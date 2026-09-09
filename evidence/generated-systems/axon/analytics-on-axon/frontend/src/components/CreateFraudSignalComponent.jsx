import React, { Component } from 'react'
import FraudSignalService from '../services/FraudSignalService';

class CreateFraudSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                ruleLogic: '',
                signalType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeruleLogicHandler = this.changeruleLogicHandler.bind(this);
        this.changeSignalTypeHandler = this.changeSignalTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FraudSignalService.getFraudSignalById(this.state.id).then( (res) =>{
                let fraudSignal = res.data;
                this.setState({
                    name: fraudSignal.name,
                    ruleLogic: fraudSignal.ruleLogic,
                    signalType: fraudSignal.signalType
                });
            });
        }        
    }
    saveOrUpdateFraudSignal = (e) => {
        e.preventDefault();
        let fraudSignal = {
                fraudSignalId: this.state.id,
                name: this.state.name,
                ruleLogic: this.state.ruleLogic,
                signalType: this.state.signalType
            };
        console.log('fraudSignal => ' + JSON.stringify(fraudSignal));

        // step 5
        if(this.state.id === '_add'){
            fraudSignal.fraudSignalId=''
            FraudSignalService.createFraudSignal(fraudSignal).then(res =>{
                this.props.history.push('/fraudSignals');
            });
        }else{
            FraudSignalService.updateFraudSignal(fraudSignal).then( res => {
                this.props.history.push('/fraudSignals');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeruleLogicHandler= (event) => {
        this.setState({ruleLogic: event.target.value});
    }
    changeSignalTypeHandler= (event) => {
        this.setState({signalType: event.target.value});
    }

    cancel(){
        this.props.history.push('/fraudSignals');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FraudSignal</h3>
        }else{
            return <h3 className="text-center">Update FraudSignal</h3>
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

                                            <label> ruleLogic:&emsp; </label>
                                                <input placeholder="ruleLogic" name="ruleLogic" className="form-control" value={this.state.ruleLogic} onChange={this.changeruleLogicHandler}/>

                                            <label> SignalType:&emsp; </label>
                                                <select value={this.state.signalType} onChange={this.changeSignalTypeHandler}>
                      <option name="SignalType" className="form-control" >
                          Velocity
                      </option>
                      <option name="SignalType" className="form-control" >
                          GeolocationMismatch
                      </option>
                      <option name="SignalType" className="form-control" >
                          AmountOutlier
                      </option>
                      <option name="SignalType" className="form-control" >
                          DeviceFingerprint
                      </option>
                      <option name="SignalType" className="form-control" >
                          BehavioralChange
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFraudSignal}>Save</button>
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

export default CreateFraudSignalComponent
