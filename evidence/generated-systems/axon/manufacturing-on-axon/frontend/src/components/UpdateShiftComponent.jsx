import React, { Component } from 'react'
import ShiftService from '../services/ShiftService';

class UpdateShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                shiftName: '',
                startTime: '',
                endTime: '',
                shiftType: ''
        }
        this.updateShift = this.updateShift.bind(this);

        this.changeshiftNameHandler = this.changeshiftNameHandler.bind(this);
        this.changestartTimeHandler = this.changestartTimeHandler.bind(this);
        this.changeendTimeHandler = this.changeendTimeHandler.bind(this);
        this.changeShiftTypeHandler = this.changeShiftTypeHandler.bind(this);
    }

    componentDidMount(){
        ShiftService.getShiftById(this.state.id).then( (res) =>{
            let shift = res.data;
            this.setState({
                shiftName: shift.shiftName,
                startTime: shift.startTime,
                endTime: shift.endTime,
                shiftType: shift.shiftType
            });
        });
    }

    updateShift = (e) => {
        e.preventDefault();
        let shift = {
            shiftId: this.state.id,
            shiftName: this.state.shiftName,
            startTime: this.state.startTime,
            endTime: this.state.endTime,
            shiftType: this.state.shiftType
        };
        console.log('shift => ' + JSON.stringify(shift));
        console.log('id => ' + JSON.stringify(this.state.id));
        ShiftService.updateShift(shift).then( res => {
            this.props.history.push('/shifts');
        });
    }

    changeshiftNameHandler= (event) => {
        this.setState({shiftName: event.target.value});
    }
    changestartTimeHandler= (event) => {
        this.setState({startTime: event.target.value});
    }
    changeendTimeHandler= (event) => {
        this.setState({endTime: event.target.value});
    }
    changeShiftTypeHandler= (event) => {
        this.setState({shiftType: event.target.value});
    }

    cancel(){
        this.props.history.push('/shifts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Shift</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> shiftName: </label>
                                                <input placeholder="shiftName" name="shiftName" className="form-control" value={this.state.shiftName} onChange={this.changeshiftNameHandler}/>

                                            <label> startTime: </label>
                                                <input placeholder="startTime" name="startTime" className="form-control" value={this.state.startTime} onChange={this.changestartTimeHandler}/>

                                            <label> endTime: </label>
                                                <input placeholder="endTime" name="endTime" className="form-control" value={this.state.endTime} onChange={this.changeendTimeHandler}/>

                                            <label> ShiftType: </label>
                                                <select value={this.state.shiftType} onChange={this.changeShiftTypeHandler}>
                      <option name="ShiftType" className="form-control" >
                          Day
                      </option>
                      <option name="ShiftType" className="form-control" >
                          Swing
                      </option>
                      <option name="ShiftType" className="form-control" >
                          Night
                      </option>
                      <option name="ShiftType" className="form-control" >
                          Weekend
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateShift}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateShiftComponent
