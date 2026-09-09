import React, { Component } from 'react'
import MedicationOrderService from '../services/MedicationOrderService'

class ListMedicationOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                medicationOrders: []
        }
        this.addMedicationOrder = this.addMedicationOrder.bind(this);
        this.editMedicationOrder = this.editMedicationOrder.bind(this);
        this.deleteMedicationOrder = this.deleteMedicationOrder.bind(this);
    }

    deleteMedicationOrder(id){
        MedicationOrderService.deleteMedicationOrder(id).then( res => {
            this.setState({medicationOrders: this.state.medicationOrders.filter(medicationOrder => medicationOrder.medicationOrderId !== id)});
        });
    }
    viewMedicationOrder(id){
        this.props.history.push(`/view-medicationOrder/${id}`);
    }
    editMedicationOrder(id){
        this.props.history.push(`/add-medicationOrder/${id}`);
    }

    componentDidMount(){
        MedicationOrderService.getMedicationOrders().then((res) => {
            this.setState({ medicationOrders: res.data});
        });
    }

    addMedicationOrder(){
        this.props.history.push('/add-medicationOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MedicationOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMedicationOrder}> Add MedicationOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> MedicationCode </th>
                                    <th> Dose </th>
                                    <th> Frequency </th>
                                    <th> Duration </th>
                                    <th> Route </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.medicationOrders.map(
                                        medicationOrder => 
                                        <tr key = {medicationOrder.medicationOrderId}>
                                             <td> { medicationOrder.medicationCode } </td>
                                             <td> { medicationOrder.dose } </td>
                                             <td> { medicationOrder.frequency } </td>
                                             <td> { medicationOrder.duration } </td>
                                             <td> { medicationOrder.route } </td>
                                             <td>
                                                 <button onClick={ () => this.editMedicationOrder(medicationOrder.medicationOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMedicationOrder(medicationOrder.medicationOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMedicationOrder(medicationOrder.medicationOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMedicationOrderComponent
