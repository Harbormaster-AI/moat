import React, { Component } from 'react'
import OperationService from '../services/OperationService'

class ListOperationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                operations: []
        }
        this.addOperation = this.addOperation.bind(this);
        this.editOperation = this.editOperation.bind(this);
        this.deleteOperation = this.deleteOperation.bind(this);
    }

    deleteOperation(id){
        OperationService.deleteOperation(id).then( res => {
            this.setState({operations: this.state.operations.filter(operation => operation.operationId !== id)});
        });
    }
    viewOperation(id){
        this.props.history.push(`/view-operation/${id}`);
    }
    editOperation(id){
        this.props.history.push(`/add-operation/${id}`);
    }

    componentDidMount(){
        OperationService.getOperations().then((res) => {
            this.setState({ operations: res.data});
        });
    }

    addOperation(){
        this.props.history.push('/add-operation/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Operation List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOperation}> Add Operation</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OperationNumber </th>
                                    <th> Name </th>
                                    <th> SetupTime </th>
                                    <th> StandardCycleTime </th>
                                    <th> OperationType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.operations.map(
                                        operation => 
                                        <tr key = {operation.operationId}>
                                             <td> { operation.operationNumber } </td>
                                             <td> { operation.name } </td>
                                             <td> { operation.setupTime } </td>
                                             <td> { operation.standardCycleTime } </td>
                                             <td> { operation.operationType } </td>
                                             <td>
                                                 <button onClick={ () => this.editOperation(operation.operationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOperation(operation.operationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOperation(operation.operationId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListOperationComponent
