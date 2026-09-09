import React, { Component } from 'react'
import CouponRedemptionService from '../services/CouponRedemptionService'

class ViewCouponRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            couponRedemption: {}
        }
    }

    componentDidMount(){
        CouponRedemptionService.getCouponRedemptionById(this.state.id).then( res => {
            this.setState({couponRedemption: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CouponRedemption Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> redeemedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.couponRedemption.redeemedAt }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCouponRedemptionComponent
