import React, { Component } from 'react'
import AuditFindingService from '../services/AuditFindingService';

class CreateAuditFindingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                description: '',
                dueDate: '',
                severity: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AuditFindingService.getAuditFindingById(this.state.id).then( (res) =>{
                let auditFinding = res.data;
                this.setState({
                    title: auditFinding.title,
                    description: auditFinding.description,
                    dueDate: auditFinding.dueDate,
                    severity: auditFinding.severity,
                    status: auditFinding.status
                });
            });
        }        
    }
    saveOrUpdateAuditFinding = (e) => {
        e.preventDefault();
        let auditFinding = {
                auditFindingId: this.state.id,
                title: this.state.title,
                description: this.state.description,
                dueDate: this.state.dueDate,
                severity: this.state.severity,
                status: this.state.status
            };
        console.log('auditFinding => ' + JSON.stringify(auditFinding));

        // step 5
        if(this.state.id === '_add'){
            auditFinding.auditFindingId=''
            AuditFindingService.createAuditFinding(auditFinding).then(res =>{
                this.props.history.push('/auditFindings');
            });
        }else{
            AuditFindingService.updateAuditFinding(auditFinding).then( res => {
                this.props.history.push('/auditFindings');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/auditFindings');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AuditFinding</h3>
        }else{
            return <h3 className="text-center">Update AuditFinding</h3>
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

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> dueDate:&emsp; </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Low
                      </option>
                      <option name="Severity" className="form-control" >
                          Medium
                      </option>
                      <option name="Severity" className="form-control" >
                          High
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          InRemediation
                      </option>
                      <option name="Status" className="form-control" >
                          Validated
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAuditFinding}>Save</button>
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

export default CreateAuditFindingComponent
