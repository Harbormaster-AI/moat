import React, { Component } from 'react'
import DiagnosisService from '../services/DiagnosisService';

class CreateDiagnosisComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                description: '',
                onsetDate: '',
                certainty: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeonsetDateHandler = this.changeonsetDateHandler.bind(this);
        this.changeCertaintyHandler = this.changeCertaintyHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateDiagnosis = (e) => {
        e.preventDefault();
        let diagnosis = {
                diagnosisId: this.state.id,
                code: this.state.code,
                description: this.state.description,
                onsetDate: this.state.onsetDate,
                certainty: this.state.certainty
            };
        console.log('diagnosis => ' + JSON.stringify(diagnosis));

        // step 5
        if(this.state.id === '_add'){
            diagnosis.diagnosisId=''
            DiagnosisService.createDiagnosis(diagnosis).then(res =>{
                this.props.history.push('/diagnosiss');
            });
        }else{
            DiagnosisService.updateDiagnosis(diagnosis).then( res => {
                this.props.history.push('/diagnosiss');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Diagnosis</h3>
        }else{
            return <h3 className="text-center">Update Diagnosis</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> onsetDate:&emsp; </label>
                                                <input type="date" placeholder="onsetDate" name="onsetDate" className="form-control" value={this.state.onsetDate} onChange={this.changeonsetDateHandler}/>

                                            <label> Certainty:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDiagnosis}>Save</button>
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

export default CreateDiagnosisComponent
