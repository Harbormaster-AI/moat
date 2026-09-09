import React, { Component } from 'react'
import GiftCardService from '../services/GiftCardService'

class ViewGiftCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            giftCard: {}
        }
    }

    componentDidMount(){
        GiftCardService.getGiftCardById(this.state.id).then( res => {
            this.setState({giftCard: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View GiftCard Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCard.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> balance:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCard.balance }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expirationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCard.expirationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.giftCard.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGiftCardComponent
