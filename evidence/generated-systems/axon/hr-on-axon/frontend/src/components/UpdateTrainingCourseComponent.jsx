import React, { Component } from 'react'
import TrainingCourseService from '../services/TrainingCourseService';

class UpdateTrainingCourseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                title: '',
                durationHours: '',
                deliveryMethod: ''
        }
        this.updateTrainingCourse = this.updateTrainingCourse.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changedurationHoursHandler = this.changedurationHoursHandler.bind(this);
        this.changeDeliveryMethodHandler = this.changeDeliveryMethodHandler.bind(this);
    }

    componentDidMount(){
        TrainingCourseService.getTrainingCourseById(this.state.id).then( (res) =>{
            let trainingCourse = res.data;
            this.setState({
                code: trainingCourse.code,
                title: trainingCourse.title,
                durationHours: trainingCourse.durationHours,
                deliveryMethod: trainingCourse.deliveryMethod
            });
        });
    }

    updateTrainingCourse = (e) => {
        e.preventDefault();
        let trainingCourse = {
            trainingCourseId: this.state.id,
            code: this.state.code,
            title: this.state.title,
            durationHours: this.state.durationHours,
            deliveryMethod: this.state.deliveryMethod
        };
        console.log('trainingCourse => ' + JSON.stringify(trainingCourse));
        console.log('id => ' + JSON.stringify(this.state.id));
        TrainingCourseService.updateTrainingCourse(trainingCourse).then( res => {
            this.props.history.push('/trainingCourses');
        });
    }

    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changedurationHoursHandler= (event) => {
        this.setState({durationHours: event.target.value});
    }
    changeDeliveryMethodHandler= (event) => {
        this.setState({deliveryMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/trainingCourses');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TrainingCourse</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> durationHours: </label>
                                                <input placeholder="durationHours" name="durationHours" className="form-control" value={this.state.durationHours} onChange={this.changedurationHoursHandler}/>

                                            <label> DeliveryMethod: </label>
                                                <select value={this.state.deliveryMethod} onChange={this.changeDeliveryMethodHandler}>
                      <option name="DeliveryMethod" className="form-control" >
                          Classroom
                      </option>
                      <option name="DeliveryMethod" className="form-control" >
                          Virtual
                      </option>
                      <option name="DeliveryMethod" className="form-control" >
                          SelfPaced
                      </option>
                      <option name="DeliveryMethod" className="form-control" >
                          Blended
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTrainingCourse}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateTrainingCourseComponent
