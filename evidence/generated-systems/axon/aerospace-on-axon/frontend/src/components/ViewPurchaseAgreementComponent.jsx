import React, { Component } from 'react'
import PurchaseAgreementService from '../services/PurchaseAgreementService'

class ViewPurchaseAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            purchaseAgreement: {}
        }
    }

    componentDidMount(){
        PurchaseAgreementService.getPurchaseAgreementById(this.state.id).then( res => {
            this.setState({purchaseAgreement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PurchaseAgreement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> agreementNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseAgreement.agreementNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.purchaseAgreement.effectiveDate }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPurchaseAgreementComponent
