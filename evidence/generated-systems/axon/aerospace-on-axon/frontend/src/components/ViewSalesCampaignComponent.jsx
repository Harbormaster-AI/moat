import React, { Component } from 'react'
import SalesCampaignService from '../services/SalesCampaignService'

class ViewSalesCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            salesCampaign: {}
        }
    }

    componentDidMount(){
        SalesCampaignService.getSalesCampaignById(this.state.id).then( res => {
            this.setState({salesCampaign: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SalesCampaign Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> campaignCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesCampaign.campaignCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.salesCampaign.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSalesCampaignComponent
