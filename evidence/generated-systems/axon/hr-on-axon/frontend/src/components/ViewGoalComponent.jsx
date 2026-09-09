import React, { Component } from 'react'
import GoalService from '../services/GoalService'

class ViewGoalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            goal: {}
        }
    }

    componentDidMount(){
        GoalService.getGoalById(this.state.id).then( res => {
            this.setState({goal: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Goal Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goal.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goal.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> targetDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goal.targetDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> weight:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goal.weight }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goal.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGoalComponent
