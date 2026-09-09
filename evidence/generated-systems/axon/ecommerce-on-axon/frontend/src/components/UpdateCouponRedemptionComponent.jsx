import React, { Component } from 'react'
import CouponRedemptionService from '../services/CouponRedemptionService';

class UpdateCouponRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                redeemedAt: ''
        }
        this.updateCouponRedemption = this.updateCouponRedemption.bind(this);

        this.changeredeemedAtHandler = this.changeredeemedAtHandler.bind(this);
    }

    componentDidMount(){
        CouponRedemptionService.getCouponRedemptionById(this.state.id).then( (res) =>{
            let couponRedemption = res.data;
            this.setState({
                redeemedAt: couponRedemption.redeemedAt
            });
        });
    }

    updateCouponRedemption = (e) => {
        e.preventDefault();
        let couponRedemption = {
            couponRedemptionId: this.state.id,
            redeemedAt: this.state.redeemedAt
        };
        console.log('couponRedemption => ' + JSON.stringify(couponRedemption));
        console.log('id => ' + JSON.stringify(this.state.id));
        CouponRedemptionService.updateCouponRedemption(couponRedemption).then( res => {
            this.props.history.push('/couponRedemptions');
        });
    }

    changeredeemedAtHandler= (event) => {
        this.setState({redeemedAt: event.target.value});
    }

    cancel(){
        this.props.history.push('/couponRedemptions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CouponRedemption</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> redeemedAt: </label>
                                                <input type="date" placeholder="redeemedAt" name="redeemedAt" className="form-control" value={this.state.redeemedAt} onChange={this.changeredeemedAtHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCouponRedemption}>Save</button>
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

export default UpdateCouponRedemptionComponent
