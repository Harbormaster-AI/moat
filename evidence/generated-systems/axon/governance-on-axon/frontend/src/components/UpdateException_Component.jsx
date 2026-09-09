import React, { Component } from 'react'
import Exception_Service from '../services/Exception_Service';

class UpdateException_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                title: '',
                justification: '',
                startDate: '',
                endDate: '',
                exceptionType: '',
                status: ''
        }
        this.updateException_ = this.updateException_.bind(this);

        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changejustificationHandler = this.changejustificationHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeExceptionTypeHandler = this.changeExceptionTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateException_ = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        Exception_Service.updateException_(exception_).then( res => {
            this.props.history.push('/exception_s');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Exception_</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> justification: </label>
                                                <input placeholder="justification" name="justification" className="form-control" value={this.state.justification} onChange={this.changejustificationHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> ExceptionType: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateException_}>Save</button>
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

export default UpdateException_Component
