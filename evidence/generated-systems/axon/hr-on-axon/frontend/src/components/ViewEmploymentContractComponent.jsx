import React, { Component } from 'react'
import EmploymentContractService from '../services/EmploymentContractService'

class ViewEmploymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            employmentContract: {}
        }
    }

    componentDidMount(){
        EmploymentContractService.getEmploymentContractById(this.state.id).then( res => {
            this.setState({employmentContract: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View EmploymentContract Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> contractNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.contractNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.endDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workHoursPerWeek:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.workHoursPerWeek }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EmploymentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.employmentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PayFrequency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.employmentContract.payFrequency }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEmploymentContractComponent
