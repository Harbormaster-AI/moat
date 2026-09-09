import React, { Component } from 'react'
import JobRequisitionService from '../services/JobRequisitionService'

class ListJobRequisitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                jobRequisitions: []
        }
        this.addJobRequisition = this.addJobRequisition.bind(this);
        this.editJobRequisition = this.editJobRequisition.bind(this);
        this.deleteJobRequisition = this.deleteJobRequisition.bind(this);
    }

    deleteJobRequisition(id){
        JobRequisitionService.deleteJobRequisition(id).then( res => {
            this.setState({jobRequisitions: this.state.jobRequisitions.filter(jobRequisition => jobRequisition.jobRequisitionId !== id)});
        });
    }
    viewJobRequisition(id){
        this.props.history.push(`/view-jobRequisition/${id}`);
    }
    editJobRequisition(id){
        this.props.history.push(`/add-jobRequisition/${id}`);
    }

    componentDidMount(){
        JobRequisitionService.getJobRequisitions().then((res) => {
            this.setState({ jobRequisitions: res.data});
        });
    }

    addJobRequisition(){
        this.props.history.push('/add-jobRequisition/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">JobRequisition List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addJobRequisition}> Add JobRequisition</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RequisitionNumber </th>
                                    <th> Title </th>
                                    <th> Openings </th>
                                    <th> TargetStartDate </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.jobRequisitions.map(
                                        jobRequisition => 
                                        <tr key = {jobRequisition.jobRequisitionId}>
                                             <td> { jobRequisition.requisitionNumber } </td>
                                             <td> { jobRequisition.title } </td>
                                             <td> { jobRequisition.openings } </td>
                                             <td> { jobRequisition.targetStartDate } </td>
                                             <td> { jobRequisition.status } </td>
                                             <td> { jobRequisition.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editJobRequisition(jobRequisition.jobRequisitionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteJobRequisition(jobRequisition.jobRequisitionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewJobRequisition(jobRequisition.jobRequisitionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListJobRequisitionComponent
