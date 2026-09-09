import React, { Component } from 'react'
import NonconformanceService from '../services/NonconformanceService'

class ViewNonconformanceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            nonconformance: {}
        }
    }

    componentDidMount(){
        NonconformanceService.getNonconformanceById(this.state.id).then( res => {
            this.setState({nonconformance: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Nonconformance Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ncNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.ncNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> containmentAction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.containmentAction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> NcType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.ncType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.severity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.nonconformance.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewNonconformanceComponent
