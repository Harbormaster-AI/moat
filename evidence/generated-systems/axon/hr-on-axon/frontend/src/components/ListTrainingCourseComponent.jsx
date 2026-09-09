import React, { Component } from 'react'
import TrainingCourseService from '../services/TrainingCourseService'

class ListTrainingCourseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                trainingCourses: []
        }
        this.addTrainingCourse = this.addTrainingCourse.bind(this);
        this.editTrainingCourse = this.editTrainingCourse.bind(this);
        this.deleteTrainingCourse = this.deleteTrainingCourse.bind(this);
    }

    deleteTrainingCourse(id){
        TrainingCourseService.deleteTrainingCourse(id).then( res => {
            this.setState({trainingCourses: this.state.trainingCourses.filter(trainingCourse => trainingCourse.trainingCourseId !== id)});
        });
    }
    viewTrainingCourse(id){
        this.props.history.push(`/view-trainingCourse/${id}`);
    }
    editTrainingCourse(id){
        this.props.history.push(`/add-trainingCourse/${id}`);
    }

    componentDidMount(){
        TrainingCourseService.getTrainingCourses().then((res) => {
            this.setState({ trainingCourses: res.data});
        });
    }

    addTrainingCourse(){
        this.props.history.push('/add-trainingCourse/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TrainingCourse List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTrainingCourse}> Add TrainingCourse</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Title </th>
                                    <th> DurationHours </th>
                                    <th> DeliveryMethod </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.trainingCourses.map(
                                        trainingCourse => 
                                        <tr key = {trainingCourse.trainingCourseId}>
                                             <td> { trainingCourse.code } </td>
                                             <td> { trainingCourse.title } </td>
                                             <td> { trainingCourse.durationHours } </td>
                                             <td> { trainingCourse.deliveryMethod } </td>
                                             <td>
                                                 <button onClick={ () => this.editTrainingCourse(trainingCourse.trainingCourseId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTrainingCourse(trainingCourse.trainingCourseId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTrainingCourse(trainingCourse.trainingCourseId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListTrainingCourseComponent
