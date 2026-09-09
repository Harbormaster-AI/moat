import React, { Component } from 'react'
import MRPRunService from '../services/MRPRunService';

class UpdateMRPRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                runNumber: '',
                runDateTime: '',
                planningHorizonDays: '',
                status: ''
        }
        this.updateMRPRun = this.updateMRPRun.bind(this);

        this.changerunNumberHandler = this.changerunNumberHandler.bind(this);
        this.changerunDateTimeHandler = this.changerunDateTimeHandler.bind(this);
        this.changeplanningHorizonDaysHandler = this.changeplanningHorizonDaysHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        MRPRunService.getMRPRunById(this.state.id).then( (res) =>{
            let mRPRun = res.data;
            this.setState({
                runNumber: mRPRun.runNumber,
                runDateTime: mRPRun.runDateTime,
                planningHorizonDays: mRPRun.planningHorizonDays,
                status: mRPRun.status
            });
        });
    }

    updateMRPRun = (e) => {
        e.preventDefault();
        let mRPRun = {
            mRPRunId: this.state.id,
            runNumber: this.state.runNumber,
            runDateTime: this.state.runDateTime,
            planningHorizonDays: this.state.planningHorizonDays,
            status: this.state.status
        };
        console.log('mRPRun => ' + JSON.stringify(mRPRun));
        console.log('id => ' + JSON.stringify(this.state.id));
        MRPRunService.updateMRPRun(mRPRun).then( res => {
            this.props.history.push('/mRPRuns');
        });
    }

    changerunNumberHandler= (event) => {
        this.setState({runNumber: event.target.value});
    }
    changerunDateTimeHandler= (event) => {
        this.setState({runDateTime: event.target.value});
    }
    changeplanningHorizonDaysHandler= (event) => {
        this.setState({planningHorizonDays: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/mRPRuns');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MRPRun</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> runNumber: </label>
                                                <input placeholder="runNumber" name="runNumber" className="form-control" value={this.state.runNumber} onChange={this.changerunNumberHandler}/>

                                            <label> runDateTime: </label>
                                                <input type="time" placeholder="runDateTime" name="runDateTime" className="form-control" value={this.state.runDateTime} onChange={this.changerunDateTimeHandler}/>

                                            <label> planningHorizonDays: </label>
                                                <input type="number" placeholder="planningHorizonDays" name="planningHorizonDays" className="form-control" value={this.state.planningHorizonDays} onChange={this.changeplanningHorizonDaysHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Started
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMRPRun}>Save</button>
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

export default UpdateMRPRunComponent
