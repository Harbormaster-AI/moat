import React, { Component } from 'react'
import CampaignService from '../services/CampaignService';

class CreateCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                totalBudget: '',
                flight: '',
                objective: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetotalBudgetHandler = this.changetotalBudgetHandler.bind(this);
        this.changeflightHandler = this.changeflightHandler.bind(this);
        this.changeObjectiveHandler = this.changeObjectiveHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
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
                    totalBudget: campaign.totalBudget,
                    flight: campaign.flight,
                    objective: campaign.objective,
                    status: campaign.status
                });
            });
        }        
    }
    saveOrUpdateCampaign = (e) => {
        e.preventDefault();
        let campaign = {
                campaignId: this.state.id,
                name: this.state.name,
                totalBudget: this.state.totalBudget,
                flight: this.state.flight,
                objective: this.state.objective,
                status: this.state.status
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
    changetotalBudgetHandler= (event) => {
        this.setState({totalBudget: event.target.value});
    }
    changeflightHandler= (event) => {
        this.setState({flight: event.target.value});
    }
    changeObjectiveHandler= (event) => {
        this.setState({objective: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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

                                            <label> totalBudget:&emsp; </label>
                                                <input placeholder="totalBudget" name="totalBudget" className="form-control" value={this.state.totalBudget} onChange={this.changetotalBudgetHandler}/>

                                            <label> flight:&emsp; </label>
                                                <input placeholder="flight" name="flight" className="form-control" value={this.state.flight} onChange={this.changeflightHandler}/>

                                            <label> Objective:&emsp; </label>
                                                <select value={this.state.objective} onChange={this.changeObjectiveHandler}>
                      <option name="Objective" className="form-control" >
                          Awareness
                      </option>
                      <option name="Objective" className="form-control" >
                          Reach
                      </option>
                      <option name="Objective" className="form-control" >
                          Traffic
                      </option>
                      <option name="Objective" className="form-control" >
                          Engagement
                      </option>
                      <option name="Objective" className="form-control" >
                          Leads
                      </option>
                      <option name="Objective" className="form-control" >
                          Sales
                      </option>
                      <option name="Objective" className="form-control" >
                          AppInstalls
                      </option>
                      <option name="Objective" className="form-control" >
                          VideoViews
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Paused
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
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
