import React, { Component } from 'react'
import MedicationDispenseService from '../services/MedicationDispenseService'

class ListMedicationDispenseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                medicationDispenses: []
        }
        this.addMedicationDispense = this.addMedicationDispense.bind(this);
        this.editMedicationDispense = this.editMedicationDispense.bind(this);
        this.deleteMedicationDispense = this.deleteMedicationDispense.bind(this);
    }

    deleteMedicationDispense(id){
        MedicationDispenseService.deleteMedicationDispense(id).then( res => {
            this.setState({medicationDispenses: this.state.medicationDispenses.filter(medicationDispense => medicationDispense.medicationDispenseId !== id)});
        });
    }
    viewMedicationDispense(id){
        this.props.history.push(`/view-medicationDispense/${id}`);
    }
    editMedicationDispense(id){
        this.props.history.push(`/add-medicationDispense/${id}`);
    }

    componentDidMount(){
        MedicationDispenseService.getMedicationDispenses().then((res) => {
            this.setState({ medicationDispenses: res.data});
        });
    }

    addMedicationDispense(){
        this.props.history.push('/add-medicationDispense/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MedicationDispense List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMedicationDispense}> Add MedicationDispense</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> DispenseNumber </th>
                                    <th> Quantity </th>
                                    <th> WhenPrepared </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.medicationDispenses.map(
                                        medicationDispense => 
                                        <tr key = {medicationDispense.medicationDispenseId}>
                                             <td> { medicationDispense.dispenseNumber } </td>
                                             <td> { medicationDispense.quantity } </td>
                                             <td> { medicationDispense.whenPrepared } </td>
                                             <td> { medicationDispense.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editMedicationDispense(medicationDispense.medicationDispenseId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMedicationDispense(medicationDispense.medicationDispenseId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMedicationDispense(medicationDispense.medicationDispenseId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMedicationDispenseComponent
