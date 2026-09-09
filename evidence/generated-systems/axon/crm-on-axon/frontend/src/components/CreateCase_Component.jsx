import React, { Component } from 'react'
import Case_Service from '../services/Case_Service';

class CreateCase_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                caseNumber: '',
                subject: '',
                description: '',
                slaDue: '',
                status: '',
                priority: '',
                origin: '',
                severity: ''
        }
        this.changecaseNumberHandler = this.changecaseNumberHandler.bind(this);
        this.changesubjectHandler = this.changesubjectHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeslaDueHandler = this.changeslaDueHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
        this.changeOriginHandler = this.changeOriginHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            Case_Service.getCase_ById(this.state.id).then( (res) =>{
                let case_ = res.data;
                this.setState({
                    caseNumber: case_.caseNumber,
                    subject: case_.subject,
                    description: case_.description,
                    slaDue: case_.slaDue,
                    status: case_.status,
                    priority: case_.priority,
                    origin: case_.origin,
                    severity: case_.severity
                });
            });
        }        
    }
    saveOrUpdateCase_ = (e) => {
        e.preventDefault();
        let case_ = {
                case_Id: this.state.id,
                caseNumber: this.state.caseNumber,
                subject: this.state.subject,
                description: this.state.description,
                slaDue: this.state.slaDue,
                status: this.state.status,
                priority: this.state.priority,
                origin: this.state.origin,
                severity: this.state.severity
            };
        console.log('case_ => ' + JSON.stringify(case_));

        // step 5
        if(this.state.id === '_add'){
            case_.case_Id=''
            Case_Service.createCase_(case_).then(res =>{
                this.props.history.push('/case_s');
            });
        }else{
            Case_Service.updateCase_(case_).then( res => {
                this.props.history.push('/case_s');
            });
        }
    }
    
    changecaseNumberHandler= (event) => {
        this.setState({caseNumber: event.target.value});
    }
    changesubjectHandler= (event) => {
        this.setState({subject: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeslaDueHandler= (event) => {
        this.setState({slaDue: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }
    changeOriginHandler= (event) => {
        this.setState({origin: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }

    cancel(){
        this.props.history.push('/case_s');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Case_</h3>
        }else{
            return <h3 className="text-center">Update Case_</h3>
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
                                            <label> caseNumber:&emsp; </label>
                                                <input placeholder="caseNumber" name="caseNumber" className="form-control" value={this.state.caseNumber} onChange={this.changecaseNumberHandler}/>

                                            <label> subject:&emsp; </label>
                                                <input placeholder="subject" name="subject" className="form-control" value={this.state.subject} onChange={this.changesubjectHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> slaDue:&emsp; </label>
                                                <input type="time" placeholder="slaDue" name="slaDue" className="form-control" value={this.state.slaDue} onChange={this.changeslaDueHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          New
                      </option>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          PendingCustomer
                      </option>
                      <option name="Status" className="form-control" >
                          PendingExternal
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Resolved
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Reopened
                      </option>
                    </select>

                                            <label> Priority:&emsp; </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Low
                      </option>
                      <option name="Priority" className="form-control" >
                          Medium
                      </option>
                      <option name="Priority" className="form-control" >
                          High
                      </option>
                      <option name="Priority" className="form-control" >
                          Critical
                      </option>
                    </select>

                                            <label> Origin:&emsp; </label>
                                                <select value={this.state.origin} onChange={this.changeOriginHandler}>
                      <option name="Origin" className="form-control" >
                          Email
                      </option>
                      <option name="Origin" className="form-control" >
                          Phone
                      </option>
                      <option name="Origin" className="form-control" >
                          Web
                      </option>
                      <option name="Origin" className="form-control" >
                          Chat
                      </option>
                      <option name="Origin" className="form-control" >
                          Social
                      </option>
                      <option name="Origin" className="form-control" >
                          Community
                      </option>
                    </select>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Minor
                      </option>
                      <option name="Severity" className="form-control" >
                          Major
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                      <option name="Severity" className="form-control" >
                          Blocker
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCase_}>Save</button>
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

export default CreateCase_Component
