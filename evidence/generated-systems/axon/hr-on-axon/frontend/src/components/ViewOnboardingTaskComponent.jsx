import React, { Component } from 'react'
import OnboardingTaskService from '../services/OnboardingTaskService'

class ViewOnboardingTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            onboardingTask: {}
        }
    }

    componentDidMount(){
        OnboardingTaskService.getOnboardingTaskById(this.state.id).then( res => {
            this.setState({onboardingTask: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View OnboardingTask Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taskNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.onboardingTask.taskNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.onboardingTask.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dueDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.onboardingTask.dueDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.onboardingTask.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOnboardingTaskComponent
