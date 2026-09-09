import React, { Component } from 'react'
import InspectionPlanService from '../services/InspectionPlanService';

class UpdateInspectionPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                planNumber: '',
                revision: '',
                samplingPlan: '',
                status: ''
        }
        this.updateInspectionPlan = this.updateInspectionPlan.bind(this);

        this.changeplanNumberHandler = this.changeplanNumberHandler.bind(this);
        this.changerevisionHandler = this.changerevisionHandler.bind(this);
        this.changeSamplingPlanHandler = this.changeSamplingPlanHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InspectionPlanService.getInspectionPlanById(this.state.id).then( (res) =>{
            let inspectionPlan = res.data;
            this.setState({
                planNumber: inspectionPlan.planNumber,
                revision: inspectionPlan.revision,
                samplingPlan: inspectionPlan.samplingPlan,
                status: inspectionPlan.status
            });
        });
    }

    updateInspectionPlan = (e) => {
        e.preventDefault();
        let inspectionPlan = {
            inspectionPlanId: this.state.id,
            planNumber: this.state.planNumber,
            revision: this.state.revision,
            samplingPlan: this.state.samplingPlan,
            status: this.state.status
        };
        console.log('inspectionPlan => ' + JSON.stringify(inspectionPlan));
        console.log('id => ' + JSON.stringify(this.state.id));
        InspectionPlanService.updateInspectionPlan(inspectionPlan).then( res => {
            this.props.history.push('/inspectionPlans');
        });
    }

    changeplanNumberHandler= (event) => {
        this.setState({planNumber: event.target.value});
    }
    changerevisionHandler= (event) => {
        this.setState({revision: event.target.value});
    }
    changeSamplingPlanHandler= (event) => {
        this.setState({samplingPlan: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inspectionPlans');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InspectionPlan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> planNumber: </label>
                                                <input placeholder="planNumber" name="planNumber" className="form-control" value={this.state.planNumber} onChange={this.changeplanNumberHandler}/>

                                            <label> revision: </label>
                                                <input placeholder="revision" name="revision" className="form-control" value={this.state.revision} onChange={this.changerevisionHandler}/>

                                            <label> SamplingPlan: </label>
                                                <select value={this.state.samplingPlan} onChange={this.changeSamplingPlanHandler}>
                      <option name="SamplingPlan" className="form-control" >
                          Fixed
                      </option>
                      <option name="SamplingPlan" className="form-control" >
                          Percentage
                      </option>
                      <option name="SamplingPlan" className="form-control" >
                          C0
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInspectionPlan}>Save</button>
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

export default UpdateInspectionPlanComponent
