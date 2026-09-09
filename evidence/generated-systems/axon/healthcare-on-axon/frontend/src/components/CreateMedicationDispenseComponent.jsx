import React, { Component } from 'react'
import MedicationDispenseService from '../services/MedicationDispenseService';

class CreateMedicationDispenseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                dispenseNumber: '',
                quantity: '',
                whenPrepared: '',
                status: ''
        }
        this.changedispenseNumberHandler = this.changedispenseNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changewhenPreparedHandler = this.changewhenPreparedHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateMedicationDispense = (e) => {
        e.preventDefault();
        let medicationDispense = {
                medicationDispenseId: this.state.id,
                dispenseNumber: this.state.dispenseNumber,
                quantity: this.state.quantity,
                whenPrepared: this.state.whenPrepared,
                status: this.state.status
            };
        console.log('medicationDispense => ' + JSON.stringify(medicationDispense));

        // step 5
        if(this.state.id === '_add'){
            medicationDispense.medicationDispenseId=''
            MedicationDispenseService.createMedicationDispense(medicationDispense).then(res =>{
                this.props.history.push('/medicationDispenses');
            });
        }else{
            MedicationDispenseService.updateMedicationDispense(medicationDispense).then( res => {
                this.props.history.push('/medicationDispenses');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MedicationDispense</h3>
        }else{
            return <h3 className="text-center">Update MedicationDispense</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> dispenseNumber:&emsp; </label>
                                                <input placeholder="dispenseNumber" name="dispenseNumber" className="form-control" value={this.state.dispenseNumber} onChange={this.changedispenseNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> whenPrepared:&emsp; </label>
                                                <input type="time" placeholder="whenPrepared" name="whenPrepared" className="form-control" value={this.state.whenPrepared} onChange={this.changewhenPreparedHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMedicationDispense}>Save</button>
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

export default CreateMedicationDispenseComponent
