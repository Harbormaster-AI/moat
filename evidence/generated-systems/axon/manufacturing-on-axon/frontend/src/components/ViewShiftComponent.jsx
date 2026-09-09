import React, { Component } from 'react'
import ShiftService from '../services/ShiftService'

class ViewShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            shift: {}
        }
    }

    componentDidMount(){
        ShiftService.getShiftById(this.state.id).then( res => {
            this.setState({shift: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Shift Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> shiftName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shift.shiftName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shift.startTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shift.endTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ShiftType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.shift.shiftType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewShiftComponent
