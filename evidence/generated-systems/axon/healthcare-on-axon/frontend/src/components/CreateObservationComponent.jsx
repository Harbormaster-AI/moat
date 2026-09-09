import React, { Component } from 'react'
import ObservationService from '../services/ObservationService';

class CreateObservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                value: '',
                unit: '',
                effectiveDateTime: '',
                interpretation: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeunitHandler = this.changeunitHandler.bind(this);
        this.changeeffectiveDateTimeHandler = this.changeeffectiveDateTimeHandler.bind(this);
        this.changeInterpretationHandler = this.changeInterpretationHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ObservationService.getObservationById(this.state.id).then( (res) =>{
                let observation = res.data;
                this.setState({
                    code: observation.code,
                    value: observation.value,
                    unit: observation.unit,
                    effectiveDateTime: observation.effectiveDateTime,
                    interpretation: observation.interpretation
                });
            });
        }        
    }
    saveOrUpdateObservation = (e) => {
        e.preventDefault();
        let observation = {
                observationId: this.state.id,
                code: this.state.code,
                value: this.state.value,
                unit: this.state.unit,
                effectiveDateTime: this.state.effectiveDateTime,
                interpretation: this.state.interpretation
            };
        console.log('observation => ' + JSON.stringify(observation));

        // step 5
        if(this.state.id === '_add'){
            observation.observationId=''
            ObservationService.createObservation(observation).then(res =>{
                this.props.history.push('/observations');
            });
        }else{
            ObservationService.updateObservation(observation).then( res => {
                this.props.history.push('/observations');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }
    changeunitHandler= (event) => {
        this.setState({unit: event.target.value});
    }
    changeeffectiveDateTimeHandler= (event) => {
        this.setState({effectiveDateTime: event.target.value});
    }
    changeInterpretationHandler= (event) => {
        this.setState({interpretation: event.target.value});
    }

    cancel(){
        this.props.history.push('/observations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Observation</h3>
        }else{
            return <h3 className="text-center">Update Observation</h3>
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

                                            <label> value:&emsp; </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> unit:&emsp; </label>
                                                <input placeholder="unit" name="unit" className="form-control" value={this.state.unit} onChange={this.changeunitHandler}/>

                                            <label> effectiveDateTime:&emsp; </label>
                                                <input type="time" placeholder="effectiveDateTime" name="effectiveDateTime" className="form-control" value={this.state.effectiveDateTime} onChange={this.changeeffectiveDateTimeHandler}/>

                                            <label> Interpretation:&emsp; </label>
                                                <select value={this.state.interpretation} onChange={this.changeInterpretationHandler}>
                      <option name="Interpretation" className="form-control" >
                          Normal
                      </option>
                      <option name="Interpretation" className="form-control" >
                          AbnormalLow
                      </option>
                      <option name="Interpretation" className="form-control" >
                          AbnormalHigh
                      </option>
                      <option name="Interpretation" className="form-control" >
                          CriticalLow
                      </option>
                      <option name="Interpretation" className="form-control" >
                          CriticalHigh
                      </option>
                      <option name="Interpretation" className="form-control" >
                          Reactive
                      </option>
                      <option name="Interpretation" className="form-control" >
                          Nonreactive
                      </option>
                      <option name="Interpretation" className="form-control" >
                          Positive
                      </option>
                      <option name="Interpretation" className="form-control" >
                          Negative
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateObservation}>Save</button>
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

export default CreateObservationComponent
