import React, { Component } from 'react'
import PolicyAcknowledgementService from '../services/PolicyAcknowledgementService';

class UpdatePolicyAcknowledgementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                acknowledgementDate: '',
                status: ''
        }
        this.updatePolicyAcknowledgement = this.updatePolicyAcknowledgement.bind(this);

        this.changeacknowledgementDateHandler = this.changeacknowledgementDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PolicyAcknowledgementService.getPolicyAcknowledgementById(this.state.id).then( (res) =>{
            let policyAcknowledgement = res.data;
            this.setState({
                acknowledgementDate: policyAcknowledgement.acknowledgementDate,
                status: policyAcknowledgement.status
            });
        });
    }

    updatePolicyAcknowledgement = (e) => {
        e.preventDefault();
        let policyAcknowledgement = {
            policyAcknowledgementId: this.state.id,
            acknowledgementDate: this.state.acknowledgementDate,
            status: this.state.status
        };
        console.log('policyAcknowledgement => ' + JSON.stringify(policyAcknowledgement));
        console.log('id => ' + JSON.stringify(this.state.id));
        PolicyAcknowledgementService.updatePolicyAcknowledgement(policyAcknowledgement).then( res => {
            this.props.history.push('/policyAcknowledgements');
        });
    }

    changeacknowledgementDateHandler= (event) => {
        this.setState({acknowledgementDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/policyAcknowledgements');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PolicyAcknowledgement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> acknowledgementDate: </label>
                                                <input type="date" placeholder="acknowledgementDate" name="acknowledgementDate" className="form-control" value={this.state.acknowledgementDate} onChange={this.changeacknowledgementDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Acknowledged
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePolicyAcknowledgement}>Save</button>
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

export default UpdatePolicyAcknowledgementComponent
