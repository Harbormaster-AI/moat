import React, { Component } from 'react'
import LineItemService from '../services/LineItemService';

class CreateLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
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
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changebidAmountHandler = this.changebidAmountHandler.bind(this);
        this.changedailyBudgetHandler = this.changedailyBudgetHandler.bind(this);
        this.changefrequencyCapHandler = this.changefrequencyCapHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePricingModelHandler = this.changePricingModelHandler.bind(this);
        this.changeBidStrategyHandler = this.changeBidStrategyHandler.bind(this);
        this.changePacingHandler = this.changePacingHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateLineItem = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            lineItem.lineItemId=''
            LineItemService.createLineItem(lineItem).then(res =>{
                this.props.history.push('/lineItems');
            });
        }else{
            LineItemService.updateLineItem(lineItem).then( res => {
                this.props.history.push('/lineItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LineItem</h3>
        }else{
            return <h3 className="text-center">Update LineItem</h3>
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

                                            <label> bidAmount:&emsp; </label>
                                                <input placeholder="bidAmount" name="bidAmount" className="form-control" value={this.state.bidAmount} onChange={this.changebidAmountHandler}/>

                                            <label> dailyBudget:&emsp; </label>
                                                <input placeholder="dailyBudget" name="dailyBudget" className="form-control" value={this.state.dailyBudget} onChange={this.changedailyBudgetHandler}/>

                                            <label> frequencyCap:&emsp; </label>
                                                <input placeholder="frequencyCap" name="frequencyCap" className="form-control" value={this.state.frequencyCap} onChange={this.changefrequencyCapHandler}/>

                                            <label> Status:&emsp; </label>
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

                                            <label> PricingModel:&emsp; </label>
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

                                            <label> BidStrategy:&emsp; </label>
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

                                            <label> Pacing:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLineItem}>Save</button>
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

export default CreateLineItemComponent
