import React, { Component } from 'react'
import RiskService from '../services/RiskService'

class ViewRiskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            risk: {}
        }
    }

    componentDidMount(){
        RiskService.getRiskById(this.state.id).then( res => {
            this.setState({risk: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Risk Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> inherentRiskScore:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.inherentRiskScore }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> residualRiskScore:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.residualRiskScore }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Category:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.category }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Impact:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.impact }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Likelihood:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.likelihood }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.risk.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRiskComponent
