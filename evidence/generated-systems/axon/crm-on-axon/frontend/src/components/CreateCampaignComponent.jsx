import React, { Component } from 'react'
import CampaignService from '../services/CampaignService';

class CreateCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                startDate: '',
                endDate: '',
                budget: '',
                actualCost: '',
                expectedRevenue: '',
                status: '',
                type: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changebudgetHandler = this.changebudgetHandler.bind(this);
        this.changeactualCostHandler = this.changeactualCostHandler.bind(this);
        this.changeexpectedRevenueHandler = this.changeexpectedRevenueHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CampaignService.getCampaignById(this.state.id).then( (res) =>{
                let campaign = res.data;
                this.setState({
                    name: campaign.name,
                    startDate: campaign.startDate,
                    endDate: campaign.endDate,
                    budget: campaign.budget,
                    actualCost: campaign.actualCost,
                    expectedRevenue: campaign.expectedRevenue,
                    status: campaign.status,
                    type: campaign.type
                });
            });
        }        
    }
    saveOrUpdateCampaign = (e) => {
        e.preventDefault();
        let campaign = {
                campaignId: this.state.id,
                name: this.state.name,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                budget: this.state.budget,
                actualCost: this.state.actualCost,
                expectedRevenue: this.state.expectedRevenue,
                status: this.state.status,
                type: this.state.type
            };
        console.log('campaign => ' + JSON.stringify(campaign));

        // step 5
        if(this.state.id === '_add'){
            campaign.campaignId=''
            CampaignService.createCampaign(campaign).then(res =>{
                this.props.history.push('/campaigns');
            });
        }else{
            CampaignService.updateCampaign(campaign).then( res => {
                this.props.history.push('/campaigns');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changebudgetHandler= (event) => {
        this.setState({budget: event.target.value});
    }
    changeactualCostHandler= (event) => {
        this.setState({actualCost: event.target.value});
    }
    changeexpectedRevenueHandler= (event) => {
        this.setState({expectedRevenue: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }

    cancel(){
        this.props.history.push('/campaigns');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Campaign</h3>
        }else{
            return <h3 className="text-center">Update Campaign</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> budget:&emsp; </label>
                                                <input placeholder="budget" name="budget" className="form-control" value={this.state.budget} onChange={this.changebudgetHandler}/>

                                            <label> actualCost:&emsp; </label>
                                                <input placeholder="actualCost" name="actualCost" className="form-control" value={this.state.actualCost} onChange={this.changeactualCostHandler}/>

                                            <label> expectedRevenue:&emsp; </label>
                                                <input placeholder="expectedRevenue" name="expectedRevenue" className="form-control" value={this.state.expectedRevenue} onChange={this.changeexpectedRevenueHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> Type:&emsp; </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          Email
                      </option>
                      <option name="Type" className="form-control" >
                          Social
                      </option>
                      <option name="Type" className="form-control" >
                          Event
                      </option>
                      <option name="Type" className="form-control" >
                          Webinar
                      </option>
                      <option name="Type" className="form-control" >
                          Advertising
                      </option>
                      <option name="Type" className="form-control" >
                          ContentMarketing
                      </option>
                      <option name="Type" className="form-control" >
                          Referral
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCampaign}>Save</button>
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

export default CreateCampaignComponent
