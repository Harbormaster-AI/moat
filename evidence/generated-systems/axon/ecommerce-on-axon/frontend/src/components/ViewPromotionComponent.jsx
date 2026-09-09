import React, { Component } from 'react'
import PromotionService from '../services/PromotionService'

class ViewPromotionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            promotion: {}
        }
    }

    componentDidMount(){
        PromotionService.getPromotionById(this.state.id).then( res => {
            this.setState({promotion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Promotion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> value:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.value }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asStackable:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.asStackable }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maxRedemptions:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.maxRedemptions }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PromotionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.promotionType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DiscountType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.promotion.discountType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPromotionComponent
