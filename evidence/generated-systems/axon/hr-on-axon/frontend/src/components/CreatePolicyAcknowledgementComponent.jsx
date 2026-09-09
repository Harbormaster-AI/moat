import React, { Component } from 'react'
import PolicyAcknowledgementService from '../services/PolicyAcknowledgementService';

class CreatePolicyAcknowledgementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                acknowledgementDate: '',
                status: ''
        }
        this.changeacknowledgementDateHandler = this.changeacknowledgementDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PolicyAcknowledgementService.getPolicyAcknowledgementById(this.state.id).then( (res) =>{
                let policyAcknowledgement = res.data;
                this.setState({
                    acknowledgementDate: policyAcknowledgement.acknowledgementDate,
                    status: policyAcknowledgement.status
                });
            });
        }        
    }
    saveOrUpdatePolicyAcknowledgement = (e) => {
        e.preventDefault();
        let policyAcknowledgement = {
                policyAcknowledgementId: this.state.id,
                acknowledgementDate: this.state.acknowledgementDate,
                status: this.state.status
            };
        console.log('policyAcknowledgement => ' + JSON.stringify(policyAcknowledgement));

        // step 5
        if(this.state.id === '_add'){
            policyAcknowledgement.policyAcknowledgementId=''
            PolicyAcknowledgementService.createPolicyAcknowledgement(policyAcknowledgement).then(res =>{
                this.props.history.push('/policyAcknowledgements');
            });
        }else{
            PolicyAcknowledgementService.updatePolicyAcknowledgement(policyAcknowledgement).then( res => {
                this.props.history.push('/policyAcknowledgements');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PolicyAcknowledgement</h3>
        }else{
            return <h3 className="text-center">Update PolicyAcknowledgement</h3>
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
                                            <label> acknowledgementDate:&emsp; </label>
                                                <input type="date" placeholder="acknowledgementDate" name="acknowledgementDate" className="form-control" value={this.state.acknowledgementDate} onChange={this.changeacknowledgementDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePolicyAcknowledgement}>Save</button>
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

export default CreatePolicyAcknowledgementComponent
