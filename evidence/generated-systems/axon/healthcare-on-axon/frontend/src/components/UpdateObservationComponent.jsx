import React, { Component } from 'react'
import ObservationService from '../services/ObservationService';

class UpdateObservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                value: '',
                unit: '',
                effectiveDateTime: '',
                interpretation: ''
        }
        this.updateObservation = this.updateObservation.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeunitHandler = this.changeunitHandler.bind(this);
        this.changeeffectiveDateTimeHandler = this.changeeffectiveDateTimeHandler.bind(this);
        this.changeInterpretationHandler = this.changeInterpretationHandler.bind(this);
    }

    componentDidMount(){
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

    updateObservation = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        ObservationService.updateObservation(observation).then( res => {
            this.props.history.push('/observations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Observation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> unit: </label>
                                                <input placeholder="unit" name="unit" className="form-control" value={this.state.unit} onChange={this.changeunitHandler}/>

                                            <label> effectiveDateTime: </label>
                                                <input type="time" placeholder="effectiveDateTime" name="effectiveDateTime" className="form-control" value={this.state.effectiveDateTime} onChange={this.changeeffectiveDateTimeHandler}/>

                                            <label> Interpretation: </label>
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
                                        <button className="btn btn-success" onClick={this.updateObservation}>Save</button>
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

export default UpdateObservationComponent
