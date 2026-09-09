import React, { Component } from 'react'
import ImagingReportService from '../services/ImagingReportService'

class ListImagingReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                imagingReports: []
        }
        this.addImagingReport = this.addImagingReport.bind(this);
        this.editImagingReport = this.editImagingReport.bind(this);
        this.deleteImagingReport = this.deleteImagingReport.bind(this);
    }

    deleteImagingReport(id){
        ImagingReportService.deleteImagingReport(id).then( res => {
            this.setState({imagingReports: this.state.imagingReports.filter(imagingReport => imagingReport.imagingReportId !== id)});
        });
    }
    viewImagingReport(id){
        this.props.history.push(`/view-imagingReport/${id}`);
    }
    editImagingReport(id){
        this.props.history.push(`/add-imagingReport/${id}`);
    }

    componentDidMount(){
        ImagingReportService.getImagingReports().then((res) => {
            this.setState({ imagingReports: res.data});
        });
    }

    addImagingReport(){
        this.props.history.push('/add-imagingReport/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ImagingReport List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addImagingReport}> Add ImagingReport</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReportNumber </th>
                                    <th> Impression </th>
                                    <th> ReportedDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.imagingReports.map(
                                        imagingReport => 
                                        <tr key = {imagingReport.imagingReportId}>
                                             <td> { imagingReport.reportNumber } </td>
                                             <td> { imagingReport.impression } </td>
                                             <td> { imagingReport.reportedDate } </td>
                                             <td> { imagingReport.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editImagingReport(imagingReport.imagingReportId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteImagingReport(imagingReport.imagingReportId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewImagingReport(imagingReport.imagingReportId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListImagingReportComponent
