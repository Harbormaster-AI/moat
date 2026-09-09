import React, { Component } from 'react'
import CampaignMemberService from '../services/CampaignMemberService';

class UpdateCampaignMemberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                responded: '',
                status: '',
                memberType: ''
        }
        this.updateCampaignMember = this.updateCampaignMember.bind(this);

        this.changerespondedHandler = this.changerespondedHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeMemberTypeHandler = this.changeMemberTypeHandler.bind(this);
    }

    componentDidMount(){
        CampaignMemberService.getCampaignMemberById(this.state.id).then( (res) =>{
            let campaignMember = res.data;
            this.setState({
                responded: campaignMember.responded,
                status: campaignMember.status,
                memberType: campaignMember.memberType
            });
        });
    }

    updateCampaignMember = (e) => {
        e.preventDefault();
        let campaignMember = {
            campaignMemberId: this.state.id,
            responded: this.state.responded,
            status: this.state.status,
            memberType: this.state.memberType
        };
        console.log('campaignMember => ' + JSON.stringify(campaignMember));
        console.log('id => ' + JSON.stringify(this.state.id));
        CampaignMemberService.updateCampaignMember(campaignMember).then( res => {
            this.props.history.push('/campaignMembers');
        });
    }

    changerespondedHandler= (event) => {
        this.setState({responded: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeMemberTypeHandler= (event) => {
        this.setState({memberType: event.target.value});
    }

    cancel(){
        this.props.history.push('/campaignMembers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CampaignMember</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> responded: </label>
                                                <input type="checkbox" placeholder="responded" name="responded" className="form-control" value={this.state.responded} onChange={this.changerespondedHandler}/>


                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Sent
                      </option>
                      <option name="Status" className="form-control" >
                          Opened
                      </option>
                      <option name="Status" className="form-control" >
                          Responded
                      </option>
                      <option name="Status" className="form-control" >
                          Unsubscribed
                      </option>
                      <option name="Status" className="form-control" >
                          Bounced
                      </option>
                      <option name="Status" className="form-control" >
                          Registered
                      </option>
                      <option name="Status" className="form-control" >
                          Attended
                      </option>
                      <option name="Status" className="form-control" >
                          NoShow
                      </option>
                    </select>

                                            <label> MemberType: </label>
                                                <select value={this.state.memberType} onChange={this.changeMemberTypeHandler}>
                      <option name="MemberType" className="form-control" >
                          Lead
                      </option>
                      <option name="MemberType" className="form-control" >
                          Contact
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCampaignMember}>Save</button>
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

export default UpdateCampaignMemberComponent
