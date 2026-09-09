import React, { Component } from 'react'
import CorrectiveActionService from '../services/CorrectiveActionService';

class CreateCorrectiveActionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                capaNumber: '',
                rootCause: '',
                correctiveAction: '',
                verificationDate: '',
                status: ''
        }
        this.changecapaNumberHandler = this.changecapaNumberHandler.bind(this);
        this.changerootCauseHandler = this.changerootCauseHandler.bind(this);
        this.changecorrectiveActionHandler = this.changecorrectiveActionHandler.bind(this);
        this.changeverificationDateHandler = this.changeverificationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCorrectiveAction = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            correctiveAction.correctiveActionId=''
            CorrectiveActionService.createCorrectiveAction(correctiveAction).then(res =>{
                this.props.history.push('/correctiveActions');
            });
        }else{
            CorrectiveActionService.updateCorrectiveAction(correctiveAction).then( res => {
                this.props.history.push('/correctiveActions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CorrectiveAction</h3>
        }else{
            return <h3 className="text-center">Update CorrectiveAction</h3>
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
                                            <label> capaNumber:&emsp; </label>
                                                <input placeholder="capaNumber" name="capaNumber" className="form-control" value={this.state.capaNumber} onChange={this.changecapaNumberHandler}/>

                                            <label> rootCause:&emsp; </label>
                                                <input placeholder="rootCause" name="rootCause" className="form-control" value={this.state.rootCause} onChange={this.changerootCauseHandler}/>

                                            <label> correctiveAction:&emsp; </label>
                                                <input placeholder="correctiveAction" name="correctiveAction" className="form-control" value={this.state.correctiveAction} onChange={this.changecorrectiveActionHandler}/>

                                            <label> verificationDate:&emsp; </label>
                                                <input type="date" placeholder="verificationDate" name="verificationDate" className="form-control" value={this.state.verificationDate} onChange={this.changeverificationDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCorrectiveAction}>Save</button>
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

export default CreateCorrectiveActionComponent
