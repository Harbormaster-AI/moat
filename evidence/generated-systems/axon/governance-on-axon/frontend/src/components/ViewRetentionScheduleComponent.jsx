import React, { Component } from 'react'
import RetentionScheduleService from '../services/RetentionScheduleService'

class ViewRetentionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            retentionSchedule: {}
        }
    }

    componentDidMount(){
        RetentionScheduleService.getRetentionScheduleById(this.state.id).then( res => {
            this.setState({retentionSchedule: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View RetentionSchedule Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.retentionSchedule.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> retentionPeriodMonths:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.retentionSchedule.retentionPeriodMonths }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RetentionTrigger:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.retentionSchedule.retentionTrigger }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DispositionAction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.retentionSchedule.dispositionAction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.retentionSchedule.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRetentionScheduleComponent
