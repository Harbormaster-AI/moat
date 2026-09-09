import React, { Component } from 'react'
import CampaignService from '../services/CampaignService';

class UpdateCampaignComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                totalBudget: '',
                flight: '',
                objective: '',
                status: ''
        }
        this.updateCampaign = this.updateCampaign.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetotalBudgetHandler = this.changetotalBudgetHandler.bind(this);
        this.changeflightHandler = this.changeflightHandler.bind(this);
        this.changeObjectiveHandler = this.changeObjectiveHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateCampaign = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        CampaignService.updateCampaign(campaign).then( res => {
            this.props.history.push('/campaigns');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Campaign</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> totalBudget: </label>
                                                <input placeholder="totalBudget" name="totalBudget" className="form-control" value={this.state.totalBudget} onChange={this.changetotalBudgetHandler}/>

                                            <label> flight: </label>
                                                <input placeholder="flight" name="flight" className="form-control" value={this.state.flight} onChange={this.changeflightHandler}/>

                                            <label> Objective: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateCampaign}>Save</button>
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

export default UpdateCampaignComponent
