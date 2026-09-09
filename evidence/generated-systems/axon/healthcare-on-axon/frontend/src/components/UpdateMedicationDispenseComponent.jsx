import React, { Component } from 'react'
import MedicationDispenseService from '../services/MedicationDispenseService';

class UpdateMedicationDispenseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                dispenseNumber: '',
                quantity: '',
                whenPrepared: '',
                status: ''
        }
        this.updateMedicationDispense = this.updateMedicationDispense.bind(this);

        this.changedispenseNumberHandler = this.changedispenseNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changewhenPreparedHandler = this.changewhenPreparedHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MedicationDispenseService.getMedicationDispenseById(this.state.id).then( (res) =>{
            let medicationDispense = res.data;
            this.setState({
                dispenseNumber: medicationDispense.dispenseNumber,
                quantity: medicationDispense.quantity,
                whenPrepared: medicationDispense.whenPrepared,
                status: medicationDispense.status
            });
        });
    }

    updateMedicationDispense = (e) => {
        e.preventDefault();
        let medicationDispense = {
            medicationDispenseId: this.state.id,
            dispenseNumber: this.state.dispenseNumber,
            quantity: this.state.quantity,
            whenPrepared: this.state.whenPrepared,
            status: this.state.status
        };
        console.log('medicationDispense => ' + JSON.stringify(medicationDispense));
        console.log('id => ' + JSON.stringify(this.state.id));
        MedicationDispenseService.updateMedicationDispense(medicationDispense).then( res => {
            this.props.history.push('/medicationDispenses');
        });
    }

    changedispenseNumberHandler= (event) => {
        this.setState({dispenseNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changewhenPreparedHandler= (event) => {
        this.setState({whenPrepared: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/medicationDispenses');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MedicationDispense</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> dispenseNumber: </label>
                                                <input placeholder="dispenseNumber" name="dispenseNumber" className="form-control" value={this.state.dispenseNumber} onChange={this.changedispenseNumberHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> whenPrepared: </label>
                                                <input type="time" placeholder="whenPrepared" name="whenPrepared" className="form-control" value={this.state.whenPrepared} onChange={this.changewhenPreparedHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Preparation
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMedicationDispense}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateMedicationDispenseComponent
