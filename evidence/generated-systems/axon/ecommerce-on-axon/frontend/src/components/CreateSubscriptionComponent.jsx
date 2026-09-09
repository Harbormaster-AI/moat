import React, { Component } from 'react'
import SubscriptionService from '../services/SubscriptionService';

class CreateSubscriptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                subscriptionNumber: '',
                nextBillingDate: '',
                startDate: '',
                endDate: '',
                status: '',
                interval: ''
        }
        this.changesubscriptionNumberHandler = this.changesubscriptionNumberHandler.bind(this);
        this.changenextBillingDateHandler = this.changenextBillingDateHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeIntervalHandler = this.changeIntervalHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SubscriptionService.getSubscriptionById(this.state.id).then( (res) =>{
                let subscription = res.data;
                this.setState({
                    subscriptionNumber: subscription.subscriptionNumber,
                    nextBillingDate: subscription.nextBillingDate,
                    startDate: subscription.startDate,
                    endDate: subscription.endDate,
                    status: subscription.status,
                    interval: subscription.interval
                });
            });
        }        
    }
    saveOrUpdateSubscription = (e) => {
        e.preventDefault();
        let subscription = {
                subscriptionId: this.state.id,
                subscriptionNumber: this.state.subscriptionNumber,
                nextBillingDate: this.state.nextBillingDate,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                status: this.state.status,
                interval: this.state.interval
            };
        console.log('subscription => ' + JSON.stringify(subscription));

        // step 5
        if(this.state.id === '_add'){
            subscription.subscriptionId=''
            SubscriptionService.createSubscription(subscription).then(res =>{
                this.props.history.push('/subscriptions');
            });
        }else{
            SubscriptionService.updateSubscription(subscription).then( res => {
                this.props.history.push('/subscriptions');
            });
        }
    }
    
    changesubscriptionNumberHandler= (event) => {
        this.setState({subscriptionNumber: event.target.value});
    }
    changenextBillingDateHandler= (event) => {
        this.setState({nextBillingDate: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeIntervalHandler= (event) => {
        this.setState({interval: event.target.value});
    }

    cancel(){
        this.props.history.push('/subscriptions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Subscription</h3>
        }else{
            return <h3 className="text-center">Update Subscription</h3>
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
                                            <label> subscriptionNumber:&emsp; </label>
                                                <input placeholder="subscriptionNumber" name="subscriptionNumber" className="form-control" value={this.state.subscriptionNumber} onChange={this.changesubscriptionNumberHandler}/>

                                            <label> nextBillingDate:&emsp; </label>
                                                <input type="date" placeholder="nextBillingDate" name="nextBillingDate" className="form-control" value={this.state.nextBillingDate} onChange={this.changenextBillingDateHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Paused
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                            <label> Interval:&emsp; </label>
                                                <select value={this.state.interval} onChange={this.changeIntervalHandler}>
                      <option name="Interval" className="form-control" >
                          Weekly
                      </option>
                      <option name="Interval" className="form-control" >
                          BiWeekly
                      </option>
                      <option name="Interval" className="form-control" >
                          Monthly
                      </option>
                      <option name="Interval" className="form-control" >
                          Quarterly
                      </option>
                      <option name="Interval" className="form-control" >
                          SemiAnnual
                      </option>
                      <option name="Interval" className="form-control" >
                          Annual
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSubscription}>Save</button>
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

export default CreateSubscriptionComponent
