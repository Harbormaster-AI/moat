import React, { Component } from 'react'
import ClaimService from '../services/ClaimService';

class UpdateClaimComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                claimNumber: '',
                totalAmount: '',
                status: ''
        }
        this.updateClaim = this.updateClaim.bind(this);

        this.changeclaimNumberHandler = this.changeclaimNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ClaimService.getClaimById(this.state.id).then( (res) =>{
            let claim = res.data;
            this.setState({
                claimNumber: claim.claimNumber,
                totalAmount: claim.totalAmount,
                status: claim.status
            });
        });
    }

    updateClaim = (e) => {
        e.preventDefault();
        let claim = {
            claimId: this.state.id,
            claimNumber: this.state.claimNumber,
            totalAmount: this.state.totalAmount,
            status: this.state.status
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
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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

                                            <label> totalAmount: </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          InProcess
                      </option>
                      <option name="Status" className="form-control" >
                          Paid
                      </option>
                      <option name="Status" className="form-control" >
                          Denied
                      </option>
                      <option name="Status" className="form-control" >
                          Adjusted
                      </option>
                      <option name="Status" className="form-control" >
                          Void
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
