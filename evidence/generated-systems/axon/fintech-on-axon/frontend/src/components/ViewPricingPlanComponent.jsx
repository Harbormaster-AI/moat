import React, { Component } from 'react'
import PricingPlanService from '../services/PricingPlanService'

class ViewPricingPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            pricingPlan: {}
        }
    }

    componentDidMount(){
        PricingPlanService.getPricingPlanById(this.state.id).then( res => {
            this.setState({pricingPlan: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PricingPlan Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.pricingPlan.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> planCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.pricingPlan.planCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> baseCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.pricingPlan.baseCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.pricingPlan.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPricingPlanComponent
