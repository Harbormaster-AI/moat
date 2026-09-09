import React, { Component } from 'react'
import CampaignService from '../services/CampaignService'

class ViewCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            campaign: {}
        }
    }

    componentDidMount(){
        CampaignService.getCampaignById(this.state.id).then( res => {
            this.setState({campaign: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Campaign Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> budget:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.budget }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> actualCost:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.actualCost }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expectedRevenue:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.expectedRevenue }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaign.type }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCampaignComponent
