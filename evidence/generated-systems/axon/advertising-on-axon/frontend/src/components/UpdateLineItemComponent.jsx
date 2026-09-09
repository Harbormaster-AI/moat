import React, { Component } from 'react'
import LineItemService from '../services/LineItemService';

class UpdateLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                bidAmount: '',
                dailyBudget: '',
                frequencyCap: '',
                status: '',
                pricingModel: '',
                bidStrategy: '',
                pacing: ''
        }
        this.updateLineItem = this.updateLineItem.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changebidAmountHandler = this.changebidAmountHandler.bind(this);
        this.changedailyBudgetHandler = this.changedailyBudgetHandler.bind(this);
        this.changefrequencyCapHandler = this.changefrequencyCapHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePricingModelHandler = this.changePricingModelHandler.bind(this);
        this.changeBidStrategyHandler = this.changeBidStrategyHandler.bind(this);
        this.changePacingHandler = this.changePacingHandler.bind(this);
    }

    componentDidMount(){
        LineItemService.getLineItemById(this.state.id).then( (res) =>{
            let lineItem = res.data;
            this.setState({
                name: lineItem.name,
                bidAmount: lineItem.bidAmount,
                dailyBudget: lineItem.dailyBudget,
                frequencyCap: lineItem.frequencyCap,
                status: lineItem.status,
                pricingModel: lineItem.pricingModel,
                bidStrategy: lineItem.bidStrategy,
                pacing: lineItem.pacing
            });
        });
    }

    updateLineItem = (e) => {
        e.preventDefault();
        let lineItem = {
            lineItemId: this.state.id,
            name: this.state.name,
            bidAmount: this.state.bidAmount,
            dailyBudget: this.state.dailyBudget,
            frequencyCap: this.state.frequencyCap,
            status: this.state.status,
            pricingModel: this.state.pricingModel,
            bidStrategy: this.state.bidStrategy,
            pacing: this.state.pacing
        };
        console.log('lineItem => ' + JSON.stringify(lineItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        LineItemService.updateLineItem(lineItem).then( res => {
            this.props.history.push('/lineItems');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changebidAmountHandler= (event) => {
        this.setState({bidAmount: event.target.value});
    }
    changedailyBudgetHandler= (event) => {
        this.setState({dailyBudget: event.target.value});
    }
    changefrequencyCapHandler= (event) => {
        this.setState({frequencyCap: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePricingModelHandler= (event) => {
        this.setState({pricingModel: event.target.value});
    }
    changeBidStrategyHandler= (event) => {
        this.setState({bidStrategy: event.target.value});
    }
    changePacingHandler= (event) => {
        this.setState({pacing: event.target.value});
    }

    cancel(){
        this.props.history.push('/lineItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LineItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> bidAmount: </label>
                                                <input placeholder="bidAmount" name="bidAmount" className="form-control" value={this.state.bidAmount} onChange={this.changebidAmountHandler}/>

                                            <label> dailyBudget: </label>
                                                <input placeholder="dailyBudget" name="dailyBudget" className="form-control" value={this.state.dailyBudget} onChange={this.changedailyBudgetHandler}/>

                                            <label> frequencyCap: </label>
                                                <input placeholder="frequencyCap" name="frequencyCap" className="form-control" value={this.state.frequencyCap} onChange={this.changefrequencyCapHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Scheduled
                      </option>
                      <option name="Status" className="form-control" >
                          Running
                      </option>
                      <option name="Status" className="form-control" >
                          Paused
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> PricingModel: </label>
                                                <select value={this.state.pricingModel} onChange={this.changePricingModelHandler}>
                      <option name="PricingModel" className="form-control" >
                          CPM
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPC
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPA
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPL
                      </option>
                      <option name="PricingModel" className="form-control" >
                          CPV
                      </option>
                      <option name="PricingModel" className="form-control" >
                          FlatFee
                      </option>
                    </select>

                                            <label> BidStrategy: </label>
                                                <select value={this.state.bidStrategy} onChange={this.changeBidStrategyHandler}>
                      <option name="BidStrategy" className="form-control" >
                          Manual
                      </option>
                      <option name="BidStrategy" className="form-control" >
                          AutoMaximizeClicks
                      </option>
                      <option name="BidStrategy" className="form-control" >
                          AutoTargetCPA
                      </option>
                      <option name="BidStrategy" className="form-control" >
                          AutoTargetROAS
                      </option>
                    </select>

                                            <label> Pacing: </label>
                                                <select value={this.state.pacing} onChange={this.changePacingHandler}>
                      <option name="Pacing" className="form-control" >
                          Even
                      </option>
                      <option name="Pacing" className="form-control" >
                          ASAP
                      </option>
                      <option name="Pacing" className="form-control" >
                          Smooth
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLineItem}>Save</button>
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

export default UpdateLineItemComponent
