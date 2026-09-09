import React, { Component } from 'react'
import ClaimService from '../services/ClaimService';

class UpdateClaimComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                claimNumber: '',
                noticeDate: '',
                lossDate: '',
                reportedBy: '',
                status: '',
                lossCause: ''
        }
        this.updateClaim = this.updateClaim.bind(this);

        this.changeclaimNumberHandler = this.changeclaimNumberHandler.bind(this);
        this.changenoticeDateHandler = this.changenoticeDateHandler.bind(this);
        this.changelossDateHandler = this.changelossDateHandler.bind(this);
        this.changereportedByHandler = this.changereportedByHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeLossCauseHandler = this.changeLossCauseHandler.bind(this);
    }

    componentDidMount(){
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

    updateClaim = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        ClaimService.updateClaim(claim).then( res => {
            this.props.history.push('/claims');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Claim</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> claimNumber: </label>
                                                <input placeholder="claimNumber" name="claimNumber" className="form-control" value={this.state.claimNumber} onChange={this.changeclaimNumberHandler}/>

                                            <label> noticeDate: </label>
                                                <input type="date" placeholder="noticeDate" name="noticeDate" className="form-control" value={this.state.noticeDate} onChange={this.changenoticeDateHandler}/>

                                            <label> lossDate: </label>
                                                <input type="date" placeholder="lossDate" name="lossDate" className="form-control" value={this.state.lossDate} onChange={this.changelossDateHandler}/>

                                            <label> reportedBy: </label>
                                                <input placeholder="reportedBy" name="reportedBy" className="form-control" value={this.state.reportedBy} onChange={this.changereportedByHandler}/>

                                            <label> Status: </label>
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

                                            <label> LossCause: </label>
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
                                        <button className="btn btn-success" onClick={this.updateClaim}>Save</button>
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

export default UpdateClaimComponent
