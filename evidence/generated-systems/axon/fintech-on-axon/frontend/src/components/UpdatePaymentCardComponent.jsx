import React, { Component } from 'react'
import PaymentCardService from '../services/PaymentCardService';

class UpdatePaymentCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                cardToken: '',
                maskedPan: '',
                expiryMonth: '',
                expiryYear: '',
                cardholderName: '',
                scheme: '',
                status: ''
        }
        this.updatePaymentCard = this.updatePaymentCard.bind(this);

        this.changecardTokenHandler = this.changecardTokenHandler.bind(this);
        this.changemaskedPanHandler = this.changemaskedPanHandler.bind(this);
        this.changeexpiryMonthHandler = this.changeexpiryMonthHandler.bind(this);
        this.changeexpiryYearHandler = this.changeexpiryYearHandler.bind(this);
        this.changecardholderNameHandler = this.changecardholderNameHandler.bind(this);
        this.changeSchemeHandler = this.changeSchemeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PaymentCardService.getPaymentCardById(this.state.id).then( (res) =>{
            let paymentCard = res.data;
            this.setState({
                cardToken: paymentCard.cardToken,
                maskedPan: paymentCard.maskedPan,
                expiryMonth: paymentCard.expiryMonth,
                expiryYear: paymentCard.expiryYear,
                cardholderName: paymentCard.cardholderName,
                scheme: paymentCard.scheme,
                status: paymentCard.status
            });
        });
    }

    updatePaymentCard = (e) => {
        e.preventDefault();
        let paymentCard = {
            paymentCardId: this.state.id,
            cardToken: this.state.cardToken,
            maskedPan: this.state.maskedPan,
            expiryMonth: this.state.expiryMonth,
            expiryYear: this.state.expiryYear,
            cardholderName: this.state.cardholderName,
            scheme: this.state.scheme,
            status: this.state.status
        };
        console.log('paymentCard => ' + JSON.stringify(paymentCard));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentCardService.updatePaymentCard(paymentCard).then( res => {
            this.props.history.push('/paymentCards');
        });
    }

    changecardTokenHandler= (event) => {
        this.setState({cardToken: event.target.value});
    }
    changemaskedPanHandler= (event) => {
        this.setState({maskedPan: event.target.value});
    }
    changeexpiryMonthHandler= (event) => {
        this.setState({expiryMonth: event.target.value});
    }
    changeexpiryYearHandler= (event) => {
        this.setState({expiryYear: event.target.value});
    }
    changecardholderNameHandler= (event) => {
        this.setState({cardholderName: event.target.value});
    }
    changeSchemeHandler= (event) => {
        this.setState({scheme: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentCards');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentCard</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> cardToken: </label>
                                                <input placeholder="cardToken" name="cardToken" className="form-control" value={this.state.cardToken} onChange={this.changecardTokenHandler}/>

                                            <label> maskedPan: </label>
                                                <input placeholder="maskedPan" name="maskedPan" className="form-control" value={this.state.maskedPan} onChange={this.changemaskedPanHandler}/>

                                            <label> expiryMonth: </label>
                                                <input type="number" placeholder="expiryMonth" name="expiryMonth" className="form-control" value={this.state.expiryMonth} onChange={this.changeexpiryMonthHandler}/>

                                            <label> expiryYear: </label>
                                                <input type="number" placeholder="expiryYear" name="expiryYear" className="form-control" value={this.state.expiryYear} onChange={this.changeexpiryYearHandler}/>

                                            <label> cardholderName: </label>
                                                <input placeholder="cardholderName" name="cardholderName" className="form-control" value={this.state.cardholderName} onChange={this.changecardholderNameHandler}/>

                                            <label> Scheme: </label>
                                                <select value={this.state.scheme} onChange={this.changeSchemeHandler}>
                      <option name="Scheme" className="form-control" >
                          Visa
                      </option>
                      <option name="Scheme" className="form-control" >
                          Mastercard
                      </option>
                      <option name="Scheme" className="form-control" >
                          Amex
                      </option>
                      <option name="Scheme" className="form-control" >
                          Discover
                      </option>
                      <option name="Scheme" className="form-control" >
                          UnionPay
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Blocked
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentCard}>Save</button>
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

export default UpdatePaymentCardComponent
