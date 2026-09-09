import React, { Component } from 'react'
import DataBreachService from '../services/DataBreachService';

class UpdateDataBreachComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                incidentDate: '',
                description: '',
                recordsAffected: '',
                notificationRequired: '',
                severity: '',
                status: ''
        }
        this.updateDataBreach = this.updateDataBreach.bind(this);

        this.changeincidentDateHandler = this.changeincidentDateHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changerecordsAffectedHandler = this.changerecordsAffectedHandler.bind(this);
        this.changenotificationRequiredHandler = this.changenotificationRequiredHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        DataBreachService.getDataBreachById(this.state.id).then( (res) =>{
            let dataBreach = res.data;
            this.setState({
                incidentDate: dataBreach.incidentDate,
                description: dataBreach.description,
                recordsAffected: dataBreach.recordsAffected,
                notificationRequired: dataBreach.notificationRequired,
                severity: dataBreach.severity,
                status: dataBreach.status
            });
        });
    }

    updateDataBreach = (e) => {
        e.preventDefault();
        let dataBreach = {
            dataBreachId: this.state.id,
            incidentDate: this.state.incidentDate,
            description: this.state.description,
            recordsAffected: this.state.recordsAffected,
            notificationRequired: this.state.notificationRequired,
            severity: this.state.severity,
            status: this.state.status
        };
        console.log('dataBreach => ' + JSON.stringify(dataBreach));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataBreachService.updateDataBreach(dataBreach).then( res => {
            this.props.history.push('/dataBreachs');
        });
    }

    changeincidentDateHandler= (event) => {
        this.setState({incidentDate: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changerecordsAffectedHandler= (event) => {
        this.setState({recordsAffected: event.target.value});
    }
    changenotificationRequiredHandler= (event) => {
        this.setState({notificationRequired: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataBreachs');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataBreach</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> incidentDate: </label>
                                                <input type="date" placeholder="incidentDate" name="incidentDate" className="form-control" value={this.state.incidentDate} onChange={this.changeincidentDateHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> recordsAffected: </label>
                                                <input type="number" placeholder="recordsAffected" name="recordsAffected" className="form-control" value={this.state.recordsAffected} onChange={this.changerecordsAffectedHandler}/>

                                            <label> notificationRequired: </label>
                                                <input type="checkbox" placeholder="notificationRequired" name="notificationRequired" className="form-control" value={this.state.notificationRequired} onChange={this.changenotificationRequiredHandler}/>


                                            <label> Severity: </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Low
                      </option>
                      <option name="Severity" className="form-control" >
                          Medium
                      </option>
                      <option name="Severity" className="form-control" >
                          High
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Identified
                      </option>
                      <option name="Status" className="form-control" >
                          Contained
                      </option>
                      <option name="Status" className="form-control" >
                          Notified
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataBreach}>Save</button>
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

export default UpdateDataBreachComponent
