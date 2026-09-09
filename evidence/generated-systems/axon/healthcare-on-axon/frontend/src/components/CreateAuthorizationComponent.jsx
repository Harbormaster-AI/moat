import React, { Component } from 'react'
import AuthorizationService from '../services/AuthorizationService';

class CreateAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                authNumber: '',
                requestedService: '',
                status: ''
        }
        this.changeauthNumberHandler = this.changeauthNumberHandler.bind(this);
        this.changerequestedServiceHandler = this.changerequestedServiceHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AuthorizationService.getAuthorizationById(this.state.id).then( (res) =>{
                let authorization = res.data;
                this.setState({
                    authNumber: authorization.authNumber,
                    requestedService: authorization.requestedService,
                    status: authorization.status
                });
            });
        }        
    }
    saveOrUpdateAuthorization = (e) => {
        e.preventDefault();
        let authorization = {
                authorizationId: this.state.id,
                authNumber: this.state.authNumber,
                requestedService: this.state.requestedService,
                status: this.state.status
            };
        console.log('authorization => ' + JSON.stringify(authorization));

        // step 5
        if(this.state.id === '_add'){
            authorization.authorizationId=''
            AuthorizationService.createAuthorization(authorization).then(res =>{
                this.props.history.push('/authorizations');
            });
        }else{
            AuthorizationService.updateAuthorization(authorization).then( res => {
                this.props.history.push('/authorizations');
            });
        }
    }
    
    changeauthNumberHandler= (event) => {
        this.setState({authNumber: event.target.value});
    }
    changerequestedServiceHandler= (event) => {
        this.setState({requestedService: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/authorizations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Authorization</h3>
        }else{
            return <h3 className="text-center">Update Authorization</h3>
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
                                            <label> authNumber:&emsp; </label>
                                                <input placeholder="authNumber" name="authNumber" className="form-control" value={this.state.authNumber} onChange={this.changeauthNumberHandler}/>

                                            <label> requestedService:&emsp; </label>
                                                <input placeholder="requestedService" name="requestedService" className="form-control" value={this.state.requestedService} onChange={this.changerequestedServiceHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Requested
                      </option>
                      <option name="Status" className="form-control" >
                          PendingReview
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Denied
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAuthorization}>Save</button>
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

export default CreateAuthorizationComponent
