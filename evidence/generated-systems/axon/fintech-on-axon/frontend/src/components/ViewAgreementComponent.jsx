import React, { Component } from 'react'
import AgreementService from '../services/AgreementService'

class ViewAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            agreement: {}
        }
    }

    componentDidMount(){
        AgreementService.getAgreementById(this.state.id).then( res => {
            this.setState({agreement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Agreement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> agreementNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agreement.agreementNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agreement.effectiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AgreementType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agreement.agreementType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.agreement.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAgreementComponent
