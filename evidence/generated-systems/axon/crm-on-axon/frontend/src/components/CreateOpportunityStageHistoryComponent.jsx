import React, { Component } from 'react'
import OpportunityStageHistoryService from '../services/OpportunityStageHistoryService';

class CreateOpportunityStageHistoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                changedAt: '',
                comment: '',
                fromStage: '',
                toStage: ''
        }
        this.changechangedAtHandler = this.changechangedAtHandler.bind(this);
        this.changecommentHandler = this.changecommentHandler.bind(this);
        this.changeFromStageHandler = this.changeFromStageHandler.bind(this);
        this.changeToStageHandler = this.changeToStageHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OpportunityStageHistoryService.getOpportunityStageHistoryById(this.state.id).then( (res) =>{
                let opportunityStageHistory = res.data;
                this.setState({
                    changedAt: opportunityStageHistory.changedAt,
                    comment: opportunityStageHistory.comment,
                    fromStage: opportunityStageHistory.fromStage,
                    toStage: opportunityStageHistory.toStage
                });
            });
        }        
    }
    saveOrUpdateOpportunityStageHistory = (e) => {
        e.preventDefault();
        let opportunityStageHistory = {
                opportunityStageHistoryId: this.state.id,
                changedAt: this.state.changedAt,
                comment: this.state.comment,
                fromStage: this.state.fromStage,
                toStage: this.state.toStage
            };
        console.log('opportunityStageHistory => ' + JSON.stringify(opportunityStageHistory));

        // step 5
        if(this.state.id === '_add'){
            opportunityStageHistory.opportunityStageHistoryId=''
            OpportunityStageHistoryService.createOpportunityStageHistory(opportunityStageHistory).then(res =>{
                this.props.history.push('/opportunityStageHistorys');
            });
        }else{
            OpportunityStageHistoryService.updateOpportunityStageHistory(opportunityStageHistory).then( res => {
                this.props.history.push('/opportunityStageHistorys');
            });
        }
    }
    
    changechangedAtHandler= (event) => {
        this.setState({changedAt: event.target.value});
    }
    changecommentHandler= (event) => {
        this.setState({comment: event.target.value});
    }
    changeFromStageHandler= (event) => {
        this.setState({fromStage: event.target.value});
    }
    changeToStageHandler= (event) => {
        this.setState({toStage: event.target.value});
    }

    cancel(){
        this.props.history.push('/opportunityStageHistorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add OpportunityStageHistory</h3>
        }else{
            return <h3 className="text-center">Update OpportunityStageHistory</h3>
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
                                            <label> changedAt:&emsp; </label>
                                                <input type="time" placeholder="changedAt" name="changedAt" className="form-control" value={this.state.changedAt} onChange={this.changechangedAtHandler}/>

                                            <label> comment:&emsp; </label>
                                                <input placeholder="comment" name="comment" className="form-control" value={this.state.comment} onChange={this.changecommentHandler}/>

                                            <label> FromStage:&emsp; </label>
                                                <select value={this.state.fromStage} onChange={this.changeFromStageHandler}>
                      <option name="FromStage" className="form-control" >
                          Qualification
                      </option>
                      <option name="FromStage" className="form-control" >
                          Discovery
                      </option>
                      <option name="FromStage" className="form-control" >
                          Proposal
                      </option>
                      <option name="FromStage" className="form-control" >
                          Negotiation
                      </option>
                      <option name="FromStage" className="form-control" >
                          ClosedWon
                      </option>
                      <option name="FromStage" className="form-control" >
                          ClosedLost
                      </option>
                    </select>

                                            <label> ToStage:&emsp; </label>
                                                <select value={this.state.toStage} onChange={this.changeToStageHandler}>
                      <option name="ToStage" className="form-control" >
                          Qualification
                      </option>
                      <option name="ToStage" className="form-control" >
                          Discovery
                      </option>
                      <option name="ToStage" className="form-control" >
                          Proposal
                      </option>
                      <option name="ToStage" className="form-control" >
                          Negotiation
                      </option>
                      <option name="ToStage" className="form-control" >
                          ClosedWon
                      </option>
                      <option name="ToStage" className="form-control" >
                          ClosedLost
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOpportunityStageHistory}>Save</button>
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

export default CreateOpportunityStageHistoryComponent
