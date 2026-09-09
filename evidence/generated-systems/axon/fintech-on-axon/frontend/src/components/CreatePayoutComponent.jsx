import React, { Component } from 'react'
import PayoutService from '../services/PayoutService';

class CreatePayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                payoutReference: '',
                amount: '',
                currency: '',
                scheduledDate: '',
                paidDate: '',
                status: ''
        }
        this.changepayoutReferenceHandler = this.changepayoutReferenceHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changescheduledDateHandler = this.changescheduledDateHandler.bind(this);
        this.changepaidDateHandler = this.changepaidDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PayoutService.getPayoutById(this.state.id).then( (res) =>{
                let payout = res.data;
                this.setState({
                    payoutReference: payout.payoutReference,
                    amount: payout.amount,
                    currency: payout.currency,
                    scheduledDate: payout.scheduledDate,
                    paidDate: payout.paidDate,
                    status: payout.status
                });
            });
        }        
    }
    saveOrUpdatePayout = (e) => {
        e.preventDefault();
        let payout = {
                payoutId: this.state.id,
                payoutReference: this.state.payoutReference,
                amount: this.state.amount,
                currency: this.state.currency,
                scheduledDate: this.state.scheduledDate,
                paidDate: this.state.paidDate,
                status: this.state.status
            };
        console.log('payout => ' + JSON.stringify(payout));

        // step 5
        if(this.state.id === '_add'){
            payout.payoutId=''
            PayoutService.createPayout(payout).then(res =>{
                this.props.history.push('/payouts');
            });
        }else{
            PayoutService.updatePayout(payout).then( res => {
                this.props.history.push('/payouts');
            });
        }
    }
    
    changepayoutReferenceHandler= (event) => {
        this.setState({payoutReference: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changescheduledDateHandler= (event) => {
        this.setState({scheduledDate: event.target.value});
    }
    changepaidDateHandler= (event) => {
        this.setState({paidDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/payouts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Payout</h3>
        }else{
            return <h3 className="text-center">Update Payout</h3>
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
                                            <label> payoutReference:&emsp; </label>
                                                <input placeholder="payoutReference" name="payoutReference" className="form-control" value={this.state.payoutReference} onChange={this.changepayoutReferenceHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> scheduledDate:&emsp; </label>
                                                <input type="date" placeholder="scheduledDate" name="scheduledDate" className="form-control" value={this.state.scheduledDate} onChange={this.changescheduledDateHandler}/>

                                            <label> paidDate:&emsp; </label>
                                                <input type="date" placeholder="paidDate" name="paidDate" className="form-control" value={this.state.paidDate} onChange={this.changepaidDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Scheduled
                      </option>
                      <option name="Status" className="form-control" >
                          Processing
                      </option>
                      <option name="Status" className="form-control" >
                          Paid
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePayout}>Save</button>
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

export default CreatePayoutComponent
