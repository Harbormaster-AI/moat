import React, { Component } from 'react'
import GiftCardRedemptionService from '../services/GiftCardRedemptionService'

class ViewGiftCardRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            giftCardRedemption: {}
        }
    }

    componentDidMount(){
        GiftCardRedemptionService.getGiftCardRedemptionById(this.state.id).then( res => {
            this.setState({giftCardRedemption: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View GiftCardRedemption Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> redeemedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCardRedemption.redeemedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCardRedemption.amount }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGiftCardRedemptionComponent
