import React, { Component } from 'react'
import ControlService from '../services/ControlService';

class UpdateControlComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                objective: '',
                ownerDepartment: '',
                controlType: '',
                frequency: '',
                status: ''
        }
        this.updateControl = this.updateControl.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeobjectiveHandler = this.changeobjectiveHandler.bind(this);
        this.changeownerDepartmentHandler = this.changeownerDepartmentHandler.bind(this);
        this.changeControlTypeHandler = this.changeControlTypeHandler.bind(this);
        this.changeFrequencyHandler = this.changeFrequencyHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ControlService.getControlById(this.state.id).then( (res) =>{
            let control = res.data;
            this.setState({
                name: control.name,
                objective: control.objective,
                ownerDepartment: control.ownerDepartment,
                controlType: control.controlType,
                frequency: control.frequency,
                status: control.status
            });
        });
    }

    updateControl = (e) => {
        e.preventDefault();
        let control = {
            controlId: this.state.id,
            name: this.state.name,
            objective: this.state.objective,
            ownerDepartment: this.state.ownerDepartment,
            controlType: this.state.controlType,
            frequency: this.state.frequency,
            status: this.state.status
        };
        console.log('control => ' + JSON.stringify(control));
        console.log('id => ' + JSON.stringify(this.state.id));
        ControlService.updateControl(control).then( res => {
            this.props.history.push('/controls');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeobjectiveHandler= (event) => {
        this.setState({objective: event.target.value});
    }
    changeownerDepartmentHandler= (event) => {
        this.setState({ownerDepartment: event.target.value});
    }
    changeControlTypeHandler= (event) => {
        this.setState({controlType: event.target.value});
    }
    changeFrequencyHandler= (event) => {
        this.setState({frequency: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/controls');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Control</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> objective: </label>
                                                <input placeholder="objective" name="objective" className="form-control" value={this.state.objective} onChange={this.changeobjectiveHandler}/>

                                            <label> ownerDepartment: </label>
                                                <input placeholder="ownerDepartment" name="ownerDepartment" className="form-control" value={this.state.ownerDepartment} onChange={this.changeownerDepartmentHandler}/>

                                            <label> ControlType: </label>
                                                <select value={this.state.controlType} onChange={this.changeControlTypeHandler}>
                      <option name="ControlType" className="form-control" >
                          Preventive
                      </option>
                      <option name="ControlType" className="form-control" >
                          Detective
                      </option>
                      <option name="ControlType" className="form-control" >
                          Corrective
                      </option>
                      <option name="ControlType" className="form-control" >
                          Directive
                      </option>
                    </select>

                                            <label> Frequency: </label>
                                                <select value={this.state.frequency} onChange={this.changeFrequencyHandler}>
                      <option name="Frequency" className="form-control" >
                          Continuous
                      </option>
                      <option name="Frequency" className="form-control" >
                          Daily
                      </option>
                      <option name="Frequency" className="form-control" >
                          Weekly
                      </option>
                      <option name="Frequency" className="form-control" >
                          Monthly
                      </option>
                      <option name="Frequency" className="form-control" >
                          Quarterly
                      </option>
                      <option name="Frequency" className="form-control" >
                          Annually
                      </option>
                      <option name="Frequency" className="form-control" >
                          AdHoc
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Designed
                      </option>
                      <option name="Status" className="form-control" >
                          Implemented
                      </option>
                      <option name="Status" className="form-control" >
                          Operating
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateControl}>Save</button>
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

export default UpdateControlComponent
