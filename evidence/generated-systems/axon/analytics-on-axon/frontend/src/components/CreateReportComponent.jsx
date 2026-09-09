import React, { Component } from 'react'
import ReportService from '../services/ReportService';

class CreateReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                audience: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeaudienceHandler = this.changeaudienceHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
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
                    title: report.title,
                    audience: report.audience,
                    status: report.status
                });
            });
        }        
    }
    saveOrUpdateReport = (e) => {
        e.preventDefault();
        let report = {
                reportId: this.state.id,
                title: this.state.title,
                audience: this.state.audience,
                status: this.state.status
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
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeaudienceHandler= (event) => {
        this.setState({audience: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> audience:&emsp; </label>
                                                <input placeholder="audience" name="audience" className="form-control" value={this.state.audience} onChange={this.changeaudienceHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Published
                      </option>
                      <option name="Status" className="form-control" >
                          Archived
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
