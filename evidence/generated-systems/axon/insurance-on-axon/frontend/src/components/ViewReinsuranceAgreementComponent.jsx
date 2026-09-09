import React, { Component } from 'react'
import ReinsuranceAgreementService from '../services/ReinsuranceAgreementService'

class ViewReinsuranceAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            reinsuranceAgreement: {}
        }
    }

    componentDidMount(){
        ReinsuranceAgreementService.getReinsuranceAgreementById(this.state.id).then( res => {
            this.setState({reinsuranceAgreement: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ReinsuranceAgreement Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> agreementNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.agreementNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectivePeriod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.effectivePeriod }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> retention:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.retention }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> limit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.limit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> cessionPercentage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.cessionPercentage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReinsuranceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.reinsuranceType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TreatyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.reinsuranceAgreement.treatyType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReinsuranceAgreementComponent
