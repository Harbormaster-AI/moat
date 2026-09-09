import React, { Component } from 'react'
import CorrectiveActionService from '../services/CorrectiveActionService';

class UpdateCorrectiveActionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                capaNumber: '',
                rootCause: '',
                correctiveAction: '',
                verificationDate: '',
                status: ''
        }
        this.updateCorrectiveAction = this.updateCorrectiveAction.bind(this);

        this.changecapaNumberHandler = this.changecapaNumberHandler.bind(this);
        this.changerootCauseHandler = this.changerootCauseHandler.bind(this);
        this.changecorrectiveActionHandler = this.changecorrectiveActionHandler.bind(this);
        this.changeverificationDateHandler = this.changeverificationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CorrectiveActionService.getCorrectiveActionById(this.state.id).then( (res) =>{
            let correctiveAction = res.data;
            this.setState({
                capaNumber: correctiveAction.capaNumber,
                rootCause: correctiveAction.rootCause,
                correctiveAction: correctiveAction.correctiveAction,
                verificationDate: correctiveAction.verificationDate,
                status: correctiveAction.status
            });
        });
    }

    updateCorrectiveAction = (e) => {
        e.preventDefault();
        let correctiveAction = {
            correctiveActionId: this.state.id,
            capaNumber: this.state.capaNumber,
            rootCause: this.state.rootCause,
            correctiveAction: this.state.correctiveAction,
            verificationDate: this.state.verificationDate,
            status: this.state.status
        };
        console.log('correctiveAction => ' + JSON.stringify(correctiveAction));
        console.log('id => ' + JSON.stringify(this.state.id));
        CorrectiveActionService.updateCorrectiveAction(correctiveAction).then( res => {
            this.props.history.push('/correctiveActions');
        });
    }

    changecapaNumberHandler= (event) => {
        this.setState({capaNumber: event.target.value});
    }
    changerootCauseHandler= (event) => {
        this.setState({rootCause: event.target.value});
    }
    changecorrectiveActionHandler= (event) => {
        this.setState({correctiveAction: event.target.value});
    }
    changeverificationDateHandler= (event) => {
        this.setState({verificationDate: event.target.value});
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
                                            <label> capaNumber: </label>
                                                <input placeholder="capaNumber" name="capaNumber" className="form-control" value={this.state.capaNumber} onChange={this.changecapaNumberHandler}/>

                                            <label> rootCause: </label>
                                                <input placeholder="rootCause" name="rootCause" className="form-control" value={this.state.rootCause} onChange={this.changerootCauseHandler}/>

                                            <label> correctiveAction: </label>
                                                <input placeholder="correctiveAction" name="correctiveAction" className="form-control" value={this.state.correctiveAction} onChange={this.changecorrectiveActionHandler}/>

                                            <label> verificationDate: </label>
                                                <input type="date" placeholder="verificationDate" name="verificationDate" className="form-control" value={this.state.verificationDate} onChange={this.changeverificationDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Proposed
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Implemented
                      </option>
                      <option name="Status" className="form-control" >
                          Verified
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
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
