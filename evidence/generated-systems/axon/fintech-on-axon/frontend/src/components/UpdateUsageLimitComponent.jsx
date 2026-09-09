import React, { Component } from 'react'
import UsageLimitService from '../services/UsageLimitService';

class UpdateUsageLimitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                amount: '',
                count: '',
                scope: '',
                period: ''
        }
        this.updateUsageLimit = this.updateUsageLimit.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changecountHandler = this.changecountHandler.bind(this);
        this.changeScopeHandler = this.changeScopeHandler.bind(this);
        this.changePeriodHandler = this.changePeriodHandler.bind(this);
    }

    componentDidMount(){
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

    updateUsageLimit = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        UsageLimitService.updateUsageLimit(usageLimit).then( res => {
            this.props.history.push('/usageLimits');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update UsageLimit</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> count: </label>
                                                <input type="number" placeholder="count" name="count" className="form-control" value={this.state.count} onChange={this.changecountHandler}/>

                                            <label> Scope: </label>
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

                                            <label> Period: </label>
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
                                        <button className="btn btn-success" onClick={this.updateUsageLimit}>Save</button>
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

export default UpdateUsageLimitComponent
