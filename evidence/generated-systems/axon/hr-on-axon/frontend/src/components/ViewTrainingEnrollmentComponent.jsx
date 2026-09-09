import React, { Component } from 'react'
import TrainingEnrollmentService from '../services/TrainingEnrollmentService'

class ViewTrainingEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            trainingEnrollment: {}
        }
    }

    componentDidMount(){
        TrainingEnrollmentService.getTrainingEnrollmentById(this.state.id).then( res => {
            this.setState({trainingEnrollment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TrainingEnrollment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> enrollmentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingEnrollment.enrollmentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> completionDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingEnrollment.completionDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> score:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingEnrollment.score }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingEnrollment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTrainingEnrollmentComponent
