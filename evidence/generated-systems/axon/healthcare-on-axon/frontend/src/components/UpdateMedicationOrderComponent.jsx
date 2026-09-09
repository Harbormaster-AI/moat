import React, { Component } from 'react'
import MedicationOrderService from '../services/MedicationOrderService';

class UpdateMedicationOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                medicationCode: '',
                dose: '',
                frequency: '',
                duration: '',
                route: ''
        }
        this.updateMedicationOrder = this.updateMedicationOrder.bind(this);

        this.changemedicationCodeHandler = this.changemedicationCodeHandler.bind(this);
        this.changedoseHandler = this.changedoseHandler.bind(this);
        this.changefrequencyHandler = this.changefrequencyHandler.bind(this);
        this.changedurationHandler = this.changedurationHandler.bind(this);
        this.changeRouteHandler = this.changeRouteHandler.bind(this);
    }

    componentDidMount(){
        MedicationOrderService.getMedicationOrderById(this.state.id).then( (res) =>{
            let medicationOrder = res.data;
            this.setState({
                medicationCode: medicationOrder.medicationCode,
                dose: medicationOrder.dose,
                frequency: medicationOrder.frequency,
                duration: medicationOrder.duration,
                route: medicationOrder.route
            });
        });
    }

    updateMedicationOrder = (e) => {
        e.preventDefault();
        let medicationOrder = {
            medicationOrderId: this.state.id,
            medicationCode: this.state.medicationCode,
            dose: this.state.dose,
            frequency: this.state.frequency,
            duration: this.state.duration,
            route: this.state.route
        };
        console.log('medicationOrder => ' + JSON.stringify(medicationOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        MedicationOrderService.updateMedicationOrder(medicationOrder).then( res => {
            this.props.history.push('/medicationOrders');
        });
    }

    changemedicationCodeHandler= (event) => {
        this.setState({medicationCode: event.target.value});
    }
    changedoseHandler= (event) => {
        this.setState({dose: event.target.value});
    }
    changefrequencyHandler= (event) => {
        this.setState({frequency: event.target.value});
    }
    changedurationHandler= (event) => {
        this.setState({duration: event.target.value});
    }
    changeRouteHandler= (event) => {
        this.setState({route: event.target.value});
    }

    cancel(){
        this.props.history.push('/medicationOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MedicationOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> medicationCode: </label>
                                                <input placeholder="medicationCode" name="medicationCode" className="form-control" value={this.state.medicationCode} onChange={this.changemedicationCodeHandler}/>

                                            <label> dose: </label>
                                                <input placeholder="dose" name="dose" className="form-control" value={this.state.dose} onChange={this.changedoseHandler}/>

                                            <label> frequency: </label>
                                                <input placeholder="frequency" name="frequency" className="form-control" value={this.state.frequency} onChange={this.changefrequencyHandler}/>

                                            <label> duration: </label>
                                                <input placeholder="duration" name="duration" className="form-control" value={this.state.duration} onChange={this.changedurationHandler}/>

                                            <label> Route: </label>
                                                <select value={this.state.route} onChange={this.changeRouteHandler}>
                      <option name="Route" className="form-control" >
                          Oral
                      </option>
                      <option name="Route" className="form-control" >
                          Intravenous
                      </option>
                      <option name="Route" className="form-control" >
                          Subcutaneous
                      </option>
                      <option name="Route" className="form-control" >
                          Intramuscular
                      </option>
                      <option name="Route" className="form-control" >
                          Topical
                      </option>
                      <option name="Route" className="form-control" >
                          Inhalation
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMedicationOrder}>Save</button>
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

export default UpdateMedicationOrderComponent
