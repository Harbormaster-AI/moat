import React, { Component } from 'react'
import ApplicationService from '../services/ApplicationService';

class CreateApplicationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                applicationNumber: '',
                submissionDate: '',
                status: ''
        }
        this.changeapplicationNumberHandler = this.changeapplicationNumberHandler.bind(this);
        this.changesubmissionDateHandler = this.changesubmissionDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ApplicationService.getApplicationById(this.state.id).then( (res) =>{
                let application = res.data;
                this.setState({
                    applicationNumber: application.applicationNumber,
                    submissionDate: application.submissionDate,
                    status: application.status
                });
            });
        }        
    }
    saveOrUpdateApplication = (e) => {
        e.preventDefault();
        let application = {
                applicationId: this.state.id,
                applicationNumber: this.state.applicationNumber,
                submissionDate: this.state.submissionDate,
                status: this.state.status
            };
        console.log('application => ' + JSON.stringify(application));

        // step 5
        if(this.state.id === '_add'){
            application.applicationId=''
            ApplicationService.createApplication(application).then(res =>{
                this.props.history.push('/applications');
            });
        }else{
            ApplicationService.updateApplication(application).then( res => {
                this.props.history.push('/applications');
            });
        }
    }
    
    changeapplicationNumberHandler= (event) => {
        this.setState({applicationNumber: event.target.value});
    }
    changesubmissionDateHandler= (event) => {
        this.setState({submissionDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/applications');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Application</h3>
        }else{
            return <h3 className="text-center">Update Application</h3>
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
                                            <label> applicationNumber:&emsp; </label>
                                                <input placeholder="applicationNumber" name="applicationNumber" className="form-control" value={this.state.applicationNumber} onChange={this.changeapplicationNumberHandler}/>

                                            <label> submissionDate:&emsp; </label>
                                                <input type="date" placeholder="submissionDate" name="submissionDate" className="form-control" value={this.state.submissionDate} onChange={this.changesubmissionDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          UnderReview
                      </option>
                      <option name="Status" className="form-control" >
                          Quoted
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                      <option name="Status" className="form-control" >
                          Bound
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateApplication}>Save</button>
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

export default CreateApplicationComponent
