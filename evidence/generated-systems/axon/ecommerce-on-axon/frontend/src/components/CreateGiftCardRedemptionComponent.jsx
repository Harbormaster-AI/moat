import React, { Component } from 'react'
import GiftCardRedemptionService from '../services/GiftCardRedemptionService';

class CreateGiftCardRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                redeemedAt: '',
                amount: ''
        }
        this.changeredeemedAtHandler = this.changeredeemedAtHandler.bind(this);
        this.changeamountHandler = this.changeamountHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            GiftCardRedemptionService.getGiftCardRedemptionById(this.state.id).then( (res) =>{
                let giftCardRedemption = res.data;
                this.setState({
                    redeemedAt: giftCardRedemption.redeemedAt,
                    amount: giftCardRedemption.amount
                });
            });
        }        
    }
    saveOrUpdateGiftCardRedemption = (e) => {
        e.preventDefault();
        let giftCardRedemption = {
                giftCardRedemptionId: this.state.id,
                redeemedAt: this.state.redeemedAt,
                amount: this.state.amount
            };
        console.log('giftCardRedemption => ' + JSON.stringify(giftCardRedemption));

        // step 5
        if(this.state.id === '_add'){
            giftCardRedemption.giftCardRedemptionId=''
            GiftCardRedemptionService.createGiftCardRedemption(giftCardRedemption).then(res =>{
                this.props.history.push('/giftCardRedemptions');
            });
        }else{
            GiftCardRedemptionService.updateGiftCardRedemption(giftCardRedemption).then( res => {
                this.props.history.push('/giftCardRedemptions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add GiftCardRedemption</h3>
        }else{
            return <h3 className="text-center">Update GiftCardRedemption</h3>
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
                                            <label> redeemedAt:&emsp; </label>
                                                <input type="date" placeholder="redeemedAt" name="redeemedAt" className="form-control" value={this.state.redeemedAt} onChange={this.changeredeemedAtHandler}/>

                                            <label> amount:&emsp; </label>
                                                <input placeholder="amount" name="amount" className="form-control" value={this.state.amount} onChange={this.changeamountHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGiftCardRedemption}>Save</button>
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

export default CreateGiftCardRedemptionComponent
