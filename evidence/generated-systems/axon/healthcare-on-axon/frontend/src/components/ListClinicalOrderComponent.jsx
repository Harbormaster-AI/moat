import React, { Component } from 'react'
import ClinicalOrderService from '../services/ClinicalOrderService'

class ListClinicalOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                clinicalOrders: []
        }
        this.addClinicalOrder = this.addClinicalOrder.bind(this);
        this.editClinicalOrder = this.editClinicalOrder.bind(this);
        this.deleteClinicalOrder = this.deleteClinicalOrder.bind(this);
    }

    deleteClinicalOrder(id){
        ClinicalOrderService.deleteClinicalOrder(id).then( res => {
            this.setState({clinicalOrders: this.state.clinicalOrders.filter(clinicalOrder => clinicalOrder.clinicalOrderId !== id)});
        });
    }
    viewClinicalOrder(id){
        this.props.history.push(`/view-clinicalOrder/${id}`);
    }
    editClinicalOrder(id){
        this.props.history.push(`/add-clinicalOrder/${id}`);
    }

    componentDidMount(){
        ClinicalOrderService.getClinicalOrders().then((res) => {
            this.setState({ clinicalOrders: res.data});
        });
    }

    addClinicalOrder(){
        this.props.history.push('/add-clinicalOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ClinicalOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addClinicalOrder}> Add ClinicalOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OrderNumber </th>
                                    <th> Status </th>
                                    <th> OrderType </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.clinicalOrders.map(
                                        clinicalOrder => 
                                        <tr key = {clinicalOrder.clinicalOrderId}>
                                             <td> { clinicalOrder.orderNumber } </td>
                                             <td> { clinicalOrder.status } </td>
                                             <td> { clinicalOrder.orderType } </td>
                                             <td> { clinicalOrder.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editClinicalOrder(clinicalOrder.clinicalOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteClinicalOrder(clinicalOrder.clinicalOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewClinicalOrder(clinicalOrder.clinicalOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListClinicalOrderComponent
