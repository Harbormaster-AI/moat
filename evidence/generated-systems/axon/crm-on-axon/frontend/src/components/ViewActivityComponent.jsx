import React, { Component } from 'react'
import ActivityService from '../services/ActivityService'

class ViewActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            activity: {}
        }
    }

    componentDidMount(){
        ActivityService.getActivityById(this.state.id).then( res => {
            this.setState({activity: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Activity Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subject:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.subject }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.startAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.endAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> location:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.location }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ActivityType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.activityType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.activity.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewActivityComponent
