import React, { Component } from 'react'
import ReportService from '../services/ReportService';

class CreateReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reportName: '',
                generatedAt: '',
                fileUrl: '',
                reportType: ''
        }
        this.changereportNameHandler = this.changereportNameHandler.bind(this);
        this.changegeneratedAtHandler = this.changegeneratedAtHandler.bind(this);
        this.changefileUrlHandler = this.changefileUrlHandler.bind(this);
        this.changeReportTypeHandler = this.changeReportTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateReport = (e) => {
        e.preventDefault();
        let report = {
                reportId: this.state.id,
                reportName: this.state.reportName,
                generatedAt: this.state.generatedAt,
                fileUrl: this.state.fileUrl,
                reportType: this.state.reportType
            };
        console.log('report => ' + JSON.stringify(report));

        // step 5
        if(this.state.id === '_add'){
            report.reportId=''
            ReportService.createReport(report).then(res =>{
                this.props.history.push('/reports');
            });
        }else{
            ReportService.updateReport(report).then( res => {
                this.props.history.push('/reports');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Report</h3>
        }else{
            return <h3 className="text-center">Update Report</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reportName:&emsp; </label>
                                                <input placeholder="reportName" name="reportName" className="form-control" value={this.state.reportName} onChange={this.changereportNameHandler}/>

                                            <label> generatedAt:&emsp; </label>
                                                <input type="time" placeholder="generatedAt" name="generatedAt" className="form-control" value={this.state.generatedAt} onChange={this.changegeneratedAtHandler}/>

                                            <label> fileUrl:&emsp; </label>
                                                <input placeholder="fileUrl" name="fileUrl" className="form-control" value={this.state.fileUrl} onChange={this.changefileUrlHandler}/>

                                            <label> ReportType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReport}>Save</button>
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

export default CreateReportComponent
