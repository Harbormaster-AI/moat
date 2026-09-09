import React, { Component } from 'react'
import CoverageService from '../services/CoverageService'

class ViewCoverageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            coverage: {}
        }
    }

    componentDidMount(){
        CoverageService.getCoverageById(this.state.id).then( res => {
            this.setState({coverage: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Coverage Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> memberId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverage.memberId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> groupNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverage.groupNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverage.effectiveDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverage.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CoverageType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.coverage.coverageType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCoverageComponent
