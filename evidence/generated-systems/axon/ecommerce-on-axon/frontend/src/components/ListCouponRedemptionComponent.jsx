import React, { Component } from 'react'
import CouponRedemptionService from '../services/CouponRedemptionService'

class ListCouponRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                couponRedemptions: []
        }
        this.addCouponRedemption = this.addCouponRedemption.bind(this);
        this.editCouponRedemption = this.editCouponRedemption.bind(this);
        this.deleteCouponRedemption = this.deleteCouponRedemption.bind(this);
    }

    deleteCouponRedemption(id){
        CouponRedemptionService.deleteCouponRedemption(id).then( res => {
            this.setState({couponRedemptions: this.state.couponRedemptions.filter(couponRedemption => couponRedemption.couponRedemptionId !== id)});
        });
    }
    viewCouponRedemption(id){
        this.props.history.push(`/view-couponRedemption/${id}`);
    }
    editCouponRedemption(id){
        this.props.history.push(`/add-couponRedemption/${id}`);
    }

    componentDidMount(){
        CouponRedemptionService.getCouponRedemptions().then((res) => {
            this.setState({ couponRedemptions: res.data});
        });
    }

    addCouponRedemption(){
        this.props.history.push('/add-couponRedemption/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CouponRedemption List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCouponRedemption}> Add CouponRedemption</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RedeemedAt </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.couponRedemptions.map(
                                        couponRedemption => 
                                        <tr key = {couponRedemption.couponRedemptionId}>
                                             <td> { couponRedemption.redeemedAt } </td>
                                             <td>
                                                 <button onClick={ () => this.editCouponRedemption(couponRedemption.couponRedemptionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCouponRedemption(couponRedemption.couponRedemptionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCouponRedemption(couponRedemption.couponRedemptionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCouponRedemptionComponent
