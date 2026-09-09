import React, { Component } from 'react'
import PolicyCoverageService from '../services/PolicyCoverageService'

class ViewPolicyCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            policyCoverage: {}
        }
    }

    componentDidMount(){
        PolicyCoverageService.getPolicyCoverageById(this.state.id).then( res => {
            this.setState({policyCoverage: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PolicyCoverage Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> limit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policyCoverage.limit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> deductible:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policyCoverage.deductible }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> premium:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policyCoverage.premium }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CoverageType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.policyCoverage.coverageType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPolicyCoverageComponent
