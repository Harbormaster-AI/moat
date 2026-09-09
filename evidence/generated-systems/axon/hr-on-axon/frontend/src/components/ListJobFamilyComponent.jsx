import React, { Component } from 'react'
import JobFamilyService from '../services/JobFamilyService'

class ListJobFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                jobFamilys: []
        }
        this.addJobFamily = this.addJobFamily.bind(this);
        this.editJobFamily = this.editJobFamily.bind(this);
        this.deleteJobFamily = this.deleteJobFamily.bind(this);
    }

    deleteJobFamily(id){
        JobFamilyService.deleteJobFamily(id).then( res => {
            this.setState({jobFamilys: this.state.jobFamilys.filter(jobFamily => jobFamily.jobFamilyId !== id)});
        });
    }
    viewJobFamily(id){
        this.props.history.push(`/view-jobFamily/${id}`);
    }
    editJobFamily(id){
        this.props.history.push(`/add-jobFamily/${id}`);
    }

    componentDidMount(){
        JobFamilyService.getJobFamilys().then((res) => {
            this.setState({ jobFamilys: res.data});
        });
    }

    addJobFamily(){
        this.props.history.push('/add-jobFamily/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">JobFamily List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addJobFamily}> Add JobFamily</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Description </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.jobFamilys.map(
                                        jobFamily => 
                                        <tr key = {jobFamily.jobFamilyId}>
                                             <td> { jobFamily.name } </td>
                                             <td> { jobFamily.description } </td>
                                             <td>
                                                 <button onClick={ () => this.editJobFamily(jobFamily.jobFamilyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteJobFamily(jobFamily.jobFamilyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewJobFamily(jobFamily.jobFamilyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListJobFamilyComponent
