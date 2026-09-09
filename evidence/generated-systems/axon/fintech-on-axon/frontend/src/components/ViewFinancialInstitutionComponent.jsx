import React, { Component } from 'react'
import FinancialInstitutionService from '../services/FinancialInstitutionService'

class ViewFinancialInstitutionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            financialInstitution: {}
        }
    }

    componentDidMount(){
        FinancialInstitutionService.getFinancialInstitutionById(this.state.id).then( res => {
            this.setState({financialInstitution: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FinancialInstitution Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.financialInstitution.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> legalName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.financialInstitution.legalName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> countryOfIncorporation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.financialInstitution.countryOfIncorporation }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.financialInstitution.bic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> website:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.financialInstitution.website }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFinancialInstitutionComponent
