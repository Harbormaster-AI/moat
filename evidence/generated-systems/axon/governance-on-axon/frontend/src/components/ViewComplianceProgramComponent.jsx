import React, { Component } from 'react'
import ComplianceProgramService from '../services/ComplianceProgramService'

class ViewComplianceProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            complianceProgram: {}
        }
    }

    componentDidMount(){
        ComplianceProgramService.getComplianceProgramById(this.state.id).then( res => {
            this.setState({complianceProgram: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ComplianceProgram Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceProgram.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> framework:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceProgram.framework }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.complianceProgram.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewComplianceProgramComponent
