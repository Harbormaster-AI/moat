import React, { Component } from 'react'
import BuildScheduleService from '../services/BuildScheduleService'

class ViewBuildScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            buildSchedule: {}
        }
    }

    componentDidMount(){
        BuildScheduleService.getBuildScheduleById(this.state.id).then( res => {
            this.setState({buildSchedule: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View BuildSchedule Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scheduleNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.buildSchedule.scheduleNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.buildSchedule.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewBuildScheduleComponent
