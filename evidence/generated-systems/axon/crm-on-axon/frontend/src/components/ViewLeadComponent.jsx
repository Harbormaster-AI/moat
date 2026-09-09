import React, { Component } from 'react'
import LeadService from '../services/LeadService'

class ViewLeadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            lead: {}
        }
    }

    componentDidMount(){
        LeadService.getLeadById(this.state.id).then( res => {
            this.setState({lead: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Lead Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> company:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.company }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> email:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.email }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> phone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.phone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> converted:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.converted }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Source:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.source }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Rating:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lead.rating }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLeadComponent
