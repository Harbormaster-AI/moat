import React, { Component } from 'react'
import DirectDebitMandateService from '../services/DirectDebitMandateService';

class CreateDirectDebitMandateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                mandateId: '',
                signedAt: '',
                scheme: '',
                status: ''
        }
        this.changemandateIdHandler = this.changemandateIdHandler.bind(this);
        this.changesignedAtHandler = this.changesignedAtHandler.bind(this);
        this.changeSchemeHandler = this.changeSchemeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DirectDebitMandateService.getDirectDebitMandateById(this.state.id).then( (res) =>{
                let directDebitMandate = res.data;
                this.setState({
                    mandateId: directDebitMandate.mandateId,
                    signedAt: directDebitMandate.signedAt,
                    scheme: directDebitMandate.scheme,
                    status: directDebitMandate.status
                });
            });
        }        
    }
    saveOrUpdateDirectDebitMandate = (e) => {
        e.preventDefault();
        let directDebitMandate = {
                directDebitMandateId: this.state.id,
                mandateId: this.state.mandateId,
                signedAt: this.state.signedAt,
                scheme: this.state.scheme,
                status: this.state.status
            };
        console.log('directDebitMandate => ' + JSON.stringify(directDebitMandate));

        // step 5
        if(this.state.id === '_add'){
            directDebitMandate.directDebitMandateId=''
            DirectDebitMandateService.createDirectDebitMandate(directDebitMandate).then(res =>{
                this.props.history.push('/directDebitMandates');
            });
        }else{
            DirectDebitMandateService.updateDirectDebitMandate(directDebitMandate).then( res => {
                this.props.history.push('/directDebitMandates');
            });
        }
    }
    
    changemandateIdHandler= (event) => {
        this.setState({mandateId: event.target.value});
    }
    changesignedAtHandler= (event) => {
        this.setState({signedAt: event.target.value});
    }
    changeSchemeHandler= (event) => {
        this.setState({scheme: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/directDebitMandates');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DirectDebitMandate</h3>
        }else{
            return <h3 className="text-center">Update DirectDebitMandate</h3>
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
                                            <label> mandateId:&emsp; </label>
                                                <input placeholder="mandateId" name="mandateId" className="form-control" value={this.state.mandateId} onChange={this.changemandateIdHandler}/>

                                            <label> signedAt:&emsp; </label>
                                                <input type="time" placeholder="signedAt" name="signedAt" className="form-control" value={this.state.signedAt} onChange={this.changesignedAtHandler}/>

                                            <label> Scheme:&emsp; </label>
                                                <select value={this.state.scheme} onChange={this.changeSchemeHandler}>
                      <option name="Scheme" className="form-control" >
                          SEPA
                      </option>
                      <option name="Scheme" className="form-control" >
                          ACH
                      </option>
                      <option name="Scheme" className="form-control" >
                          BACS
                      </option>
                      <option name="Scheme" className="form-control" >
                          BECS
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDirectDebitMandate}>Save</button>
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

export default CreateDirectDebitMandateComponent
