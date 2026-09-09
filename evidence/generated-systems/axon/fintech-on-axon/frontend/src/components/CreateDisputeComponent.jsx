import React, { Component } from 'react'
import DisputeService from '../services/DisputeService';

class CreateDisputeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                disputeReference: '',
                openedAt: '',
                closedAt: '',
                reason: '',
                status: ''
        }
        this.changedisputeReferenceHandler = this.changedisputeReferenceHandler.bind(this);
        this.changeopenedAtHandler = this.changeopenedAtHandler.bind(this);
        this.changeclosedAtHandler = this.changeclosedAtHandler.bind(this);
        this.changeReasonHandler = this.changeReasonHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DisputeService.getDisputeById(this.state.id).then( (res) =>{
                let dispute = res.data;
                this.setState({
                    disputeReference: dispute.disputeReference,
                    openedAt: dispute.openedAt,
                    closedAt: dispute.closedAt,
                    reason: dispute.reason,
                    status: dispute.status
                });
            });
        }        
    }
    saveOrUpdateDispute = (e) => {
        e.preventDefault();
        let dispute = {
                disputeId: this.state.id,
                disputeReference: this.state.disputeReference,
                openedAt: this.state.openedAt,
                closedAt: this.state.closedAt,
                reason: this.state.reason,
                status: this.state.status
            };
        console.log('dispute => ' + JSON.stringify(dispute));

        // step 5
        if(this.state.id === '_add'){
            dispute.disputeId=''
            DisputeService.createDispute(dispute).then(res =>{
                this.props.history.push('/disputes');
            });
        }else{
            DisputeService.updateDispute(dispute).then( res => {
                this.props.history.push('/disputes');
            });
        }
    }
    
    changedisputeReferenceHandler= (event) => {
        this.setState({disputeReference: event.target.value});
    }
    changeopenedAtHandler= (event) => {
        this.setState({openedAt: event.target.value});
    }
    changeclosedAtHandler= (event) => {
        this.setState({closedAt: event.target.value});
    }
    changeReasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/disputes');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Dispute</h3>
        }else{
            return <h3 className="text-center">Update Dispute</h3>
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
                                            <label> disputeReference:&emsp; </label>
                                                <input placeholder="disputeReference" name="disputeReference" className="form-control" value={this.state.disputeReference} onChange={this.changedisputeReferenceHandler}/>

                                            <label> openedAt:&emsp; </label>
                                                <input type="time" placeholder="openedAt" name="openedAt" className="form-control" value={this.state.openedAt} onChange={this.changeopenedAtHandler}/>

                                            <label> closedAt:&emsp; </label>
                                                <input type="time" placeholder="closedAt" name="closedAt" className="form-control" value={this.state.closedAt} onChange={this.changeclosedAtHandler}/>

                                            <label> Reason:&emsp; </label>
                                                <select value={this.state.reason} onChange={this.changeReasonHandler}>
                      <option name="Reason" className="form-control" >
                          Fraud
                      </option>
                      <option name="Reason" className="form-control" >
                          Duplicate
                      </option>
                      <option name="Reason" className="form-control" >
                          NotAsDescribed
                      </option>
                      <option name="Reason" className="form-control" >
                          NotReceived
                      </option>
                      <option name="Reason" className="form-control" >
                          ProcessingError
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Represented
                      </option>
                      <option name="Status" className="form-control" >
                          Won
                      </option>
                      <option name="Status" className="form-control" >
                          Lost
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDispute}>Save</button>
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

export default CreateDisputeComponent
