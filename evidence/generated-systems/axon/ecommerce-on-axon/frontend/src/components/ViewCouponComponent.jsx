import React, { Component } from 'react'
import CouponService from '../services/CouponService'

class ViewCouponComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            coupon: {}
        }
    }

    componentDidMount(){
        CouponService.getCouponById(this.state.id).then( res => {
            this.setState({coupon: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Coupon Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coupon.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> usageLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coupon.usageLimit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> perCustomerLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coupon.perCustomerLimit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> expirationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coupon.expirationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coupon.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCouponComponent
