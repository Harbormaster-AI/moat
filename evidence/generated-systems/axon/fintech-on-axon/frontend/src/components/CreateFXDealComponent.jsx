import React, { Component } from 'react'
import FXDealService from '../services/FXDealService';

class CreateFXDealComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                dealReference: '',
                baseCurrency: '',
                quoteCurrency: '',
                rate: '',
                amount: '',
                settlementDate: '',
                status: ''
        }
        this.changedealReferenceHandler = this.changedealReferenceHandler.bind(this);
        this.changebaseCurrencyHandler = this.changebaseCurrencyHandler.bind(this);
        this.changequoteCurrencyHandler = this.changequoteCurrencyHandler.bind(this);
        this.changerateHandler = this.changerateHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
        this.changesettlementDateHandler = this.changesettlementDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FXDealService.getFXDealById(this.state.id).then( (res) =>{
                let fXDeal = res.data;
                this.setState({
                    dealReference: fXDeal.dealReference,
                    baseCurrency: fXDeal.baseCurrency,
                    quoteCurrency: fXDeal.quoteCurrency,
                    rate: fXDeal.rate,
                    amount: fXDeal.amount,
                    settlementDate: fXDeal.settlementDate,
                    status: fXDeal.status
                });
            });
        }        
    }
    saveOrUpdateFXDeal = (e) => {
        e.preventDefault();
        let fXDeal = {
                fXDealId: this.state.id,
                dealReference: this.state.dealReference,
                baseCurrency: this.state.baseCurrency,
                quoteCurrency: this.state.quoteCurrency,
                rate: this.state.rate,
                amount: this.state.amount,
                settlementDate: this.state.settlementDate,
                status: this.state.status
            };
        console.log('fXDeal => ' + JSON.stringify(fXDeal));

        // step 5
        if(this.state.id === '_add'){
            fXDeal.fXDealId=''
            FXDealService.createFXDeal(fXDeal).then(res =>{
                this.props.history.push('/fXDeals');
            });
        }else{
            FXDealService.updateFXDeal(fXDeal).then( res => {
                this.props.history.push('/fXDeals');
            });
        }
    }
    
    changedealReferenceHandler= (event) => {
        this.setState({dealReference: event.target.value});
    }
    changebaseCurrencyHandler= (event) => {
        this.setState({baseCurrency: event.target.value});
    }
    changequoteCurrencyHandler= (event) => {
        this.setState({quoteCurrency: event.target.value});
    }
    changerateHandler= (event) => {
        this.setState({rate: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }
    changesettlementDateHandler= (event) => {
        this.setState({settlementDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/fXDeals');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FXDeal</h3>
        }else{
            return <h3 className="text-center">Update FXDeal</h3>
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
                                            <label> dealReference:&emsp; </label>
                                                <input placeholder="dealReference" name="dealReference" className="form-control" value={this.state.dealReference} onChange={this.changedealReferenceHandler}/>

                                            <label> baseCurrency:&emsp; </label>
                                                <input placeholder="baseCurrency" name="baseCurrency" className="form-control" value={this.state.baseCurrency} onChange={this.changebaseCurrencyHandler}/>

                                            <label> quoteCurrency:&emsp; </label>
                                                <input placeholder="quoteCurrency" name="quoteCurrency" className="form-control" value={this.state.quoteCurrency} onChange={this.changequoteCurrencyHandler}/>

                                            <label> rate:&emsp; </label>
                                                <input placeholder="rate" name="rate" className="form-control" value={this.state.rate} onChange={this.changerateHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                            <label> settlementDate:&emsp; </label>
                                                <input type="date" placeholder="settlementDate" name="settlementDate" className="form-control" value={this.state.settlementDate} onChange={this.changesettlementDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Booked
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Settled
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFXDeal}>Save</button>
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

export default CreateFXDealComponent
