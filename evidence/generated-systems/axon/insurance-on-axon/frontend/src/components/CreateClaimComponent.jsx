import React, { Component } from 'react'
import ClaimService from '../services/ClaimService';

class CreateClaimComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                claimNumber: '',
                noticeDate: '',
                lossDate: '',
                reportedBy: '',
                status: '',
                lossCause: ''
        }
        this.changeclaimNumberHandler = this.changeclaimNumberHandler.bind(this);
        this.changenoticeDateHandler = this.changenoticeDateHandler.bind(this);
        this.changelossDateHandler = this.changelossDateHandler.bind(this);
        this.changereportedByHandler = this.changereportedByHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeLossCauseHandler = this.changeLossCauseHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ClaimService.getClaimById(this.state.id).then( (res) =>{
                let claim = res.data;
                this.setState({
                    claimNumber: claim.claimNumber,
                    noticeDate: claim.noticeDate,
                    lossDate: claim.lossDate,
                    reportedBy: claim.reportedBy,
                    status: claim.status,
                    lossCause: claim.lossCause
                });
            });
        }        
    }
    saveOrUpdateClaim = (e) => {
        e.preventDefault();
        let claim = {
                claimId: this.state.id,
                claimNumber: this.state.claimNumber,
                noticeDate: this.state.noticeDate,
                lossDate: this.state.lossDate,
                reportedBy: this.state.reportedBy,
                status: this.state.status,
                lossCause: this.state.lossCause
            };
        console.log('claim => ' + JSON.stringify(claim));

        // step 5
        if(this.state.id === '_add'){
            claim.claimId=''
            ClaimService.createClaim(claim).then(res =>{
                this.props.history.push('/claims');
            });
        }else{
            ClaimService.updateClaim(claim).then( res => {
                this.props.history.push('/claims');
            });
        }
    }
    
    changeclaimNumberHandler= (event) => {
        this.setState({claimNumber: event.target.value});
    }
    changenoticeDateHandler= (event) => {
        this.setState({noticeDate: event.target.value});
    }
    changelossDateHandler= (event) => {
        this.setState({lossDate: event.target.value});
    }
    changereportedByHandler= (event) => {
        this.setState({reportedBy: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeLossCauseHandler= (event) => {
        this.setState({lossCause: event.target.value});
    }

    cancel(){
        this.props.history.push('/claims');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Claim</h3>
        }else{
            return <h3 className="text-center">Update Claim</h3>
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
                                            <label> claimNumber:&emsp; </label>
                                                <input placeholder="claimNumber" name="claimNumber" className="form-control" value={this.state.claimNumber} onChange={this.changeclaimNumberHandler}/>

                                            <label> noticeDate:&emsp; </label>
                                                <input type="date" placeholder="noticeDate" name="noticeDate" className="form-control" value={this.state.noticeDate} onChange={this.changenoticeDateHandler}/>

                                            <label> lossDate:&emsp; </label>
                                                <input type="date" placeholder="lossDate" name="lossDate" className="form-control" value={this.state.lossDate} onChange={this.changelossDateHandler}/>

                                            <label> reportedBy:&emsp; </label>
                                                <input placeholder="reportedBy" name="reportedBy" className="form-control" value={this.state.reportedBy} onChange={this.changereportedByHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Reopened
                      </option>
                      <option name="Status" className="form-control" >
                          Denied
                      </option>
                      <option name="Status" className="form-control" >
                          PendingInvestigation
                      </option>
                      <option name="Status" className="form-control" >
                          Litigation
                      </option>
                    </select>

                                            <label> LossCause:&emsp; </label>
                                                <select value={this.state.lossCause} onChange={this.changeLossCauseHandler}>
                      <option name="LossCause" className="form-control" >
                          Collision
                      </option>
                      <option name="LossCause" className="form-control" >
                          Weather
                      </option>
                      <option name="LossCause" className="form-control" >
                          MechanicalFailure
                      </option>
                      <option name="LossCause" className="form-control" >
                          HumanError
                      </option>
                      <option name="LossCause" className="form-control" >
                          NaturalDisaster
                      </option>
                      <option name="LossCause" className="form-control" >
                          Theft
                      </option>
                      <option name="LossCause" className="form-control" >
                          Vandalism
                      </option>
                      <option name="LossCause" className="form-control" >
                          LiabilityClaim
                      </option>
                      <option name="LossCause" className="form-control" >
                          Illness
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateClaim}>Save</button>
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

export default CreateClaimComponent
