import React, { Component } from 'react'
import CouponService from '../services/CouponService'

class ListCouponComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                coupons: []
        }
        this.addCoupon = this.addCoupon.bind(this);
        this.editCoupon = this.editCoupon.bind(this);
        this.deleteCoupon = this.deleteCoupon.bind(this);
    }

    deleteCoupon(id){
        CouponService.deleteCoupon(id).then( res => {
            this.setState({coupons: this.state.coupons.filter(coupon => coupon.couponId !== id)});
        });
    }
    viewCoupon(id){
        this.props.history.push(`/view-coupon/${id}`);
    }
    editCoupon(id){
        this.props.history.push(`/add-coupon/${id}`);
    }

    componentDidMount(){
        CouponService.getCoupons().then((res) => {
            this.setState({ coupons: res.data});
        });
    }

    addCoupon(){
        this.props.history.push('/add-coupon/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Coupon List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCoupon}> Add Coupon</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> UsageLimit </th>
                                    <th> PerCustomerLimit </th>
                                    <th> ExpirationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.coupons.map(
                                        coupon => 
                                        <tr key = {coupon.couponId}>
                                             <td> { coupon.code } </td>
                                             <td> { coupon.usageLimit } </td>
                                             <td> { coupon.perCustomerLimit } </td>
                                             <td> { coupon.expirationDate } </td>
                                             <td> { coupon.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCoupon(coupon.couponId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCoupon(coupon.couponId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCoupon(coupon.couponId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListCouponComponent
