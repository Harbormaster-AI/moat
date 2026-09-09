import React, { Component } from 'react'
import ShippingMethodService from '../services/ShippingMethodService'

class ViewShippingMethodComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            shippingMethod: {}
        }
    }

    componentDidMount(){
        ShippingMethodService.getShippingMethodById(this.state.id).then( res => {
            this.setState({shippingMethod: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ShippingMethod Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shippingMethod.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> flatRate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shippingMethod.flatRate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> estimatedDays:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shippingMethod.estimatedDays }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shippingMethod.asActive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MethodType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shippingMethod.methodType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewShippingMethodComponent
