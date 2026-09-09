import React, { Component } from 'react'
import CampaignMemberService from '../services/CampaignMemberService'

class ListCampaignMemberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                campaignMembers: []
        }
        this.addCampaignMember = this.addCampaignMember.bind(this);
        this.editCampaignMember = this.editCampaignMember.bind(this);
        this.deleteCampaignMember = this.deleteCampaignMember.bind(this);
    }

    deleteCampaignMember(id){
        CampaignMemberService.deleteCampaignMember(id).then( res => {
            this.setState({campaignMembers: this.state.campaignMembers.filter(campaignMember => campaignMember.campaignMemberId !== id)});
        });
    }
    viewCampaignMember(id){
        this.props.history.push(`/view-campaignMember/${id}`);
    }
    editCampaignMember(id){
        this.props.history.push(`/add-campaignMember/${id}`);
    }

    componentDidMount(){
        CampaignMemberService.getCampaignMembers().then((res) => {
            this.setState({ campaignMembers: res.data});
        });
    }

    addCampaignMember(){
        this.props.history.push('/add-campaignMember/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CampaignMember List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCampaignMember}> Add CampaignMember</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Responded </th>
                                    <th> Status </th>
                                    <th> MemberType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.campaignMembers.map(
                                        campaignMember => 
                                        <tr key = {campaignMember.campaignMemberId}>
                                             <td> { campaignMember.responded } </td>
                                             <td> { campaignMember.status } </td>
                                             <td> { campaignMember.memberType } </td>
                                             <td>
                                                 <button onClick={ () => this.editCampaignMember(campaignMember.campaignMemberId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCampaignMember(campaignMember.campaignMemberId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCampaignMember(campaignMember.campaignMemberId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCampaignMemberComponent
