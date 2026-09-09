import React, { Component } from 'react'
import TimeEntryService from '../services/TimeEntryService'

class ListTimeEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                timeEntrys: []
        }
        this.addTimeEntry = this.addTimeEntry.bind(this);
        this.editTimeEntry = this.editTimeEntry.bind(this);
        this.deleteTimeEntry = this.deleteTimeEntry.bind(this);
    }

    deleteTimeEntry(id){
        TimeEntryService.deleteTimeEntry(id).then( res => {
            this.setState({timeEntrys: this.state.timeEntrys.filter(timeEntry => timeEntry.timeEntryId !== id)});
        });
    }
    viewTimeEntry(id){
        this.props.history.push(`/view-timeEntry/${id}`);
    }
    editTimeEntry(id){
        this.props.history.push(`/add-timeEntry/${id}`);
    }

    componentDidMount(){
        TimeEntryService.getTimeEntrys().then((res) => {
            this.setState({ timeEntrys: res.data});
        });
    }

    addTimeEntry(){
        this.props.history.push('/add-timeEntry/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TimeEntry List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTimeEntry}> Add TimeEntry</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EntryDate </th>
                                    <th> HoursWorked </th>
                                    <th> EntryType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.timeEntrys.map(
                                        timeEntry => 
                                        <tr key = {timeEntry.timeEntryId}>
                                             <td> { timeEntry.entryDate } </td>
                                             <td> { timeEntry.hoursWorked } </td>
                                             <td> { timeEntry.entryType } </td>
                                             <td>
                                                 <button onClick={ () => this.editTimeEntry(timeEntry.timeEntryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTimeEntry(timeEntry.timeEntryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTimeEntry(timeEntry.timeEntryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTimeEntryComponent
