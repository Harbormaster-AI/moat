import React, { Component } from 'react'
import ApplicationService from '../services/ApplicationService'

class ListApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                applications: []
        }
        this.addApplication = this.addApplication.bind(this);
        this.editApplication = this.editApplication.bind(this);
        this.deleteApplication = this.deleteApplication.bind(this);
    }

    deleteApplication(id){
        ApplicationService.deleteApplication(id).then( res => {
            this.setState({applications: this.state.applications.filter(application => application.applicationId !== id)});
        });
    }
    viewApplication(id){
        this.props.history.push(`/view-application/${id}`);
    }
    editApplication(id){
        this.props.history.push(`/add-application/${id}`);
    }

    componentDidMount(){
        ApplicationService.getApplications().then((res) => {
            this.setState({ applications: res.data});
        });
    }

    addApplication(){
        this.props.history.push('/add-application/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Application List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addApplication}> Add Application</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ApplicationNumber </th>
                                    <th> SubmissionDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.applications.map(
                                        application => 
                                        <tr key = {application.applicationId}>
                                             <td> { application.applicationNumber } </td>
                                             <td> { application.submissionDate } </td>
                                             <td> { application.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editApplication(application.applicationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteApplication(application.applicationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewApplication(application.applicationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListApplicationComponent
