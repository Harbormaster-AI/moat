import React, { Component } from 'react'
import JobProfileService from '../services/JobProfileService'

class ListJobProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                jobProfiles: []
        }
        this.addJobProfile = this.addJobProfile.bind(this);
        this.editJobProfile = this.editJobProfile.bind(this);
        this.deleteJobProfile = this.deleteJobProfile.bind(this);
    }

    deleteJobProfile(id){
        JobProfileService.deleteJobProfile(id).then( res => {
            this.setState({jobProfiles: this.state.jobProfiles.filter(jobProfile => jobProfile.jobProfileId !== id)});
        });
    }
    viewJobProfile(id){
        this.props.history.push(`/view-jobProfile/${id}`);
    }
    editJobProfile(id){
        this.props.history.push(`/add-jobProfile/${id}`);
    }

    componentDidMount(){
        JobProfileService.getJobProfiles().then((res) => {
            this.setState({ jobProfiles: res.data});
        });
    }

    addJobProfile(){
        this.props.history.push('/add-jobProfile/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">JobProfile List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addJobProfile}> Add JobProfile</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> JobCode </th>
                                    <th> JobLevel </th>
                                    <th> ExemptStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.jobProfiles.map(
                                        jobProfile => 
                                        <tr key = {jobProfile.jobProfileId}>
                                             <td> { jobProfile.title } </td>
                                             <td> { jobProfile.jobCode } </td>
                                             <td> { jobProfile.jobLevel } </td>
                                             <td> { jobProfile.exemptStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editJobProfile(jobProfile.jobProfileId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteJobProfile(jobProfile.jobProfileId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewJobProfile(jobProfile.jobProfileId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListJobProfileComponent
