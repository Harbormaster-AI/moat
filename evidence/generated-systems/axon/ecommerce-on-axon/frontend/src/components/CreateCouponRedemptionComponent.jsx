import React, { Component } from 'react'
import CouponRedemptionService from '../services/CouponRedemptionService';

class CreateCouponRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                redeemedAt: ''
        }
        this.changeredeemedAtHandler = this.changeredeemedAtHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CouponRedemptionService.getCouponRedemptionById(this.state.id).then( (res) =>{
                let couponRedemption = res.data;
                this.setState({
                    redeemedAt: couponRedemption.redeemedAt
                });
            });
        }        
    }
    saveOrUpdateCouponRedemption = (e) => {
        e.preventDefault();
        let couponRedemption = {
                couponRedemptionId: this.state.id,
                redeemedAt: this.state.redeemedAt
            };
        console.log('couponRedemption => ' + JSON.stringify(couponRedemption));

        // step 5
        if(this.state.id === '_add'){
            couponRedemption.couponRedemptionId=''
            CouponRedemptionService.createCouponRedemption(couponRedemption).then(res =>{
                this.props.history.push('/couponRedemptions');
            });
        }else{
            CouponRedemptionService.updateCouponRedemption(couponRedemption).then( res => {
                this.props.history.push('/couponRedemptions');
            });
        }
    }
    
    changeredeemedAtHandler= (event) => {
        this.setState({redeemedAt: event.target.value});
    }

    cancel(){
        this.props.history.push('/couponRedemptions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CouponRedemption</h3>
        }else{
            return <h3 className="text-center">Update CouponRedemption</h3>
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

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCouponRedemption}>Save</button>
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

export default CreateCouponRedemptionComponent
