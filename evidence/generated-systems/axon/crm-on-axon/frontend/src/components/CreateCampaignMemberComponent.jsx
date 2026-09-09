import React, { Component } from 'react'
import CampaignMemberService from '../services/CampaignMemberService';

class CreateCampaignMemberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                responded: '',
                status: '',
                memberType: ''
        }
        this.changerespondedHandler = this.changerespondedHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeMemberTypeHandler = this.changeMemberTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CampaignMemberService.getCampaignMemberById(this.state.id).then( (res) =>{
                let campaignMember = res.data;
                this.setState({
                    responded: campaignMember.responded,
                    status: campaignMember.status,
                    memberType: campaignMember.memberType
                });
            });
        }        
    }
    saveOrUpdateCampaignMember = (e) => {
        e.preventDefault();
        let campaignMember = {
                campaignMemberId: this.state.id,
                responded: this.state.responded,
                status: this.state.status,
                memberType: this.state.memberType
            };
        console.log('campaignMember => ' + JSON.stringify(campaignMember));

        // step 5
        if(this.state.id === '_add'){
            campaignMember.campaignMemberId=''
            CampaignMemberService.createCampaignMember(campaignMember).then(res =>{
                this.props.history.push('/campaignMembers');
            });
        }else{
            CampaignMemberService.updateCampaignMember(campaignMember).then( res => {
                this.props.history.push('/campaignMembers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CampaignMember</h3>
        }else{
            return <h3 className="text-center">Update CampaignMember</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> responded:&emsp; </label>
                                                <input type="checkbox" placeholder="responded" name="responded" className="form-control" value={this.state.responded} onChange={this.changerespondedHandler}/>


                                            <label> Status:&emsp; </label>
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

                                            <label> MemberType:&emsp; </label>
                                                <select value={this.state.memberType} onChange={this.changeMemberTypeHandler}>
                      <option name="MemberType" className="form-control" >
                          Lead
                      </option>
                      <option name="MemberType" className="form-control" >
                          Contact
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCampaignMember}>Save</button>
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

export default CreateCampaignMemberComponent
