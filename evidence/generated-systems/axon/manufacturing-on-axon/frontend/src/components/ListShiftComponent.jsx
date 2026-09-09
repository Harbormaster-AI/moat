import React, { Component } from 'react'
import ShiftService from '../services/ShiftService'

class ListShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                shifts: []
        }
        this.addShift = this.addShift.bind(this);
        this.editShift = this.editShift.bind(this);
        this.deleteShift = this.deleteShift.bind(this);
    }

    deleteShift(id){
        ShiftService.deleteShift(id).then( res => {
            this.setState({shifts: this.state.shifts.filter(shift => shift.shiftId !== id)});
        });
    }
    viewShift(id){
        this.props.history.push(`/view-shift/${id}`);
    }
    editShift(id){
        this.props.history.push(`/add-shift/${id}`);
    }

    componentDidMount(){
        ShiftService.getShifts().then((res) => {
            this.setState({ shifts: res.data});
        });
    }

    addShift(){
        this.props.history.push('/add-shift/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Shift List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addShift}> Add Shift</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ShiftName </th>
                                    <th> StartTime </th>
                                    <th> EndTime </th>
                                    <th> ShiftType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.shifts.map(
                                        shift => 
                                        <tr key = {shift.shiftId}>
                                             <td> { shift.shiftName } </td>
                                             <td> { shift.startTime } </td>
                                             <td> { shift.endTime } </td>
                                             <td> { shift.shiftType } </td>
                                             <td>
                                                 <button onClick={ () => this.editShift(shift.shiftId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteShift(shift.shiftId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewShift(shift.shiftId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListShiftComponent
