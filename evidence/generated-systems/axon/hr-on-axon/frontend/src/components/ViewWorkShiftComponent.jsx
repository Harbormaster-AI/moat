import React, { Component } from 'react'
import WorkShiftService from '../services/WorkShiftService'

class ViewWorkShiftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            workShift: {}
        }
    }

    componentDidMount(){
        WorkShiftService.getWorkShiftById(this.state.id).then( res => {
            this.setState({workShift: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View WorkShift Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workShift.startTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workShift.endTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> breakMinutes:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workShift.breakMinutes }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DayOfWeek:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.workShift.dayOfWeek }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewWorkShiftComponent
