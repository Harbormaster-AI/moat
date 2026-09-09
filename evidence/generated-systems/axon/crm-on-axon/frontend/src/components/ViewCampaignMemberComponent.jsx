import React, { Component } from 'react'
import CampaignMemberService from '../services/CampaignMemberService'

class ViewCampaignMemberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            campaignMember: {}
        }
    }

    componentDidMount(){
        CampaignMemberService.getCampaignMemberById(this.state.id).then( res => {
            this.setState({campaignMember: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CampaignMember Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> responded:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaignMember.responded }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaignMember.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MemberType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.campaignMember.memberType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCampaignMemberComponent
