import React, { Component } from 'react'
import GiftCardRedemptionService from '../services/GiftCardRedemptionService';

class UpdateGiftCardRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                redeemedAt: '',
                amount: ''
        }
        this.updateGiftCardRedemption = this.updateGiftCardRedemption.bind(this);

        this.changeredeemedAtHandler = this.changeredeemedAtHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
    }

    componentDidMount(){
        GiftCardRedemptionService.getGiftCardRedemptionById(this.state.id).then( (res) =>{
            let giftCardRedemption = res.data;
            this.setState({
                redeemedAt: giftCardRedemption.redeemedAt,
                amount: giftCardRedemption.amount
            });
        });
    }

    updateGiftCardRedemption = (e) => {
        e.preventDefault();
        let giftCardRedemption = {
            giftCardRedemptionId: this.state.id,
            redeemedAt: this.state.redeemedAt,
            amount: this.state.amount
        };
        console.log('giftCardRedemption => ' + JSON.stringify(giftCardRedemption));
        console.log('id => ' + JSON.stringify(this.state.id));
        GiftCardRedemptionService.updateGiftCardRedemption(giftCardRedemption).then( res => {
            this.props.history.push('/giftCardRedemptions');
        });
    }

    changeredeemedAtHandler= (event) => {
        this.setState({redeemedAt: event.target.value});
    }
    changeamountHandler= (event) => {
        this.setState({amount: event.target.value});
    }

    cancel(){
        this.props.history.push('/giftCardRedemptions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GiftCardRedemption</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> redeemedAt: </label>
                                                <input type="date" placeholder="redeemedAt" name="redeemedAt" className="form-control" value={this.state.redeemedAt} onChange={this.changeredeemedAtHandler}/>

                                            <label> amount: </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateGiftCardRedemption}>Save</button>
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

export default UpdateGiftCardRedemptionComponent
