import React, { Component } from 'react'
import TrainingEnrollmentService from '../services/TrainingEnrollmentService'

class ListTrainingEnrollmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                trainingEnrollments: []
        }
        this.addTrainingEnrollment = this.addTrainingEnrollment.bind(this);
        this.editTrainingEnrollment = this.editTrainingEnrollment.bind(this);
        this.deleteTrainingEnrollment = this.deleteTrainingEnrollment.bind(this);
    }

    deleteTrainingEnrollment(id){
        TrainingEnrollmentService.deleteTrainingEnrollment(id).then( res => {
            this.setState({trainingEnrollments: this.state.trainingEnrollments.filter(trainingEnrollment => trainingEnrollment.trainingEnrollmentId !== id)});
        });
    }
    viewTrainingEnrollment(id){
        this.props.history.push(`/view-trainingEnrollment/${id}`);
    }
    editTrainingEnrollment(id){
        this.props.history.push(`/add-trainingEnrollment/${id}`);
    }

    componentDidMount(){
        TrainingEnrollmentService.getTrainingEnrollments().then((res) => {
            this.setState({ trainingEnrollments: res.data});
        });
    }

    addTrainingEnrollment(){
        this.props.history.push('/add-trainingEnrollment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TrainingEnrollment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTrainingEnrollment}> Add TrainingEnrollment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EnrollmentNumber </th>
                                    <th> CompletionDate </th>
                                    <th> Score </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.trainingEnrollments.map(
                                        trainingEnrollment => 
                                        <tr key = {trainingEnrollment.trainingEnrollmentId}>
                                             <td> { trainingEnrollment.enrollmentNumber } </td>
                                             <td> { trainingEnrollment.completionDate } </td>
                                             <td> { trainingEnrollment.score } </td>
                                             <td> { trainingEnrollment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editTrainingEnrollment(trainingEnrollment.trainingEnrollmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTrainingEnrollment(trainingEnrollment.trainingEnrollmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTrainingEnrollment(trainingEnrollment.trainingEnrollmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTrainingEnrollmentComponent
