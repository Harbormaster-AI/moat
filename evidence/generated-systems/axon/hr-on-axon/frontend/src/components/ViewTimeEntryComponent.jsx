import React, { Component } from 'react'
import TimeEntryService from '../services/TimeEntryService'

class ViewTimeEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            timeEntry: {}
        }
    }

    componentDidMount(){
        TimeEntryService.getTimeEntryById(this.state.id).then( res => {
            this.setState({timeEntry: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TimeEntry Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> entryDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timeEntry.entryDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> hoursWorked:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timeEntry.hoursWorked }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EntryType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.timeEntry.entryType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTimeEntryComponent
