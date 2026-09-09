import React, { Component } from 'react'
import UnderwritingDecisionService from '../services/UnderwritingDecisionService';

class UpdateUnderwritingDecisionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                notes: '',
                decisionDate: '',
                decision: ''
        }
        this.updateUnderwritingDecision = this.updateUnderwritingDecision.bind(this);

        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changedecisionDateHandler = this.changedecisionDateHandler.bind(this);
        this.changeDecisionHandler = this.changeDecisionHandler.bind(this);
    }

    componentDidMount(){
        UnderwritingDecisionService.getUnderwritingDecisionById(this.state.id).then( (res) =>{
            let underwritingDecision = res.data;
            this.setState({
                notes: underwritingDecision.notes,
                decisionDate: underwritingDecision.decisionDate,
                decision: underwritingDecision.decision
            });
        });
    }

    updateUnderwritingDecision = (e) => {
        e.preventDefault();
        let underwritingDecision = {
            underwritingDecisionId: this.state.id,
            notes: this.state.notes,
            decisionDate: this.state.decisionDate,
            decision: this.state.decision
        };
        console.log('underwritingDecision => ' + JSON.stringify(underwritingDecision));
        console.log('id => ' + JSON.stringify(this.state.id));
        UnderwritingDecisionService.updateUnderwritingDecision(underwritingDecision).then( res => {
            this.props.history.push('/underwritingDecisions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update UnderwritingDecision</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> notes: </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> decisionDate: </label>
                                                <input type="date" placeholder="decisionDate" name="decisionDate" className="form-control" value={this.state.decisionDate} onChange={this.changedecisionDateHandler}/>

                                            <label> Decision: </label>
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
                                        <button className="btn btn-success" onClick={this.updateUnderwritingDecision}>Save</button>
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

export default UpdateUnderwritingDecisionComponent
