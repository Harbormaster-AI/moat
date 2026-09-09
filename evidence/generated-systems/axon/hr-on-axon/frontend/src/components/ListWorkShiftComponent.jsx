import React, { Component } from 'react'
import WorkShiftService from '../services/WorkShiftService'

class ListWorkShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                workShifts: []
        }
        this.addWorkShift = this.addWorkShift.bind(this);
        this.editWorkShift = this.editWorkShift.bind(this);
        this.deleteWorkShift = this.deleteWorkShift.bind(this);
    }

    deleteWorkShift(id){
        WorkShiftService.deleteWorkShift(id).then( res => {
            this.setState({workShifts: this.state.workShifts.filter(workShift => workShift.workShiftId !== id)});
        });
    }
    viewWorkShift(id){
        this.props.history.push(`/view-workShift/${id}`);
    }
    editWorkShift(id){
        this.props.history.push(`/add-workShift/${id}`);
    }

    componentDidMount(){
        WorkShiftService.getWorkShifts().then((res) => {
            this.setState({ workShifts: res.data});
        });
    }

    addWorkShift(){
        this.props.history.push('/add-workShift/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WorkShift List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWorkShift}> Add WorkShift</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> StartTime </th>
                                    <th> EndTime </th>
                                    <th> BreakMinutes </th>
                                    <th> DayOfWeek </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.workShifts.map(
                                        workShift => 
                                        <tr key = {workShift.workShiftId}>
                                             <td> { workShift.startTime } </td>
                                             <td> { workShift.endTime } </td>
                                             <td> { workShift.breakMinutes } </td>
                                             <td> { workShift.dayOfWeek } </td>
                                             <td>
                                                 <button onClick={ () => this.editWorkShift(workShift.workShiftId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWorkShift(workShift.workShiftId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWorkShift(workShift.workShiftId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWorkShiftComponent
