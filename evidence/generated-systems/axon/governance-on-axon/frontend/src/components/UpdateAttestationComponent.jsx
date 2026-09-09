import React, { Component } from 'react'
import AttestationService from '../services/AttestationService';

class UpdateAttestationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                statement: '',
                attestor: '',
                dateSigned: '',
                result: ''
        }
        this.updateAttestation = this.updateAttestation.bind(this);

        this.changestatementHandler = this.changestatementHandler.bind(this);
        this.changeattestorHandler = this.changeattestorHandler.bind(this);
        this.changedateSignedHandler = this.changedateSignedHandler.bind(this);
        this.changeResultHandler = this.changeResultHandler.bind(this);
    }

    componentDidMount(){
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

    updateAttestation = (e) => {
        e.preventDefault();
        let attestation = {
            attestationId: this.state.id,
            statement: this.state.statement,
            attestor: this.state.attestor,
            dateSigned: this.state.dateSigned,
            result: this.state.result
        };
        console.log('attestation => ' + JSON.stringify(attestation));
        console.log('id => ' + JSON.stringify(this.state.id));
        AttestationService.updateAttestation(attestation).then( res => {
            this.props.history.push('/attestations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Attestation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> statement: </label>
                                                <input placeholder="statement" name="statement" className="form-control" value={this.state.statement} onChange={this.changestatementHandler}/>

                                            <label> attestor: </label>
                                                <input placeholder="attestor" name="attestor" className="form-control" value={this.state.attestor} onChange={this.changeattestorHandler}/>

                                            <label> dateSigned: </label>
                                                <input type="date" placeholder="dateSigned" name="dateSigned" className="form-control" value={this.state.dateSigned} onChange={this.changedateSignedHandler}/>

                                            <label> Result: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAttestation}>Save</button>
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

export default UpdateAttestationComponent
