import React, { Component } from 'react'
import OperatorService from '../services/OperatorService'

class ListOperatorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                operators: []
        }
        this.addOperator = this.addOperator.bind(this);
        this.editOperator = this.editOperator.bind(this);
        this.deleteOperator = this.deleteOperator.bind(this);
    }

    deleteOperator(id){
        OperatorService.deleteOperator(id).then( res => {
            this.setState({operators: this.state.operators.filter(operator => operator.operatorId !== id)});
        });
    }
    viewOperator(id){
        this.props.history.push(`/view-operator/${id}`);
    }
    editOperator(id){
        this.props.history.push(`/add-operator/${id}`);
    }

    componentDidMount(){
        OperatorService.getOperators().then((res) => {
            this.setState({ operators: res.data});
        });
    }

    addOperator(){
        this.props.history.push('/add-operator/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Operator List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOperator}> Add Operator</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> IcaoDesignator </th>
                                    <th> OperatorType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.operators.map(
                                        operator => 
                                        <tr key = {operator.operatorId}>
                                             <td> { operator.name } </td>
                                             <td> { operator.icaoDesignator } </td>
                                             <td> { operator.operatorType } </td>
                                             <td>
                                                 <button onClick={ () => this.editOperator(operator.operatorId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOperator(operator.operatorId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOperator(operator.operatorId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOperatorComponent
