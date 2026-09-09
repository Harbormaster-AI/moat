import React, { Component } from 'react'
import Exception_Service from '../services/Exception_Service';

class CreateException_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                justification: '',
                startDate: '',
                endDate: '',
                exceptionType: '',
                status: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changejustificationHandler = this.changejustificationHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeExceptionTypeHandler = this.changeExceptionTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            Exception_Service.getException_ById(this.state.id).then( (res) =>{
                let exception_ = res.data;
                this.setState({
                    title: exception_.title,
                    justification: exception_.justification,
                    startDate: exception_.startDate,
                    endDate: exception_.endDate,
                    exceptionType: exception_.exceptionType,
                    status: exception_.status
                });
            });
        }        
    }
    saveOrUpdateException_ = (e) => {
        e.preventDefault();
        let exception_ = {
                exception_Id: this.state.id,
                title: this.state.title,
                justification: this.state.justification,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                exceptionType: this.state.exceptionType,
                status: this.state.status
            };
        console.log('exception_ => ' + JSON.stringify(exception_));

        // step 5
        if(this.state.id === '_add'){
            exception_.exception_Id=''
            Exception_Service.createException_(exception_).then(res =>{
                this.props.history.push('/exception_s');
            });
        }else{
            Exception_Service.updateException_(exception_).then( res => {
                this.props.history.push('/exception_s');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changejustificationHandler= (event) => {
        this.setState({justification: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeExceptionTypeHandler= (event) => {
        this.setState({exceptionType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/exception_s');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Exception_</h3>
        }else{
            return <h3 className="text-center">Update Exception_</h3>
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

                                            <label> justification:&emsp; </label>
                                                <input placeholder="justification" name="justification" className="form-control" value={this.state.justification} onChange={this.changejustificationHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> ExceptionType:&emsp; </label>
                                                <select value={this.state.exceptionType} onChange={this.changeExceptionTypeHandler}>
                      <option name="ExceptionType" className="form-control" >
                          PolicyException
                      </option>
                      <option name="ExceptionType" className="form-control" >
                          ControlException
                      </option>
                      <option name="ExceptionType" className="form-control" >
                          RetentionException
                      </option>
                      <option name="ExceptionType" className="form-control" >
                          RiskAcceptance
                      </option>
                      <option name="ExceptionType" className="form-control" >
                          ComplianceWaiver
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Submitted
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateException_}>Save</button>
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

export default CreateException_Component
