import React, { Component } from 'react'
import CarePlanService from '../services/CarePlanService';

class UpdateCarePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                planNumber: '',
                goalSummary: '',
                status: ''
        }
        this.updateCarePlan = this.updateCarePlan.bind(this);

        this.changeplanNumberHandler = this.changeplanNumberHandler.bind(this);
        this.changegoalSummaryHandler = this.changegoalSummaryHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CarePlanService.getCarePlanById(this.state.id).then( (res) =>{
            let carePlan = res.data;
            this.setState({
                planNumber: carePlan.planNumber,
                goalSummary: carePlan.goalSummary,
                status: carePlan.status
            });
        });
    }

    updateCarePlan = (e) => {
        e.preventDefault();
        let carePlan = {
            carePlanId: this.state.id,
            planNumber: this.state.planNumber,
            goalSummary: this.state.goalSummary,
            status: this.state.status
        };
        console.log('carePlan => ' + JSON.stringify(carePlan));
        console.log('id => ' + JSON.stringify(this.state.id));
        CarePlanService.updateCarePlan(carePlan).then( res => {
            this.props.history.push('/carePlans');
        });
    }

    changeplanNumberHandler= (event) => {
        this.setState({planNumber: event.target.value});
    }
    changegoalSummaryHandler= (event) => {
        this.setState({goalSummary: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/carePlans');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CarePlan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> planNumber: </label>
                                                <input placeholder="planNumber" name="planNumber" className="form-control" value={this.state.planNumber} onChange={this.changeplanNumberHandler}/>

                                            <label> goalSummary: </label>
                                                <input placeholder="goalSummary" name="goalSummary" className="form-control" value={this.state.goalSummary} onChange={this.changegoalSummaryHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCarePlan}>Save</button>
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

export default UpdateCarePlanComponent
