import React, { Component } from 'react'
import OpportunityService from '../services/OpportunityService';

class CreateOpportunityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                amount: '',
                closeDate: '',
                probability: '',
                description: '',
                stage: '',
                type: '',
                forecastCategory: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changecloseDateHandler = this.changecloseDateHandler.bind(this);
        this.changeprobabilityHandler = this.changeprobabilityHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeStageHandler = this.changeStageHandler.bind(this);
        this.changeTypeHandler = this.changeTypeHandler.bind(this);
        this.changeForecastCategoryHandler = this.changeForecastCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OpportunityService.getOpportunityById(this.state.id).then( (res) =>{
                let opportunity = res.data;
                this.setState({
                    name: opportunity.name,
                    amount: opportunity.amount,
                    closeDate: opportunity.closeDate,
                    probability: opportunity.probability,
                    description: opportunity.description,
                    stage: opportunity.stage,
                    type: opportunity.type,
                    forecastCategory: opportunity.forecastCategory
                });
            });
        }        
    }
    saveOrUpdateOpportunity = (e) => {
        e.preventDefault();
        let opportunity = {
                opportunityId: this.state.id,
                name: this.state.name,
                amount: this.state.amount,
                closeDate: this.state.closeDate,
                probability: this.state.probability,
                description: this.state.description,
                stage: this.state.stage,
                type: this.state.type,
                forecastCategory: this.state.forecastCategory
            };
        console.log('opportunity => ' + JSON.stringify(opportunity));

        // step 5
        if(this.state.id === '_add'){
            opportunity.opportunityId=''
            OpportunityService.createOpportunity(opportunity).then(res =>{
                this.props.history.push('/opportunitys');
            });
        }else{
            OpportunityService.updateOpportunity(opportunity).then( res => {
                this.props.history.push('/opportunitys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changecloseDateHandler= (event) => {
        this.setState({closeDate: event.target.value});
    }
    changeprobabilityHandler= (event) => {
        this.setState({probability: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeStageHandler= (event) => {
        this.setState({stage: event.target.value});
    }
    changeTypeHandler= (event) => {
        this.setState({type: event.target.value});
    }
    changeForecastCategoryHandler= (event) => {
        this.setState({forecastCategory: event.target.value});
    }

    cancel(){
        this.props.history.push('/opportunitys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Opportunity</h3>
        }else{
            return <h3 className="text-center">Update Opportunity</h3>
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

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> closeDate:&emsp; </label>
                                                <input type="date" placeholder="closeDate" name="closeDate" className="form-control" value={this.state.closeDate} onChange={this.changecloseDateHandler}/>

                                            <label> probability:&emsp; </label>
                                                <input placeholder="probability" name="probability" className="form-control" value={this.state.probability} onChange={this.changeprobabilityHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> Stage:&emsp; </label>
                                                <select value={this.state.stage} onChange={this.changeStageHandler}>
                      <option name="Stage" className="form-control" >
                          Qualification
                      </option>
                      <option name="Stage" className="form-control" >
                          Discovery
                      </option>
                      <option name="Stage" className="form-control" >
                          Proposal
                      </option>
                      <option name="Stage" className="form-control" >
                          Negotiation
                      </option>
                      <option name="Stage" className="form-control" >
                          ClosedWon
                      </option>
                      <option name="Stage" className="form-control" >
                          ClosedLost
                      </option>
                    </select>

                                            <label> Type:&emsp; </label>
                                                <select value={this.state.type} onChange={this.changeTypeHandler}>
                      <option name="Type" className="form-control" >
                          NewBusiness
                      </option>
                      <option name="Type" className="form-control" >
                          ExistingBusiness
                      </option>
                      <option name="Type" className="form-control" >
                          Renewal
                      </option>
                      <option name="Type" className="form-control" >
                          Upsell
                      </option>
                      <option name="Type" className="form-control" >
                          CrossSell
                      </option>
                    </select>

                                            <label> ForecastCategory:&emsp; </label>
                                                <select value={this.state.forecastCategory} onChange={this.changeForecastCategoryHandler}>
                      <option name="ForecastCategory" className="form-control" >
                          Pipeline
                      </option>
                      <option name="ForecastCategory" className="form-control" >
                          BestCase
                      </option>
                      <option name="ForecastCategory" className="form-control" >
                          Commit
                      </option>
                      <option name="ForecastCategory" className="form-control" >
                          Omitted
                      </option>
                      <option name="ForecastCategory" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOpportunity}>Save</button>
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

export default CreateOpportunityComponent
