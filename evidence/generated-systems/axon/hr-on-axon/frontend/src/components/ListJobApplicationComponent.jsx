import React, { Component } from 'react'
import JobApplicationService from '../services/JobApplicationService'

class ListJobApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                jobApplications: []
        }
        this.addJobApplication = this.addJobApplication.bind(this);
        this.editJobApplication = this.editJobApplication.bind(this);
        this.deleteJobApplication = this.deleteJobApplication.bind(this);
    }

    deleteJobApplication(id){
        JobApplicationService.deleteJobApplication(id).then( res => {
            this.setState({jobApplications: this.state.jobApplications.filter(jobApplication => jobApplication.jobApplicationId !== id)});
        });
    }
    viewJobApplication(id){
        this.props.history.push(`/view-jobApplication/${id}`);
    }
    editJobApplication(id){
        this.props.history.push(`/add-jobApplication/${id}`);
    }

    componentDidMount(){
        JobApplicationService.getJobApplications().then((res) => {
            this.setState({ jobApplications: res.data});
        });
    }

    addJobApplication(){
        this.props.history.push('/add-jobApplication/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">JobApplication List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addJobApplication}> Add JobApplication</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ApplicationNumber </th>
                                    <th> AppliedDate </th>
                                    <th> ResumeUrl </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.jobApplications.map(
                                        jobApplication => 
                                        <tr key = {jobApplication.jobApplicationId}>
                                             <td> { jobApplication.applicationNumber } </td>
                                             <td> { jobApplication.appliedDate } </td>
                                             <td> { jobApplication.resumeUrl } </td>
                                             <td> { jobApplication.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editJobApplication(jobApplication.jobApplicationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteJobApplication(jobApplication.jobApplicationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewJobApplication(jobApplication.jobApplicationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListJobApplicationComponent
