import React, { Component } from 'react'
import OfferService from '../services/OfferService'

class ViewOfferComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            offer: {}
        }
    }

    componentDidMount(){
        OfferService.getOfferById(this.state.id).then( res => {
            this.setState({offer: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Offer Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> offerNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.offer.offerNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> proposedStartDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.offer.proposedStartDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> baseSalary:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.offer.baseSalary }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> signOnBonus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.offer.signOnBonus }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.offer.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOfferComponent
