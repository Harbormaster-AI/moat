import React, { Component } from 'react'
import PayrollCalendarService from '../services/PayrollCalendarService'

class ViewPayrollCalendarComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            payrollCalendar: {}
        }
    }

    componentDidMount(){
        PayrollCalendarService.getPayrollCalendarById(this.state.id).then( res => {
            this.setState({payrollCalendar: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PayrollCalendar Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollCalendar.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> country:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollCalendar.country }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PayFrequency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.payrollCalendar.payFrequency }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPayrollCalendarComponent
