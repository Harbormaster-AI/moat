import React, { Component } from 'react'
import BrandSafetyPolicyService from '../services/BrandSafetyPolicyService'

class ViewBrandSafetyPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            brandSafetyPolicy: {}
        }
    }

    componentDidMount(){
        BrandSafetyPolicyService.getBrandSafetyPolicyById(this.state.id).then( res => {
            this.setState({brandSafetyPolicy: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BrandSafetyPolicy Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Level:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.brandSafetyPolicy.level }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ContentRatingThreshold:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.brandSafetyPolicy.contentRatingThreshold }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBrandSafetyPolicyComponent
