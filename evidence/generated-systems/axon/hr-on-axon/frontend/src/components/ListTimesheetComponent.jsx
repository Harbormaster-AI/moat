import React, { Component } from 'react'
import TimesheetService from '../services/TimesheetService'

class ListTimesheetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                timesheets: []
        }
        this.addTimesheet = this.addTimesheet.bind(this);
        this.editTimesheet = this.editTimesheet.bind(this);
        this.deleteTimesheet = this.deleteTimesheet.bind(this);
    }

    deleteTimesheet(id){
        TimesheetService.deleteTimesheet(id).then( res => {
            this.setState({timesheets: this.state.timesheets.filter(timesheet => timesheet.timesheetId !== id)});
        });
    }
    viewTimesheet(id){
        this.props.history.push(`/view-timesheet/${id}`);
    }
    editTimesheet(id){
        this.props.history.push(`/add-timesheet/${id}`);
    }

    componentDidMount(){
        TimesheetService.getTimesheets().then((res) => {
            this.setState({ timesheets: res.data});
        });
    }

    addTimesheet(){
        this.props.history.push('/add-timesheet/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Timesheet List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTimesheet}> Add Timesheet</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PeriodStart </th>
                                    <th> PeriodEnd </th>
                                    <th> SubmissionDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.timesheets.map(
                                        timesheet => 
                                        <tr key = {timesheet.timesheetId}>
                                             <td> { timesheet.periodStart } </td>
                                             <td> { timesheet.periodEnd } </td>
                                             <td> { timesheet.submissionDate } </td>
                                             <td> { timesheet.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editTimesheet(timesheet.timesheetId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTimesheet(timesheet.timesheetId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTimesheet(timesheet.timesheetId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListTimesheetComponent
