import React, { Component } from 'react'
import AlertService from '../services/AlertService';

class UpdateAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                createdAt: '',
                severity: '',
                status: ''
        }
        this.updateAlert = this.updateAlert.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AlertService.getAlertById(this.state.id).then( (res) =>{
            let alert = res.data;
            this.setState({
                title: alert.title,
                createdAt: alert.createdAt,
                severity: alert.severity,
                status: alert.status
            });
        });
    }

    updateAlert = (e) => {
        e.preventDefault();
        let alert = {
            alertId: this.state.id,
            title: this.state.title,
            createdAt: this.state.createdAt,
            severity: this.state.severity,
            status: this.state.status
        };
        console.log('alert => ' + JSON.stringify(alert));
        console.log('id => ' + JSON.stringify(this.state.id));
        AlertService.updateAlert(alert).then( res => {
            this.props.history.push('/alerts');
        });
    }

    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/alerts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Alert</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Severity: </label>
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

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Acknowledged
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          Suppressed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAlert}>Save</button>
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

export default UpdateAlertComponent
