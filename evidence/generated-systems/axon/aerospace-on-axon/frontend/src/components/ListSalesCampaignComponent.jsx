import React, { Component } from 'react'
import SalesCampaignService from '../services/SalesCampaignService'

class ListSalesCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                salesCampaigns: []
        }
        this.addSalesCampaign = this.addSalesCampaign.bind(this);
        this.editSalesCampaign = this.editSalesCampaign.bind(this);
        this.deleteSalesCampaign = this.deleteSalesCampaign.bind(this);
    }

    deleteSalesCampaign(id){
        SalesCampaignService.deleteSalesCampaign(id).then( res => {
            this.setState({salesCampaigns: this.state.salesCampaigns.filter(salesCampaign => salesCampaign.salesCampaignId !== id)});
        });
    }
    viewSalesCampaign(id){
        this.props.history.push(`/view-salesCampaign/${id}`);
    }
    editSalesCampaign(id){
        this.props.history.push(`/add-salesCampaign/${id}`);
    }

    componentDidMount(){
        SalesCampaignService.getSalesCampaigns().then((res) => {
            this.setState({ salesCampaigns: res.data});
        });
    }

    addSalesCampaign(){
        this.props.history.push('/add-salesCampaign/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SalesCampaign List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSalesCampaign}> Add SalesCampaign</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CampaignCode </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.salesCampaigns.map(
                                        salesCampaign => 
                                        <tr key = {salesCampaign.salesCampaignId}>
                                             <td> { salesCampaign.campaignCode } </td>
                                             <td> { salesCampaign.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSalesCampaign(salesCampaign.salesCampaignId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSalesCampaign(salesCampaign.salesCampaignId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSalesCampaign(salesCampaign.salesCampaignId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSalesCampaignComponent
