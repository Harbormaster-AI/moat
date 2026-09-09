import React, { Component } from 'react'
import ProcedureOrderService from '../services/ProcedureOrderService'

class ListProcedureOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                procedureOrders: []
        }
        this.addProcedureOrder = this.addProcedureOrder.bind(this);
        this.editProcedureOrder = this.editProcedureOrder.bind(this);
        this.deleteProcedureOrder = this.deleteProcedureOrder.bind(this);
    }

    deleteProcedureOrder(id){
        ProcedureOrderService.deleteProcedureOrder(id).then( res => {
            this.setState({procedureOrders: this.state.procedureOrders.filter(procedureOrder => procedureOrder.procedureOrderId !== id)});
        });
    }
    viewProcedureOrder(id){
        this.props.history.push(`/view-procedureOrder/${id}`);
    }
    editProcedureOrder(id){
        this.props.history.push(`/add-procedureOrder/${id}`);
    }

    componentDidMount(){
        ProcedureOrderService.getProcedureOrders().then((res) => {
            this.setState({ procedureOrders: res.data});
        });
    }

    addProcedureOrder(){
        this.props.history.push('/add-procedureOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProcedureOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProcedureOrder}> Add ProcedureOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ProcedureCode </th>
                                    <th> ConsentObtained </th>
                                    <th> AnesthesiaType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.procedureOrders.map(
                                        procedureOrder => 
                                        <tr key = {procedureOrder.procedureOrderId}>
                                             <td> { procedureOrder.procedureCode } </td>
                                             <td> { procedureOrder.consentObtained } </td>
                                             <td> { procedureOrder.anesthesiaType } </td>
                                             <td>
                                                 <button onClick={ () => this.editProcedureOrder(procedureOrder.procedureOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProcedureOrder(procedureOrder.procedureOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProcedureOrder(procedureOrder.procedureOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProcedureOrderComponent
