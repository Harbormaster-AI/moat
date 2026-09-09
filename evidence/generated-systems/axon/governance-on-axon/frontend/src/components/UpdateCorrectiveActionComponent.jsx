import React, { Component } from 'react'
import CorrectiveActionService from '../services/CorrectiveActionService';

class UpdateCorrectiveActionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                actionTitle: '',
                owner: '',
                targetDate: '',
                status: ''
        }
        this.updateCorrectiveAction = this.updateCorrectiveAction.bind(this);

        this.changeactionTitleHandler = this.changeactionTitleHandler.bind(this);
        this.changeownerHandler = this.changeownerHandler.bind(this);
        this.changetargetDateHandler = this.changetargetDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CorrectiveActionService.getCorrectiveActionById(this.state.id).then( (res) =>{
            let correctiveAction = res.data;
            this.setState({
                actionTitle: correctiveAction.actionTitle,
                owner: correctiveAction.owner,
                targetDate: correctiveAction.targetDate,
                status: correctiveAction.status
            });
        });
    }

    updateCorrectiveAction = (e) => {
        e.preventDefault();
        let correctiveAction = {
            correctiveActionId: this.state.id,
            actionTitle: this.state.actionTitle,
            owner: this.state.owner,
            targetDate: this.state.targetDate,
            status: this.state.status
        };
        console.log('correctiveAction => ' + JSON.stringify(correctiveAction));
        console.log('id => ' + JSON.stringify(this.state.id));
        CorrectiveActionService.updateCorrectiveAction(correctiveAction).then( res => {
            this.props.history.push('/correctiveActions');
        });
    }

    changeactionTitleHandler= (event) => {
        this.setState({actionTitle: event.target.value});
    }
    changeownerHandler= (event) => {
        this.setState({owner: event.target.value});
    }
    changetargetDateHandler= (event) => {
        this.setState({targetDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/correctiveActions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CorrectiveAction</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> actionTitle: </label>
                                                <input placeholder="actionTitle" name="actionTitle" className="form-control" value={this.state.actionTitle} onChange={this.changeactionTitleHandler}/>

                                            <label> owner: </label>
                                                <input placeholder="owner" name="owner" className="form-control" value={this.state.owner} onChange={this.changeownerHandler}/>

                                            <label> targetDate: </label>
                                                <input type="date" placeholder="targetDate" name="targetDate" className="form-control" value={this.state.targetDate} onChange={this.changetargetDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Deferred
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCorrectiveAction}>Save</button>
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

export default UpdateCorrectiveActionComponent
