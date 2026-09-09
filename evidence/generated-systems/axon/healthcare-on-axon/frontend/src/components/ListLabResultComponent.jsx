import React, { Component } from 'react'
import LabResultService from '../services/LabResultService'

class ListLabResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                labResults: []
        }
        this.addLabResult = this.addLabResult.bind(this);
        this.editLabResult = this.editLabResult.bind(this);
        this.deleteLabResult = this.deleteLabResult.bind(this);
    }

    deleteLabResult(id){
        LabResultService.deleteLabResult(id).then( res => {
            this.setState({labResults: this.state.labResults.filter(labResult => labResult.labResultId !== id)});
        });
    }
    viewLabResult(id){
        this.props.history.push(`/view-labResult/${id}`);
    }
    editLabResult(id){
        this.props.history.push(`/add-labResult/${id}`);
    }

    componentDidMount(){
        LabResultService.getLabResults().then((res) => {
            this.setState({ labResults: res.data});
        });
    }

    addLabResult(){
        this.props.history.push('/add-labResult/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LabResult List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLabResult}> Add LabResult</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ResultCode </th>
                                    <th> IssuedDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.labResults.map(
                                        labResult => 
                                        <tr key = {labResult.labResultId}>
                                             <td> { labResult.resultCode } </td>
                                             <td> { labResult.issuedDate } </td>
                                             <td> { labResult.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editLabResult(labResult.labResultId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLabResult(labResult.labResultId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLabResult(labResult.labResultId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLabResultComponent
