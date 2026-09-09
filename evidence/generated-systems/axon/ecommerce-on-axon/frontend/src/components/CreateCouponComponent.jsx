import React, { Component } from 'react'
import CouponService from '../services/CouponService';

class CreateCouponComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                usageLimit: '',
                perCustomerLimit: '',
                expirationDate: '',
                status: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeusageLimitHandler = this.changeusageLimitHandler.bind(this);
        this.changeperCustomerLimitHandler = this.changeperCustomerLimitHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CouponService.getCouponById(this.state.id).then( (res) =>{
                let coupon = res.data;
                this.setState({
                    code: coupon.code,
                    usageLimit: coupon.usageLimit,
                    perCustomerLimit: coupon.perCustomerLimit,
                    expirationDate: coupon.expirationDate,
                    status: coupon.status
                });
            });
        }        
    }
    saveOrUpdateCoupon = (e) => {
        e.preventDefault();
        let coupon = {
                couponId: this.state.id,
                code: this.state.code,
                usageLimit: this.state.usageLimit,
                perCustomerLimit: this.state.perCustomerLimit,
                expirationDate: this.state.expirationDate,
                status: this.state.status
            };
        console.log('coupon => ' + JSON.stringify(coupon));

        // step 5
        if(this.state.id === '_add'){
            coupon.couponId=''
            CouponService.createCoupon(coupon).then(res =>{
                this.props.history.push('/coupons');
            });
        }else{
            CouponService.updateCoupon(coupon).then( res => {
                this.props.history.push('/coupons');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeusageLimitHandler= (event) => {
        this.setState({usageLimit: event.target.value});
    }
    changeperCustomerLimitHandler= (event) => {
        this.setState({perCustomerLimit: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/coupons');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Coupon</h3>
        }else{
            return <h3 className="text-center">Update Coupon</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> usageLimit:&emsp; </label>
                                                <input type="number" placeholder="usageLimit" name="usageLimit" className="form-control" value={this.state.usageLimit} onChange={this.changeusageLimitHandler}/>

                                            <label> perCustomerLimit:&emsp; </label>
                                                <input type="number" placeholder="perCustomerLimit" name="perCustomerLimit" className="form-control" value={this.state.perCustomerLimit} onChange={this.changeperCustomerLimitHandler}/>

                                            <label> expirationDate:&emsp; </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Disabled
                      </option>
                      <option name="Status" className="form-control" >
                          Exhausted
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCoupon}>Save</button>
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

export default CreateCouponComponent
