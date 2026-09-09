import React, { Component } from 'react'
import PromotionService from '../services/PromotionService';

class CreatePromotionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                code: '',
                value: '',
                startDate: '',
                endDate: '',
                asStackable: '',
                maxRedemptions: '',
                promotionType: '',
                discountType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeasStackableHandler = this.changeasStackableHandler.bind(this);
        this.changemaxRedemptionsHandler = this.changemaxRedemptionsHandler.bind(this);
        this.changePromotionTypeHandler = this.changePromotionTypeHandler.bind(this);
        this.changeDiscountTypeHandler = this.changeDiscountTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PromotionService.getPromotionById(this.state.id).then( (res) =>{
                let promotion = res.data;
                this.setState({
                    name: promotion.name,
                    code: promotion.code,
                    value: promotion.value,
                    startDate: promotion.startDate,
                    endDate: promotion.endDate,
                    asStackable: promotion.asStackable,
                    maxRedemptions: promotion.maxRedemptions,
                    promotionType: promotion.promotionType,
                    discountType: promotion.discountType
                });
            });
        }        
    }
    saveOrUpdatePromotion = (e) => {
        e.preventDefault();
        let promotion = {
                promotionId: this.state.id,
                name: this.state.name,
                code: this.state.code,
                value: this.state.value,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                asStackable: this.state.asStackable,
                maxRedemptions: this.state.maxRedemptions,
                promotionType: this.state.promotionType,
                discountType: this.state.discountType
            };
        console.log('promotion => ' + JSON.stringify(promotion));

        // step 5
        if(this.state.id === '_add'){
            promotion.promotionId=''
            PromotionService.createPromotion(promotion).then(res =>{
                this.props.history.push('/promotions');
            });
        }else{
            PromotionService.updatePromotion(promotion).then( res => {
                this.props.history.push('/promotions');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeasStackableHandler= (event) => {
        this.setState({asStackable: event.target.value});
    }
    changemaxRedemptionsHandler= (event) => {
        this.setState({maxRedemptions: event.target.value});
    }
    changePromotionTypeHandler= (event) => {
        this.setState({promotionType: event.target.value});
    }
    changeDiscountTypeHandler= (event) => {
        this.setState({discountType: event.target.value});
    }

    cancel(){
        this.props.history.push('/promotions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Promotion</h3>
        }else{
            return <h3 className="text-center">Update Promotion</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> value:&emsp; </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> asStackable:&emsp; </label>
                                                <input type="checkbox" placeholder="asStackable" name="asStackable" className="form-control" value={this.state.asStackable} onChange={this.changeasStackableHandler}/>


                                            <label> maxRedemptions:&emsp; </label>
                                                <input type="number" placeholder="maxRedemptions" name="maxRedemptions" className="form-control" value={this.state.maxRedemptions} onChange={this.changemaxRedemptionsHandler}/>

                                            <label> PromotionType:&emsp; </label>
                                                <select value={this.state.promotionType} onChange={this.changePromotionTypeHandler}>
                      <option name="PromotionType" className="form-control" >
                          Catalog
                      </option>
                      <option name="PromotionType" className="form-control" >
                          Cart
                      </option>
                      <option name="PromotionType" className="form-control" >
                          Shipping
                      </option>
                    </select>

                                            <label> DiscountType:&emsp; </label>
                                                <select value={this.state.discountType} onChange={this.changeDiscountTypeHandler}>
                      <option name="DiscountType" className="form-control" >
                          AmountOff
                      </option>
                      <option name="DiscountType" className="form-control" >
                          PercentOff
                      </option>
                      <option name="DiscountType" className="form-control" >
                          BuyXGetY
                      </option>
                      <option name="DiscountType" className="form-control" >
                          FreeShipping
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePromotion}>Save</button>
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

export default CreatePromotionComponent
