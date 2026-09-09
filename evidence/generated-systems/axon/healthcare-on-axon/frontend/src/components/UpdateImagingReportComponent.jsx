import React, { Component } from 'react'
import ImagingReportService from '../services/ImagingReportService';

class UpdateImagingReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                reportNumber: '',
                impression: '',
                reportedDate: '',
                status: ''
        }
        this.updateImagingReport = this.updateImagingReport.bind(this);

        this.changereportNumberHandler = this.changereportNumberHandler.bind(this);
        this.changeimpressionHandler = this.changeimpressionHandler.bind(this);
        this.changereportedDateHandler = this.changereportedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ImagingReportService.getImagingReportById(this.state.id).then( (res) =>{
            let imagingReport = res.data;
            this.setState({
                reportNumber: imagingReport.reportNumber,
                impression: imagingReport.impression,
                reportedDate: imagingReport.reportedDate,
                status: imagingReport.status
            });
        });
    }

    updateImagingReport = (e) => {
        e.preventDefault();
        let imagingReport = {
            imagingReportId: this.state.id,
            reportNumber: this.state.reportNumber,
            impression: this.state.impression,
            reportedDate: this.state.reportedDate,
            status: this.state.status
        };
        console.log('imagingReport => ' + JSON.stringify(imagingReport));
        console.log('id => ' + JSON.stringify(this.state.id));
        ImagingReportService.updateImagingReport(imagingReport).then( res => {
            this.props.history.push('/imagingReports');
        });
    }

    changereportNumberHandler= (event) => {
        this.setState({reportNumber: event.target.value});
    }
    changeimpressionHandler= (event) => {
        this.setState({impression: event.target.value});
    }
    changereportedDateHandler= (event) => {
        this.setState({reportedDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/imagingReports');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ImagingReport</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reportNumber: </label>
                                                <input placeholder="reportNumber" name="reportNumber" className="form-control" value={this.state.reportNumber} onChange={this.changereportNumberHandler}/>

                                            <label> impression: </label>
                                                <input placeholder="impression" name="impression" className="form-control" value={this.state.impression} onChange={this.changeimpressionHandler}/>

                                            <label> reportedDate: </label>
                                                <input type="time" placeholder="reportedDate" name="reportedDate" className="form-control" value={this.state.reportedDate} onChange={this.changereportedDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Registered
                      </option>
                      <option name="Status" className="form-control" >
                          Partial
                      </option>
                      <option name="Status" className="form-control" >
                          Final
                      </option>
                      <option name="Status" className="form-control" >
                          Corrected
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateImagingReport}>Save</button>
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

export default UpdateImagingReportComponent
