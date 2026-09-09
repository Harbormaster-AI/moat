import React, { Component } from 'react'
import UnderwritingDecisionService from '../services/UnderwritingDecisionService';

class CreateUnderwritingDecisionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                notes: '',
                decisionDate: '',
                decision: ''
        }
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changedecisionDateHandler = this.changedecisionDateHandler.bind(this);
        this.changeDecisionHandler = this.changeDecisionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            UnderwritingDecisionService.getUnderwritingDecisionById(this.state.id).then( (res) =>{
                let underwritingDecision = res.data;
                this.setState({
                    notes: underwritingDecision.notes,
                    decisionDate: underwritingDecision.decisionDate,
                    decision: underwritingDecision.decision
                });
            });
        }        
    }
    saveOrUpdateUnderwritingDecision = (e) => {
        e.preventDefault();
        let underwritingDecision = {
                underwritingDecisionId: this.state.id,
                notes: this.state.notes,
                decisionDate: this.state.decisionDate,
                decision: this.state.decision
            };
        console.log('underwritingDecision => ' + JSON.stringify(underwritingDecision));

        // step 5
        if(this.state.id === '_add'){
            underwritingDecision.underwritingDecisionId=''
            UnderwritingDecisionService.createUnderwritingDecision(underwritingDecision).then(res =>{
                this.props.history.push('/underwritingDecisions');
            });
        }else{
            UnderwritingDecisionService.updateUnderwritingDecision(underwritingDecision).then( res => {
                this.props.history.push('/underwritingDecisions');
            });
        }
    }
    
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changedecisionDateHandler= (event) => {
        this.setState({decisionDate: event.target.value});
    }
    changeDecisionHandler= (event) => {
        this.setState({decision: event.target.value});
    }

    cancel(){
        this.props.history.push('/underwritingDecisions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add UnderwritingDecision</h3>
        }else{
            return <h3 className="text-center">Update UnderwritingDecision</h3>
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
                                            <label> notes:&emsp; </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> decisionDate:&emsp; </label>
                                                <input type="date" placeholder="decisionDate" name="decisionDate" className="form-control" value={this.state.decisionDate} onChange={this.changedecisionDateHandler}/>

                                            <label> Decision:&emsp; </label>
                                                <select value={this.state.decision} onChange={this.changeDecisionHandler}>
                      <option name="Decision" className="form-control" >
                          Approve
                      </option>
                      <option name="Decision" className="form-control" >
                          ConditionalApprove
                      </option>
                      <option name="Decision" className="form-control" >
                          Refer
                      </option>
                      <option name="Decision" className="form-control" >
                          Decline
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateUnderwritingDecision}>Save</button>
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

export default CreateUnderwritingDecisionComponent
