import React, { Component } from 'react'
import SalesCampaignService from '../services/SalesCampaignService';

class UpdateSalesCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                campaignCode: '',
                status: ''
        }
        this.updateSalesCampaign = this.updateSalesCampaign.bind(this);

        this.changecampaignCodeHandler = this.changecampaignCodeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        SalesCampaignService.getSalesCampaignById(this.state.id).then( (res) =>{
            let salesCampaign = res.data;
            this.setState({
                campaignCode: salesCampaign.campaignCode,
                status: salesCampaign.status
            });
        });
    }

    updateSalesCampaign = (e) => {
        e.preventDefault();
        let salesCampaign = {
            salesCampaignId: this.state.id,
            campaignCode: this.state.campaignCode,
            status: this.state.status
        };
        console.log('salesCampaign => ' + JSON.stringify(salesCampaign));
        console.log('id => ' + JSON.stringify(this.state.id));
        SalesCampaignService.updateSalesCampaign(salesCampaign).then( res => {
            this.props.history.push('/salesCampaigns');
        });
    }

    changecampaignCodeHandler= (event) => {
        this.setState({campaignCode: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/salesCampaigns');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SalesCampaign</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> campaignCode: </label>
                                                <input placeholder="campaignCode" name="campaignCode" className="form-control" value={this.state.campaignCode} onChange={this.changecampaignCodeHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Prospecting
                      </option>
                      <option name="Status" className="form-control" >
                          Proposal
                      </option>
                      <option name="Status" className="form-control" >
                          Negotiation
                      </option>
                      <option name="Status" className="form-control" >
                          Won
                      </option>
                      <option name="Status" className="form-control" >
                          Lost
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSalesCampaign}>Save</button>
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

export default UpdateSalesCampaignComponent
