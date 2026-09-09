import React, { Component } from 'react'
import ControlTest_Service from '../services/ControlTest_Service';

class UpdateControlTest_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                testPeriodStart: '',
                testPeriodEnd: '',
                sampleSize: '',
                testType: '',
                effectiveness: '',
                status: ''
        }
        this.updateControlTest_ = this.updateControlTest_.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetestPeriodStartHandler = this.changetestPeriodStartHandler.bind(this);
        this.changetestPeriodEndHandler = this.changetestPeriodEndHandler.bind(this);
        this.changesampleSizeHandler = this.changesampleSizeHandler.bind(this);
        this.changeTestTypeHandler = this.changeTestTypeHandler.bind(this);
        this.changeEffectivenessHandler = this.changeEffectivenessHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ControlTest_Service.getControlTest_ById(this.state.id).then( (res) =>{
            let controlTest_ = res.data;
            this.setState({
                name: controlTest_.name,
                testPeriodStart: controlTest_.testPeriodStart,
                testPeriodEnd: controlTest_.testPeriodEnd,
                sampleSize: controlTest_.sampleSize,
                testType: controlTest_.testType,
                effectiveness: controlTest_.effectiveness,
                status: controlTest_.status
            });
        });
    }

    updateControlTest_ = (e) => {
        e.preventDefault();
        let controlTest_ = {
            controlTest_Id: this.state.id,
            name: this.state.name,
            testPeriodStart: this.state.testPeriodStart,
            testPeriodEnd: this.state.testPeriodEnd,
            sampleSize: this.state.sampleSize,
            testType: this.state.testType,
            effectiveness: this.state.effectiveness,
            status: this.state.status
        };
        console.log('controlTest_ => ' + JSON.stringify(controlTest_));
        console.log('id => ' + JSON.stringify(this.state.id));
        ControlTest_Service.updateControlTest_(controlTest_).then( res => {
            this.props.history.push('/controlTest_s');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetestPeriodStartHandler= (event) => {
        this.setState({testPeriodStart: event.target.value});
    }
    changetestPeriodEndHandler= (event) => {
        this.setState({testPeriodEnd: event.target.value});
    }
    changesampleSizeHandler= (event) => {
        this.setState({sampleSize: event.target.value});
    }
    changeTestTypeHandler= (event) => {
        this.setState({testType: event.target.value});
    }
    changeEffectivenessHandler= (event) => {
        this.setState({effectiveness: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/controlTest_s');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ControlTest_</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> testPeriodStart: </label>
                                                <input type="date" placeholder="testPeriodStart" name="testPeriodStart" className="form-control" value={this.state.testPeriodStart} onChange={this.changetestPeriodStartHandler}/>

                                            <label> testPeriodEnd: </label>
                                                <input type="date" placeholder="testPeriodEnd" name="testPeriodEnd" className="form-control" value={this.state.testPeriodEnd} onChange={this.changetestPeriodEndHandler}/>

                                            <label> sampleSize: </label>
                                                <input type="number" placeholder="sampleSize" name="sampleSize" className="form-control" value={this.state.sampleSize} onChange={this.changesampleSizeHandler}/>

                                            <label> TestType: </label>
                                                <select value={this.state.testType} onChange={this.changeTestTypeHandler}>
                      <option name="TestType" className="form-control" >
                          DesignEffectiveness
                      </option>
                      <option name="TestType" className="form-control" >
                          OperatingEffectiveness
                      </option>
                      <option name="TestType" className="form-control" >
                          Walkthrough
                      </option>
                      <option name="TestType" className="form-control" >
                          Reperformance
                      </option>
                      <option name="TestType" className="form-control" >
                          Inquiry
                      </option>
                      <option name="TestType" className="form-control" >
                          Observation
                      </option>
                      <option name="TestType" className="form-control" >
                          Inspection
                      </option>
                      <option name="TestType" className="form-control" >
                          DataAnalysis
                      </option>
                    </select>

                                            <label> Effectiveness: </label>
                                                <select value={this.state.effectiveness} onChange={this.changeEffectivenessHandler}>
                      <option name="Effectiveness" className="form-control" >
                          Effective
                      </option>
                      <option name="Effectiveness" className="form-control" >
                          PartiallyEffective
                      </option>
                      <option name="Effectiveness" className="form-control" >
                          Ineffective
                      </option>
                      <option name="Effectiveness" className="form-control" >
                          NotTested
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Blocked
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateControlTest_}>Save</button>
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

export default UpdateControlTest_Component
