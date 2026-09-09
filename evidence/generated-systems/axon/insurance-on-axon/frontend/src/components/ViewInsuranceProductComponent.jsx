import React, { Component } from 'react'
import InsuranceProductService from '../services/InsuranceProductService'

class ViewInsuranceProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            insuranceProduct: {}
        }
    }

    componentDidMount(){
        InsuranceProductService.getInsuranceProductById(this.state.id).then( res => {
            this.setState({insuranceProduct: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InsuranceProduct Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuranceProduct.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> productCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuranceProduct.productCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LineOfBusiness:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.insuranceProduct.lineOfBusiness }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInsuranceProductComponent
