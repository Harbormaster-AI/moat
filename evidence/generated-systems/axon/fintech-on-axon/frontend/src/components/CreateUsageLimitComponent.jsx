import React, { Component } from 'react'
import UsageLimitService from '../services/UsageLimitService';

class CreateUsageLimitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                amount: '',
                count: '',
                scope: '',
                period: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changecountHandler = this.changecountHandler.bind(this);
        this.changeScopeHandler = this.changeScopeHandler.bind(this);
        this.changePeriodHandler = this.changePeriodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            UsageLimitService.getUsageLimitById(this.state.id).then( (res) =>{
                let usageLimit = res.data;
                this.setState({
                    name: usageLimit.name,
                    amount: usageLimit.amount,
                    count: usageLimit.count,
                    scope: usageLimit.scope,
                    period: usageLimit.period
                });
            });
        }        
    }
    saveOrUpdateUsageLimit = (e) => {
        e.preventDefault();
        let usageLimit = {
                usageLimitId: this.state.id,
                name: this.state.name,
                amount: this.state.amount,
                count: this.state.count,
                scope: this.state.scope,
                period: this.state.period
            };
        console.log('usageLimit => ' + JSON.stringify(usageLimit));

        // step 5
        if(this.state.id === '_add'){
            usageLimit.usageLimitId=''
            UsageLimitService.createUsageLimit(usageLimit).then(res =>{
                this.props.history.push('/usageLimits');
            });
        }else{
            UsageLimitService.updateUsageLimit(usageLimit).then( res => {
                this.props.history.push('/usageLimits');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changecountHandler= (event) => {
        this.setState({count: event.target.value});
    }
    changeScopeHandler= (event) => {
        this.setState({scope: event.target.value});
    }
    changePeriodHandler= (event) => {
        this.setState({period: event.target.value});
    }

    cancel(){
        this.props.history.push('/usageLimits');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add UsageLimit</h3>
        }else{
            return <h3 className="text-center">Update UsageLimit</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> count:&emsp; </label>
                                                <input type="number" placeholder="count" name="count" className="form-control" value={this.state.count} onChange={this.changecountHandler}/>

                                            <label> Scope:&emsp; </label>
                                                <select value={this.state.scope} onChange={this.changeScopeHandler}>
                      <option name="Scope" className="form-control" >
                          PerTransaction
                      </option>
                      <option name="Scope" className="form-control" >
                          Daily
                      </option>
                      <option name="Scope" className="form-control" >
                          Monthly
                      </option>
                      <option name="Scope" className="form-control" >
                          Yearly
                      </option>
                      <option name="Scope" className="form-control" >
                          Rolling24h
                      </option>
                    </select>

                                            <label> Period:&emsp; </label>
                                                <select value={this.state.period} onChange={this.changePeriodHandler}>
                      <option name="Period" className="form-control" >
                          None
                      </option>
                      <option name="Period" className="form-control" >
                          Day
                      </option>
                      <option name="Period" className="form-control" >
                          Week
                      </option>
                      <option name="Period" className="form-control" >
                          Month
                      </option>
                      <option name="Period" className="form-control" >
                          Year
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateUsageLimit}>Save</button>
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

export default CreateUsageLimitComponent
