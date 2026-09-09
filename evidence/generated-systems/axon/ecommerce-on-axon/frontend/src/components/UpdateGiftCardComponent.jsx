import React, { Component } from 'react'
import GiftCardService from '../services/GiftCardService';

class UpdateGiftCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                balance: '',
                expirationDate: '',
                status: ''
        }
        this.updateGiftCard = this.updateGiftCard.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changebalanceHandler = this.changebalanceHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        GiftCardService.getGiftCardById(this.state.id).then( (res) =>{
            let giftCard = res.data;
            this.setState({
                code: giftCard.code,
                balance: giftCard.balance,
                expirationDate: giftCard.expirationDate,
                status: giftCard.status
            });
        });
    }

    updateGiftCard = (e) => {
        e.preventDefault();
        let giftCard = {
            giftCardId: this.state.id,
            code: this.state.code,
            balance: this.state.balance,
            expirationDate: this.state.expirationDate,
            status: this.state.status
        };
        console.log('giftCard => ' + JSON.stringify(giftCard));
        console.log('id => ' + JSON.stringify(this.state.id));
        GiftCardService.updateGiftCard(giftCard).then( res => {
            this.props.history.push('/giftCards');
        });
    }

    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changebalanceHandler= (event) => {
        this.setState({balance: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/giftCards');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GiftCard</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> balance: </label>
                                                <input placeholder="balance" name="balance" className="form-control" value={this.state.balance} onChange={this.changebalanceHandler}/>

                                            <label> expirationDate: </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Redeemed
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Disabled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateGiftCard}>Save</button>
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

export default UpdateGiftCardComponent
