import React, { Component } from 'react'
import CarePlanService from '../services/CarePlanService';

class CreateCarePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                planNumber: '',
                goalSummary: '',
                status: ''
        }
        this.changeplanNumberHandler = this.changeplanNumberHandler.bind(this);
        this.changegoalSummaryHandler = this.changegoalSummaryHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CarePlanService.getCarePlanById(this.state.id).then( (res) =>{
                let carePlan = res.data;
                this.setState({
                    planNumber: carePlan.planNumber,
                    goalSummary: carePlan.goalSummary,
                    status: carePlan.status
                });
            });
        }        
    }
    saveOrUpdateCarePlan = (e) => {
        e.preventDefault();
        let carePlan = {
                carePlanId: this.state.id,
                planNumber: this.state.planNumber,
                goalSummary: this.state.goalSummary,
                status: this.state.status
            };
        console.log('carePlan => ' + JSON.stringify(carePlan));

        // step 5
        if(this.state.id === '_add'){
            carePlan.carePlanId=''
            CarePlanService.createCarePlan(carePlan).then(res =>{
                this.props.history.push('/carePlans');
            });
        }else{
            CarePlanService.updateCarePlan(carePlan).then( res => {
                this.props.history.push('/carePlans');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CarePlan</h3>
        }else{
            return <h3 className="text-center">Update CarePlan</h3>
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
                                            <label> planNumber:&emsp; </label>
                                                <input placeholder="planNumber" name="planNumber" className="form-control" value={this.state.planNumber} onChange={this.changeplanNumberHandler}/>

                                            <label> goalSummary:&emsp; </label>
                                                <input placeholder="goalSummary" name="goalSummary" className="form-control" value={this.state.goalSummary} onChange={this.changegoalSummaryHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCarePlan}>Save</button>
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

export default CreateCarePlanComponent
