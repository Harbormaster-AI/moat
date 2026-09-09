import React, { Component } from 'react'
import DiagnosisService from '../services/DiagnosisService';

class UpdateDiagnosisComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                description: '',
                onsetDate: '',
                certainty: ''
        }
        this.updateDiagnosis = this.updateDiagnosis.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeonsetDateHandler = this.changeonsetDateHandler.bind(this);
        this.changeCertaintyHandler = this.changeCertaintyHandler.bind(this);
    }

    componentDidMount(){
        DiagnosisService.getDiagnosisById(this.state.id).then( (res) =>{
            let diagnosis = res.data;
            this.setState({
                code: diagnosis.code,
                description: diagnosis.description,
                onsetDate: diagnosis.onsetDate,
                certainty: diagnosis.certainty
            });
        });
    }

    updateDiagnosis = (e) => {
        e.preventDefault();
        let diagnosis = {
            diagnosisId: this.state.id,
            code: this.state.code,
            description: this.state.description,
            onsetDate: this.state.onsetDate,
            certainty: this.state.certainty
        };
        console.log('diagnosis => ' + JSON.stringify(diagnosis));
        console.log('id => ' + JSON.stringify(this.state.id));
        DiagnosisService.updateDiagnosis(diagnosis).then( res => {
            this.props.history.push('/diagnosiss');
        });
    }

    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeonsetDateHandler= (event) => {
        this.setState({onsetDate: event.target.value});
    }
    changeCertaintyHandler= (event) => {
        this.setState({certainty: event.target.value});
    }

    cancel(){
        this.props.history.push('/diagnosiss');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Diagnosis</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> onsetDate: </label>
                                                <input type="date" placeholder="onsetDate" name="onsetDate" className="form-control" value={this.state.onsetDate} onChange={this.changeonsetDateHandler}/>

                                            <label> Certainty: </label>
                                                <select value={this.state.certainty} onChange={this.changeCertaintyHandler}>
                      <option name="Certainty" className="form-control" >
                          Suspected
                      </option>
                      <option name="Certainty" className="form-control" >
                          Presumptive
                      </option>
                      <option name="Certainty" className="form-control" >
                          Confirmed
                      </option>
                      <option name="Certainty" className="form-control" >
                          RuledOut
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDiagnosis}>Save</button>
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

export default UpdateDiagnosisComponent
