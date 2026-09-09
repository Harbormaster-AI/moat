import React, { Component } from 'react'
import ImagingReportService from '../services/ImagingReportService';

class CreateImagingReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reportNumber: '',
                impression: '',
                reportedDate: '',
                status: ''
        }
        this.changereportNumberHandler = this.changereportNumberHandler.bind(this);
        this.changeimpressionHandler = this.changeimpressionHandler.bind(this);
        this.changereportedDateHandler = this.changereportedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateImagingReport = (e) => {
        e.preventDefault();
        let imagingReport = {
                imagingReportId: this.state.id,
                reportNumber: this.state.reportNumber,
                impression: this.state.impression,
                reportedDate: this.state.reportedDate,
                status: this.state.status
            };
        console.log('imagingReport => ' + JSON.stringify(imagingReport));

        // step 5
        if(this.state.id === '_add'){
            imagingReport.imagingReportId=''
            ImagingReportService.createImagingReport(imagingReport).then(res =>{
                this.props.history.push('/imagingReports');
            });
        }else{
            ImagingReportService.updateImagingReport(imagingReport).then( res => {
                this.props.history.push('/imagingReports');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ImagingReport</h3>
        }else{
            return <h3 className="text-center">Update ImagingReport</h3>
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
                                            <label> reportNumber:&emsp; </label>
                                                <input placeholder="reportNumber" name="reportNumber" className="form-control" value={this.state.reportNumber} onChange={this.changereportNumberHandler}/>

                                            <label> impression:&emsp; </label>
                                                <input placeholder="impression" name="impression" className="form-control" value={this.state.impression} onChange={this.changeimpressionHandler}/>

                                            <label> reportedDate:&emsp; </label>
                                                <input type="time" placeholder="reportedDate" name="reportedDate" className="form-control" value={this.state.reportedDate} onChange={this.changereportedDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateImagingReport}>Save</button>
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

export default CreateImagingReportComponent
