import React, { Component } from 'react'
import AgentService from '../services/AgentService';

class UpdateAgentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                licenseId: '',
                status: ''
        }
        this.updateAgent = this.updateAgent.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changelicenseIdHandler = this.changelicenseIdHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AgentService.getAgentById(this.state.id).then( (res) =>{
            let agent = res.data;
            this.setState({
                firstName: agent.firstName,
                lastName: agent.lastName,
                licenseId: agent.licenseId,
                status: agent.status
            });
        });
    }

    updateAgent = (e) => {
        e.preventDefault();
        let agent = {
            agentId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            licenseId: this.state.licenseId,
            status: this.state.status
        };
        console.log('agent => ' + JSON.stringify(agent));
        console.log('id => ' + JSON.stringify(this.state.id));
        AgentService.updateAgent(agent).then( res => {
            this.props.history.push('/agents');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changelicenseIdHandler= (event) => {
        this.setState({licenseId: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/agents');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Agent</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> licenseId: </label>
                                                <input placeholder="licenseId" name="licenseId" className="form-control" value={this.state.licenseId} onChange={this.changelicenseIdHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAgent}>Save</button>
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

export default UpdateAgentComponent
