import React, { Component } from 'react'
import SalesCampaignService from '../services/SalesCampaignService';

class CreateSalesCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                campaignCode: '',
                status: ''
        }
        this.changecampaignCodeHandler = this.changecampaignCodeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SalesCampaignService.getSalesCampaignById(this.state.id).then( (res) =>{
                let salesCampaign = res.data;
                this.setState({
                    campaignCode: salesCampaign.campaignCode,
                    status: salesCampaign.status
                });
            });
        }        
    }
    saveOrUpdateSalesCampaign = (e) => {
        e.preventDefault();
        let salesCampaign = {
                salesCampaignId: this.state.id,
                campaignCode: this.state.campaignCode,
                status: this.state.status
            };
        console.log('salesCampaign => ' + JSON.stringify(salesCampaign));

        // step 5
        if(this.state.id === '_add'){
            salesCampaign.salesCampaignId=''
            SalesCampaignService.createSalesCampaign(salesCampaign).then(res =>{
                this.props.history.push('/salesCampaigns');
            });
        }else{
            SalesCampaignService.updateSalesCampaign(salesCampaign).then( res => {
                this.props.history.push('/salesCampaigns');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SalesCampaign</h3>
        }else{
            return <h3 className="text-center">Update SalesCampaign</h3>
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
                                            <label> campaignCode:&emsp; </label>
                                                <input placeholder="campaignCode" name="campaignCode" className="form-control" value={this.state.campaignCode} onChange={this.changecampaignCodeHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSalesCampaign}>Save</button>
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

export default CreateSalesCampaignComponent
