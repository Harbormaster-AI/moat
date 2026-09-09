import React, { Component } from 'react'
import WorkAuthorizationService from '../services/WorkAuthorizationService';

class UpdateWorkAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                country: '',
                expirationDate: '',
                status: ''
        }
        this.updateWorkAuthorization = this.updateWorkAuthorization.bind(this);

        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        WorkAuthorizationService.getWorkAuthorizationById(this.state.id).then( (res) =>{
            let workAuthorization = res.data;
            this.setState({
                country: workAuthorization.country,
                expirationDate: workAuthorization.expirationDate,
                status: workAuthorization.status
            });
        });
    }

    updateWorkAuthorization = (e) => {
        e.preventDefault();
        let workAuthorization = {
            workAuthorizationId: this.state.id,
            country: this.state.country,
            expirationDate: this.state.expirationDate,
            status: this.state.status
        };
        console.log('workAuthorization => ' + JSON.stringify(workAuthorization));
        console.log('id => ' + JSON.stringify(this.state.id));
        WorkAuthorizationService.updateWorkAuthorization(workAuthorization).then( res => {
            this.props.history.push('/workAuthorizations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WorkAuthorization</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> country: </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> expirationDate: </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateWorkAuthorization}>Save</button>
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

export default UpdateWorkAuthorizationComponent
