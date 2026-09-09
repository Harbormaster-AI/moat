import React, { Component } from 'react'
import DataSubjectRequestService from '../services/DataSubjectRequestService';

class UpdateDataSubjectRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                receivedDate: '',
                dueDate: '',
                requesterCountry: '',
                requestType: '',
                status: ''
        }
        this.updateDataSubjectRequest = this.updateDataSubjectRequest.bind(this);

        this.changereceivedDateHandler = this.changereceivedDateHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
        this.changerequesterCountryHandler = this.changerequesterCountryHandler.bind(this);
        this.changeRequestTypeHandler = this.changeRequestTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        DataSubjectRequestService.getDataSubjectRequestById(this.state.id).then( (res) =>{
            let dataSubjectRequest = res.data;
            this.setState({
                receivedDate: dataSubjectRequest.receivedDate,
                dueDate: dataSubjectRequest.dueDate,
                requesterCountry: dataSubjectRequest.requesterCountry,
                requestType: dataSubjectRequest.requestType,
                status: dataSubjectRequest.status
            });
        });
    }

    updateDataSubjectRequest = (e) => {
        e.preventDefault();
        let dataSubjectRequest = {
            dataSubjectRequestId: this.state.id,
            receivedDate: this.state.receivedDate,
            dueDate: this.state.dueDate,
            requesterCountry: this.state.requesterCountry,
            requestType: this.state.requestType,
            status: this.state.status
        };
        console.log('dataSubjectRequest => ' + JSON.stringify(dataSubjectRequest));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataSubjectRequestService.updateDataSubjectRequest(dataSubjectRequest).then( res => {
            this.props.history.push('/dataSubjectRequests');
        });
    }

    changereceivedDateHandler= (event) => {
        this.setState({receivedDate: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }
    changerequesterCountryHandler= (event) => {
        this.setState({requesterCountry: event.target.value});
    }
    changeRequestTypeHandler= (event) => {
        this.setState({requestType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataSubjectRequests');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataSubjectRequest</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> receivedDate: </label>
                                                <input type="date" placeholder="receivedDate" name="receivedDate" className="form-control" value={this.state.receivedDate} onChange={this.changereceivedDateHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                            <label> requesterCountry: </label>
                                                <input placeholder="requesterCountry" name="requesterCountry" className="form-control" value={this.state.requesterCountry} onChange={this.changerequesterCountryHandler}/>

                                            <label> RequestType: </label>
                                                <select value={this.state.requestType} onChange={this.changeRequestTypeHandler}>
                      <option name="RequestType" className="form-control" >
                          Access
                      </option>
                      <option name="RequestType" className="form-control" >
                          Rectification
                      </option>
                      <option name="RequestType" className="form-control" >
                          Erasure
                      </option>
                      <option name="RequestType" className="form-control" >
                          Restriction
                      </option>
                      <option name="RequestType" className="form-control" >
                          Portability
                      </option>
                      <option name="RequestType" className="form-control" >
                          Objection
                      </option>
                      <option name="RequestType" className="form-control" >
                          AutomatedDecisioningReview
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Received
                      </option>
                      <option name="Status" className="form-control" >
                          InValidation
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Fulfilled
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataSubjectRequest}>Save</button>
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

export default UpdateDataSubjectRequestComponent
