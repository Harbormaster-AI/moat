import React, { Component } from 'react'
import WorkAuthorizationService from '../services/WorkAuthorizationService';

class CreateWorkAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                country: '',
                expirationDate: '',
                status: ''
        }
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WorkAuthorizationService.getWorkAuthorizationById(this.state.id).then( (res) =>{
                let workAuthorization = res.data;
                this.setState({
                    country: workAuthorization.country,
                    expirationDate: workAuthorization.expirationDate,
                    status: workAuthorization.status
                });
            });
        }        
    }
    saveOrUpdateWorkAuthorization = (e) => {
        e.preventDefault();
        let workAuthorization = {
                workAuthorizationId: this.state.id,
                country: this.state.country,
                expirationDate: this.state.expirationDate,
                status: this.state.status
            };
        console.log('workAuthorization => ' + JSON.stringify(workAuthorization));

        // step 5
        if(this.state.id === '_add'){
            workAuthorization.workAuthorizationId=''
            WorkAuthorizationService.createWorkAuthorization(workAuthorization).then(res =>{
                this.props.history.push('/workAuthorizations');
            });
        }else{
            WorkAuthorizationService.updateWorkAuthorization(workAuthorization).then( res => {
                this.props.history.push('/workAuthorizations');
            });
        }
    }
    
    changecountryHandler= (event) => {
        this.setState({country: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/workAuthorizations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add WorkAuthorization</h3>
        }else{
            return <h3 className="text-center">Update WorkAuthorization</h3>
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
                                            <label> country:&emsp; </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> expirationDate:&emsp; </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotRequired
                      </option>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Authorized
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWorkAuthorization}>Save</button>
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

export default CreateWorkAuthorizationComponent
