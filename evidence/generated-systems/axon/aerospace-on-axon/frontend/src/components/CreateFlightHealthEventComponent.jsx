import React, { Component } from 'react'
import FlightHealthEventService from '../services/FlightHealthEventService';

class CreateFlightHealthEventComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                eventCode: '',
                severity: ''
        }
        this.changeeventCodeHandler = this.changeeventCodeHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FlightHealthEventService.getFlightHealthEventById(this.state.id).then( (res) =>{
                let flightHealthEvent = res.data;
                this.setState({
                    eventCode: flightHealthEvent.eventCode,
                    severity: flightHealthEvent.severity
                });
            });
        }        
    }
    saveOrUpdateFlightHealthEvent = (e) => {
        e.preventDefault();
        let flightHealthEvent = {
                flightHealthEventId: this.state.id,
                eventCode: this.state.eventCode,
                severity: this.state.severity
            };
        console.log('flightHealthEvent => ' + JSON.stringify(flightHealthEvent));

        // step 5
        if(this.state.id === '_add'){
            flightHealthEvent.flightHealthEventId=''
            FlightHealthEventService.createFlightHealthEvent(flightHealthEvent).then(res =>{
                this.props.history.push('/flightHealthEvents');
            });
        }else{
            FlightHealthEventService.updateFlightHealthEvent(flightHealthEvent).then( res => {
                this.props.history.push('/flightHealthEvents');
            });
        }
    }
    
    changeeventCodeHandler= (event) => {
        this.setState({eventCode: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }

    cancel(){
        this.props.history.push('/flightHealthEvents');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FlightHealthEvent</h3>
        }else{
            return <h3 className="text-center">Update FlightHealthEvent</h3>
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
                                            <label> eventCode:&emsp; </label>
                                                <input placeholder="eventCode" name="eventCode" className="form-control" value={this.state.eventCode} onChange={this.changeeventCodeHandler}/>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Info
                      </option>
                      <option name="Severity" className="form-control" >
                          Warning
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFlightHealthEvent}>Save</button>
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

export default CreateFlightHealthEventComponent
