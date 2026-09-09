import React, { Component } from 'react'
import FraudScenarioService from '../services/FraudScenarioService';

class CreateFraudScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                riskAppetite: '',
                detectionType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeriskAppetiteHandler = this.changeriskAppetiteHandler.bind(this);
        this.changeDetectionTypeHandler = this.changeDetectionTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FraudScenarioService.getFraudScenarioById(this.state.id).then( (res) =>{
                let fraudScenario = res.data;
                this.setState({
                    name: fraudScenario.name,
                    riskAppetite: fraudScenario.riskAppetite,
                    detectionType: fraudScenario.detectionType
                });
            });
        }        
    }
    saveOrUpdateFraudScenario = (e) => {
        e.preventDefault();
        let fraudScenario = {
                fraudScenarioId: this.state.id,
                name: this.state.name,
                riskAppetite: this.state.riskAppetite,
                detectionType: this.state.detectionType
            };
        console.log('fraudScenario => ' + JSON.stringify(fraudScenario));

        // step 5
        if(this.state.id === '_add'){
            fraudScenario.fraudScenarioId=''
            FraudScenarioService.createFraudScenario(fraudScenario).then(res =>{
                this.props.history.push('/fraudScenarios');
            });
        }else{
            FraudScenarioService.updateFraudScenario(fraudScenario).then( res => {
                this.props.history.push('/fraudScenarios');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FraudScenario</h3>
        }else{
            return <h3 className="text-center">Update FraudScenario</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> riskAppetite:&emsp; </label>
                                                <input placeholder="riskAppetite" name="riskAppetite" className="form-control" value={this.state.riskAppetite} onChange={this.changeriskAppetiteHandler}/>

                                            <label> DetectionType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFraudScenario}>Save</button>
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

export default CreateFraudScenarioComponent
