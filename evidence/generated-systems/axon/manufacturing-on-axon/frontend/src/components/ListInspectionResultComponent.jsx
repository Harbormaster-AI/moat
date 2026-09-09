import React, { Component } from 'react'
import InspectionResultService from '../services/InspectionResultService'

class ListInspectionResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inspectionResults: []
        }
        this.addInspectionResult = this.addInspectionResult.bind(this);
        this.editInspectionResult = this.editInspectionResult.bind(this);
        this.deleteInspectionResult = this.deleteInspectionResult.bind(this);
    }

    deleteInspectionResult(id){
        InspectionResultService.deleteInspectionResult(id).then( res => {
            this.setState({inspectionResults: this.state.inspectionResults.filter(inspectionResult => inspectionResult.inspectionResultId !== id)});
        });
    }
    viewInspectionResult(id){
        this.props.history.push(`/view-inspectionResult/${id}`);
    }
    editInspectionResult(id){
        this.props.history.push(`/add-inspectionResult/${id}`);
    }

    componentDidMount(){
        InspectionResultService.getInspectionResults().then((res) => {
            this.setState({ inspectionResults: res.data});
        });
    }

    addInspectionResult(){
        this.props.history.push('/add-inspectionResult/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InspectionResult List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInspectionResult}> Add InspectionResult</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ResultValue </th>
                                    <th> RecordedOn </th>
                                    <th> Notes </th>
                                    <th> ResultStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inspectionResults.map(
                                        inspectionResult => 
                                        <tr key = {inspectionResult.inspectionResultId}>
                                             <td> { inspectionResult.resultValue } </td>
                                             <td> { inspectionResult.recordedOn } </td>
                                             <td> { inspectionResult.notes } </td>
                                             <td> { inspectionResult.resultStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editInspectionResult(inspectionResult.inspectionResultId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInspectionResult(inspectionResult.inspectionResultId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInspectionResult(inspectionResult.inspectionResultId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInspectionResultComponent
