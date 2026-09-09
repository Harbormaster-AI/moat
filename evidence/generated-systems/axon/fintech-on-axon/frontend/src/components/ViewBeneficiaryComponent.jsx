import React, { Component } from 'react'
import BeneficiaryService from '../services/BeneficiaryService'

class ViewBeneficiaryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            beneficiary: {}
        }
    }

    componentDidMount(){
        BeneficiaryService.getBeneficiaryById(this.state.id).then( res => {
            this.setState({beneficiary: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Beneficiary Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.beneficiary.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> accountIdentifier:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.beneficiary.accountIdentifier }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> iban:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.beneficiary.iban }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.beneficiary.bic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.beneficiary.address }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBeneficiaryComponent
