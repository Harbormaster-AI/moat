import React, { Component } from 'react'
import AgentService from '../services/AgentService'

class ViewAgentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            agent: {}
        }
    }

    componentDidMount(){
        AgentService.getAgentById(this.state.id).then( res => {
            this.setState({agent: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Agent Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agent.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agent.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> licenseId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agent.licenseId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agent.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAgentComponent
