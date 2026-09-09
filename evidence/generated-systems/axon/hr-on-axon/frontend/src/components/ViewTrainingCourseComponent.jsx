import React, { Component } from 'react'
import TrainingCourseService from '../services/TrainingCourseService'

class ViewTrainingCourseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            trainingCourse: {}
        }
    }

    componentDidMount(){
        TrainingCourseService.getTrainingCourseById(this.state.id).then( res => {
            this.setState({trainingCourse: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TrainingCourse Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingCourse.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingCourse.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> durationHours:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingCourse.durationHours }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DeliveryMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.trainingCourse.deliveryMethod }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTrainingCourseComponent
