import React, { Component } from 'react'
import AgentService from '../services/AgentService'

class ListAgentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                agents: []
        }
        this.addAgent = this.addAgent.bind(this);
        this.editAgent = this.editAgent.bind(this);
        this.deleteAgent = this.deleteAgent.bind(this);
    }

    deleteAgent(id){
        AgentService.deleteAgent(id).then( res => {
            this.setState({agents: this.state.agents.filter(agent => agent.agentId !== id)});
        });
    }
    viewAgent(id){
        this.props.history.push(`/view-agent/${id}`);
    }
    editAgent(id){
        this.props.history.push(`/add-agent/${id}`);
    }

    componentDidMount(){
        AgentService.getAgents().then((res) => {
            this.setState({ agents: res.data});
        });
    }

    addAgent(){
        this.props.history.push('/add-agent/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Agent List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAgent}> Add Agent</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> LicenseId </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.agents.map(
                                        agent => 
                                        <tr key = {agent.agentId}>
                                             <td> { agent.firstName } </td>
                                             <td> { agent.lastName } </td>
                                             <td> { agent.licenseId } </td>
                                             <td> { agent.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAgent(agent.agentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAgent(agent.agentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAgent(agent.agentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAgentComponent
