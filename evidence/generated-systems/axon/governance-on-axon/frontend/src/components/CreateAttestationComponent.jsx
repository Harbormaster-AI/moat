import React, { Component } from 'react'
import AttestationService from '../services/AttestationService';

class CreateAttestationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                statement: '',
                attestor: '',
                dateSigned: '',
                result: ''
        }
        this.changestatementHandler = this.changestatementHandler.bind(this);
        this.changeattestorHandler = this.changeattestorHandler.bind(this);
        this.changedateSignedHandler = this.changedateSignedHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AttestationService.getAttestationById(this.state.id).then( (res) =>{
                let attestation = res.data;
                this.setState({
                    statement: attestation.statement,
                    attestor: attestation.attestor,
                    dateSigned: attestation.dateSigned,
                    result: attestation.result
                });
            });
        }        
    }
    saveOrUpdateAttestation = (e) => {
        e.preventDefault();
        let attestation = {
                attestationId: this.state.id,
                statement: this.state.statement,
                attestor: this.state.attestor,
                dateSigned: this.state.dateSigned,
                result: this.state.result
            };
        console.log('attestation => ' + JSON.stringify(attestation));

        // step 5
        if(this.state.id === '_add'){
            attestation.attestationId=''
            AttestationService.createAttestation(attestation).then(res =>{
                this.props.history.push('/attestations');
            });
        }else{
            AttestationService.updateAttestation(attestation).then( res => {
                this.props.history.push('/attestations');
            });
        }
    }
    
    changestatementHandler= (event) => {
        this.setState({statement: event.target.value});
    }
    changeattestorHandler= (event) => {
        this.setState({attestor: event.target.value});
    }
    changedateSignedHandler= (event) => {
        this.setState({dateSigned: event.target.value});
    }
    changeResultHandler= (event) => {
        this.setState({result: event.target.value});
    }

    cancel(){
        this.props.history.push('/attestations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Attestation</h3>
        }else{
            return <h3 className="text-center">Update Attestation</h3>
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
                                            <label> statement:&emsp; </label>
                                                <input placeholder="statement" name="statement" className="form-control" value={this.state.statement} onChange={this.changestatementHandler}/>

                                            <label> attestor:&emsp; </label>
                                                <input placeholder="attestor" name="attestor" className="form-control" value={this.state.attestor} onChange={this.changeattestorHandler}/>

                                            <label> dateSigned:&emsp; </label>
                                                <input type="date" placeholder="dateSigned" name="dateSigned" className="form-control" value={this.state.dateSigned} onChange={this.changedateSignedHandler}/>

                                            <label> Result:&emsp; </label>
                                                <select value={this.state.result} onChange={this.changeResultHandler}>
                      <option name="Result" className="form-control" >
                          Affirmative
                      </option>
                      <option name="Result" className="form-control" >
                          Negative
                      </option>
                      <option name="Result" className="form-control" >
                          Qualified
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAttestation}>Save</button>
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

export default CreateAttestationComponent
