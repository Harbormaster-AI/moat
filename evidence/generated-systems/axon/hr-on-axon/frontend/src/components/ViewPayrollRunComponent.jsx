import React, { Component } from 'react'
import PayrollRunService from '../services/PayrollRunService'

class ViewPayrollRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            payrollRun: {}
        }
    }

    componentDidMount(){
        PayrollRunService.getPayrollRunById(this.state.id).then( res => {
            this.setState({payrollRun: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PayrollRun Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> runNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollRun.runNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollRun.periodStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollRun.periodEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> paymentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollRun.paymentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollRun.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPayrollRunComponent
