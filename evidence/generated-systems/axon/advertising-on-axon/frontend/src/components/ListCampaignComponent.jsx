import React, { Component } from 'react'
import CampaignService from '../services/CampaignService'

class ListCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                campaigns: []
        }
        this.addCampaign = this.addCampaign.bind(this);
        this.editCampaign = this.editCampaign.bind(this);
        this.deleteCampaign = this.deleteCampaign.bind(this);
    }

    deleteCampaign(id){
        CampaignService.deleteCampaign(id).then( res => {
            this.setState({campaigns: this.state.campaigns.filter(campaign => campaign.campaignId !== id)});
        });
    }
    viewCampaign(id){
        this.props.history.push(`/view-campaign/${id}`);
    }
    editCampaign(id){
        this.props.history.push(`/add-campaign/${id}`);
    }

    componentDidMount(){
        CampaignService.getCampaigns().then((res) => {
            this.setState({ campaigns: res.data});
        });
    }

    addCampaign(){
        this.props.history.push('/add-campaign/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Campaign List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCampaign}> Add Campaign</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TotalBudget </th>
                                    <th> Flight </th>
                                    <th> Objective </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.campaigns.map(
                                        campaign => 
                                        <tr key = {campaign.campaignId}>
                                             <td> { campaign.name } </td>
                                             <td> { campaign.totalBudget } </td>
                                             <td> { campaign.flight } </td>
                                             <td> { campaign.objective } </td>
                                             <td> { campaign.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCampaign(campaign.campaignId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCampaign(campaign.campaignId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCampaign(campaign.campaignId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCampaignComponent
