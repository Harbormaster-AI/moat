import React, { Component } from 'react'
import TimesheetService from '../services/TimesheetService'

class ViewTimesheetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            timesheet: {}
        }
    }

    componentDidMount(){
        TimesheetService.getTimesheetById(this.state.id).then( res => {
            this.setState({timesheet: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Timesheet Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timesheet.periodStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> periodEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timesheet.periodEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> submissionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timesheet.submissionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timesheet.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTimesheetComponent
