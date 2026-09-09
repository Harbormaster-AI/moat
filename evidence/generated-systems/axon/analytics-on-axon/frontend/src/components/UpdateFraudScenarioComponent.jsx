import React, { Component } from 'react'
import FraudScenarioService from '../services/FraudScenarioService';

class UpdateFraudScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                riskAppetite: '',
                detectionType: ''
        }
        this.updateFraudScenario = this.updateFraudScenario.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeriskAppetiteHandler = this.changeriskAppetiteHandler.bind(this);
        this.changeDetectionTypeHandler = this.changeDetectionTypeHandler.bind(this);
    }

    componentDidMount(){
        FraudScenarioService.getFraudScenarioById(this.state.id).then( (res) =>{
            let fraudScenario = res.data;
            this.setState({
                name: fraudScenario.name,
                riskAppetite: fraudScenario.riskAppetite,
                detectionType: fraudScenario.detectionType
            });
        });
    }

    updateFraudScenario = (e) => {
        e.preventDefault();
        let fraudScenario = {
            fraudScenarioId: this.state.id,
            name: this.state.name,
            riskAppetite: this.state.riskAppetite,
            detectionType: this.state.detectionType
        };
        console.log('fraudScenario => ' + JSON.stringify(fraudScenario));
        console.log('id => ' + JSON.stringify(this.state.id));
        FraudScenarioService.updateFraudScenario(fraudScenario).then( res => {
            this.props.history.push('/fraudScenarios');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeriskAppetiteHandler= (event) => {
        this.setState({riskAppetite: event.target.value});
    }
    changeDetectionTypeHandler= (event) => {
        this.setState({detectionType: event.target.value});
    }

    cancel(){
        this.props.history.push('/fraudScenarios');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update FraudScenario</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> riskAppetite: </label>
                                                <input placeholder="riskAppetite" name="riskAppetite" className="form-control" value={this.state.riskAppetite} onChange={this.changeriskAppetiteHandler}/>

                                            <label> DetectionType: </label>
                                                <select value={this.state.detectionType} onChange={this.changeDetectionTypeHandler}>
                      <option name="DetectionType" className="form-control" >
                          RuleBased
                      </option>
                      <option name="DetectionType" className="form-control" >
                          SupervisedML
                      </option>
                      <option name="DetectionType" className="form-control" >
                          UnsupervisedML
                      </option>
                      <option name="DetectionType" className="form-control" >
                          Hybrid
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFraudScenario}>Save</button>
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

export default UpdateFraudScenarioComponent
