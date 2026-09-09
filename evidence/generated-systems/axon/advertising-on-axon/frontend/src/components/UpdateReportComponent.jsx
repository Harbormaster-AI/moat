import React, { Component } from 'react'
import ReportService from '../services/ReportService';

class UpdateReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                reportName: '',
                generatedAt: '',
                fileUrl: '',
                reportType: ''
        }
        this.updateReport = this.updateReport.bind(this);

        this.changereportNameHandler = this.changereportNameHandler.bind(this);
        this.changegeneratedAtHandler = this.changegeneratedAtHandler.bind(this);
        this.changefileUrlHandler = this.changefileUrlHandler.bind(this);
        this.changeReportTypeHandler = this.changeReportTypeHandler.bind(this);
    }

    componentDidMount(){
        ReportService.getReportById(this.state.id).then( (res) =>{
            let report = res.data;
            this.setState({
                reportName: report.reportName,
                generatedAt: report.generatedAt,
                fileUrl: report.fileUrl,
                reportType: report.reportType
            });
        });
    }

    updateReport = (e) => {
        e.preventDefault();
        let report = {
            reportId: this.state.id,
            reportName: this.state.reportName,
            generatedAt: this.state.generatedAt,
            fileUrl: this.state.fileUrl,
            reportType: this.state.reportType
        };
        console.log('report => ' + JSON.stringify(report));
        console.log('id => ' + JSON.stringify(this.state.id));
        ReportService.updateReport(report).then( res => {
            this.props.history.push('/reports');
        });
    }

    changereportNameHandler= (event) => {
        this.setState({reportName: event.target.value});
    }
    changegeneratedAtHandler= (event) => {
        this.setState({generatedAt: event.target.value});
    }
    changefileUrlHandler= (event) => {
        this.setState({fileUrl: event.target.value});
    }
    changeReportTypeHandler= (event) => {
        this.setState({reportType: event.target.value});
    }

    cancel(){
        this.props.history.push('/reports');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Report</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reportName: </label>
                                                <input placeholder="reportName" name="reportName" className="form-control" value={this.state.reportName} onChange={this.changereportNameHandler}/>

                                            <label> generatedAt: </label>
                                                <input type="time" placeholder="generatedAt" name="generatedAt" className="form-control" value={this.state.generatedAt} onChange={this.changegeneratedAtHandler}/>

                                            <label> fileUrl: </label>
                                                <input placeholder="fileUrl" name="fileUrl" className="form-control" value={this.state.fileUrl} onChange={this.changefileUrlHandler}/>

                                            <label> ReportType: </label>
                                                <select value={this.state.reportType} onChange={this.changeReportTypeHandler}>
                      <option name="ReportType" className="form-control" >
                          Performance
                      </option>
                      <option name="ReportType" className="form-control" >
                          Delivery
                      </option>
                      <option name="ReportType" className="form-control" >
                          Inventory
                      </option>
                      <option name="ReportType" className="form-control" >
                          Billing
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateReport}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateReportComponent
