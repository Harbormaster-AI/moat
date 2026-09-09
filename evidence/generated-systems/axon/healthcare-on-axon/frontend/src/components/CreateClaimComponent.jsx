import React, { Component } from 'react'
import ClaimService from '../services/ClaimService';

class CreateClaimComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                claimNumber: '',
                totalAmount: '',
                status: ''
        }
        this.changeclaimNumberHandler = this.changeclaimNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
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
                    totalAmount: claim.totalAmount,
                    status: claim.status
                });
            });
        }        
    }
    saveOrUpdateClaim = (e) => {
        e.preventDefault();
        let claim = {
                claimId: this.state.id,
                claimNumber: this.state.claimNumber,
                totalAmount: this.state.totalAmount,
                status: this.state.status
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
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                            <label> Status:&emsp; </label>
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
