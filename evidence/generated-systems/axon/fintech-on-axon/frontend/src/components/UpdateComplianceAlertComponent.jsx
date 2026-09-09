import React, { Component } from 'react'
import ComplianceAlertService from '../services/ComplianceAlertService';

class UpdateComplianceAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                alertCode: '',
                raisedAt: '',
                notes: '',
                severity: '',
                status: ''
        }
        this.updateComplianceAlert = this.updateComplianceAlert.bind(this);

        this.changealertCodeHandler = this.changealertCodeHandler.bind(this);
        this.changeraisedAtHandler = this.changeraisedAtHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ComplianceAlertService.getComplianceAlertById(this.state.id).then( (res) =>{
            let complianceAlert = res.data;
            this.setState({
                alertCode: complianceAlert.alertCode,
                raisedAt: complianceAlert.raisedAt,
                notes: complianceAlert.notes,
                severity: complianceAlert.severity,
                status: complianceAlert.status
            });
        });
    }

    updateComplianceAlert = (e) => {
        e.preventDefault();
        let complianceAlert = {
            complianceAlertId: this.state.id,
            alertCode: this.state.alertCode,
            raisedAt: this.state.raisedAt,
            notes: this.state.notes,
            severity: this.state.severity,
            status: this.state.status
        };
        console.log('complianceAlert => ' + JSON.stringify(complianceAlert));
        console.log('id => ' + JSON.stringify(this.state.id));
        ComplianceAlertService.updateComplianceAlert(complianceAlert).then( res => {
            this.props.history.push('/complianceAlerts');
        });
    }

    changealertCodeHandler= (event) => {
        this.setState({alertCode: event.target.value});
    }
    changeraisedAtHandler= (event) => {
        this.setState({raisedAt: event.target.value});
    }
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/complianceAlerts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ComplianceAlert</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> alertCode: </label>
                                                <input placeholder="alertCode" name="alertCode" className="form-control" value={this.state.alertCode} onChange={this.changealertCodeHandler}/>

                                            <label> raisedAt: </label>
                                                <input type="time" placeholder="raisedAt" name="raisedAt" className="form-control" value={this.state.raisedAt} onChange={this.changeraisedAtHandler}/>

                                            <label> notes: </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

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
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Investigating
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          Dismissed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateComplianceAlert}>Save</button>
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

export default UpdateComplianceAlertComponent
