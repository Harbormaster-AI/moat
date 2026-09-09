import React, { Component } from 'react'
import TrainingCourseService from '../services/TrainingCourseService';

class CreateTrainingCourseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                title: '',
                durationHours: '',
                deliveryMethod: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changedurationHoursHandler = this.changedurationHoursHandler.bind(this);
        this.changeDeliveryMethodHandler = this.changeDeliveryMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateTrainingCourse = (e) => {
        e.preventDefault();
        let trainingCourse = {
                trainingCourseId: this.state.id,
                code: this.state.code,
                title: this.state.title,
                durationHours: this.state.durationHours,
                deliveryMethod: this.state.deliveryMethod
            };
        console.log('trainingCourse => ' + JSON.stringify(trainingCourse));

        // step 5
        if(this.state.id === '_add'){
            trainingCourse.trainingCourseId=''
            TrainingCourseService.createTrainingCourse(trainingCourse).then(res =>{
                this.props.history.push('/trainingCourses');
            });
        }else{
            TrainingCourseService.updateTrainingCourse(trainingCourse).then( res => {
                this.props.history.push('/trainingCourses');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TrainingCourse</h3>
        }else{
            return <h3 className="text-center">Update TrainingCourse</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> durationHours:&emsp; </label>
                                                <input placeholder="durationHours" name="durationHours" className="form-control" value={this.state.durationHours} onChange={this.changedurationHoursHandler}/>

                                            <label> DeliveryMethod:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTrainingCourse}>Save</button>
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

export default CreateTrainingCourseComponent
