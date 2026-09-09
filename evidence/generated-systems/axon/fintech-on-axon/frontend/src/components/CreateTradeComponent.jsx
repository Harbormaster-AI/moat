import React, { Component } from 'react'
import TradeService from '../services/TradeService';

class CreateTradeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                executedAt: '',
                quantity: '',
                price: '',
                fees: '',
                settlementDate: ''
        }
        this.changeexecutedAtHandler = this.changeexecutedAtHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changepriceHandler = this.changepriceHandler.bind(this);
        this.changefeesHandler = this.changefeesHandler.bind(this);
        this.changesettlementDateHandler = this.changesettlementDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TradeService.getTradeById(this.state.id).then( (res) =>{
                let trade = res.data;
                this.setState({
                    executedAt: trade.executedAt,
                    quantity: trade.quantity,
                    price: trade.price,
                    fees: trade.fees,
                    settlementDate: trade.settlementDate
                });
            });
        }        
    }
    saveOrUpdateTrade = (e) => {
        e.preventDefault();
        let trade = {
                tradeId: this.state.id,
                executedAt: this.state.executedAt,
                quantity: this.state.quantity,
                price: this.state.price,
                fees: this.state.fees,
                settlementDate: this.state.settlementDate
            };
        console.log('trade => ' + JSON.stringify(trade));

        // step 5
        if(this.state.id === '_add'){
            trade.tradeId=''
            TradeService.createTrade(trade).then(res =>{
                this.props.history.push('/trades');
            });
        }else{
            TradeService.updateTrade(trade).then( res => {
                this.props.history.push('/trades');
            });
        }
    }
    
    changeexecutedAtHandler= (event) => {
        this.setState({executedAt: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changepriceHandler= (event) => {
        this.setState({price: event.target.value});
    }
    changefeesHandler= (event) => {
        this.setState({fees: event.target.value});
    }
    changesettlementDateHandler= (event) => {
        this.setState({settlementDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/trades');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Trade</h3>
        }else{
            return <h3 className="text-center">Update Trade</h3>
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
                                            <label> executedAt:&emsp; </label>
                                                <input type="time" placeholder="executedAt" name="executedAt" className="form-control" value={this.state.executedAt} onChange={this.changeexecutedAtHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> price:&emsp; </label>
                                                <input placeholder="price" name="price" className="form-control" value={this.state.price} onChange={this.changepriceHandler}/>

                                            <label> fees:&emsp; </label>
                                                <input placeholder="fees" name="fees" className="form-control" value={this.state.fees} onChange={this.changefeesHandler}/>

                                            <label> settlementDate:&emsp; </label>
                                                <input type="date" placeholder="settlementDate" name="settlementDate" className="form-control" value={this.state.settlementDate} onChange={this.changesettlementDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTrade}>Save</button>
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

export default CreateTradeComponent
