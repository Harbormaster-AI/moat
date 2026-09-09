import React, { Component } from 'react'
import AuditEngagementService from '../services/AuditEngagementService';

class CreateAuditEngagementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                startDate: '',
                endDate: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AuditEngagementService.getAuditEngagementById(this.state.id).then( (res) =>{
                let auditEngagement = res.data;
                this.setState({
                    title: auditEngagement.title,
                    startDate: auditEngagement.startDate,
                    endDate: auditEngagement.endDate,
                    status: auditEngagement.status
                });
            });
        }        
    }
    saveOrUpdateAuditEngagement = (e) => {
        e.preventDefault();
        let auditEngagement = {
                auditEngagementId: this.state.id,
                title: this.state.title,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                status: this.state.status
            };
        console.log('auditEngagement => ' + JSON.stringify(auditEngagement));

        // step 5
        if(this.state.id === '_add'){
            auditEngagement.auditEngagementId=''
            AuditEngagementService.createAuditEngagement(auditEngagement).then(res =>{
                this.props.history.push('/auditEngagements');
            });
        }else{
            AuditEngagementService.updateAuditEngagement(auditEngagement).then( res => {
                this.props.history.push('/auditEngagements');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/auditEngagements');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AuditEngagement</h3>
        }else{
            return <h3 className="text-center">Update AuditEngagement</h3>
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

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Fieldwork
                      </option>
                      <option name="Status" className="form-control" >
                          Reporting
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAuditEngagement}>Save</button>
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

export default CreateAuditEngagementComponent
