import React, { Component } from 'react'
import TrainingRunService from '../services/TrainingRunService'

class ViewTrainingRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            trainingRun: {}
        }
    }

    componentDidMount(){
        TrainingRunService.getTrainingRunById(this.state.id).then( res => {
            this.setState({trainingRun: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TrainingRun Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> runLabel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingRun.runLabel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingRun.startedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> completedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingRun.completedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingRun.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTrainingRunComponent
