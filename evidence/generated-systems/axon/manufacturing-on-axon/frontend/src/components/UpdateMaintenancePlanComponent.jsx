import React, { Component } from 'react'
import MaintenancePlanService from '../services/MaintenancePlanService';

class UpdateMaintenancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                planNumber: '',
                interval: '',
                lastServiceDate: '',
                strategy: ''
        }
        this.updateMaintenancePlan = this.updateMaintenancePlan.bind(this);

        this.changeplanNumberHandler = this.changeplanNumberHandler.bind(this);
        this.changeintervalHandler = this.changeintervalHandler.bind(this);
        this.changelastServiceDateHandler = this.changelastServiceDateHandler.bind(this);
        this.changeStrategyHandler = this.changeStrategyHandler.bind(this);
    }

    componentDidMount(){
        MaintenancePlanService.getMaintenancePlanById(this.state.id).then( (res) =>{
            let maintenancePlan = res.data;
            this.setState({
                planNumber: maintenancePlan.planNumber,
                interval: maintenancePlan.interval,
                lastServiceDate: maintenancePlan.lastServiceDate,
                strategy: maintenancePlan.strategy
            });
        });
    }

    updateMaintenancePlan = (e) => {
        e.preventDefault();
        let maintenancePlan = {
            maintenancePlanId: this.state.id,
            planNumber: this.state.planNumber,
            interval: this.state.interval,
            lastServiceDate: this.state.lastServiceDate,
            strategy: this.state.strategy
        };
        console.log('maintenancePlan => ' + JSON.stringify(maintenancePlan));
        console.log('id => ' + JSON.stringify(this.state.id));
        MaintenancePlanService.updateMaintenancePlan(maintenancePlan).then( res => {
            this.props.history.push('/maintenancePlans');
        });
    }

    changeplanNumberHandler= (event) => {
        this.setState({planNumber: event.target.value});
    }
    changeintervalHandler= (event) => {
        this.setState({interval: event.target.value});
    }
    changelastServiceDateHandler= (event) => {
        this.setState({lastServiceDate: event.target.value});
    }
    changeStrategyHandler= (event) => {
        this.setState({strategy: event.target.value});
    }

    cancel(){
        this.props.history.push('/maintenancePlans');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MaintenancePlan</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> planNumber: </label>
                                                <input placeholder="planNumber" name="planNumber" className="form-control" value={this.state.planNumber} onChange={this.changeplanNumberHandler}/>

                                            <label> interval: </label>
                                                <input placeholder="interval" name="interval" className="form-control" value={this.state.interval} onChange={this.changeintervalHandler}/>

                                            <label> lastServiceDate: </label>
                                                <input type="date" placeholder="lastServiceDate" name="lastServiceDate" className="form-control" value={this.state.lastServiceDate} onChange={this.changelastServiceDateHandler}/>

                                            <label> Strategy: </label>
                                                <select value={this.state.strategy} onChange={this.changeStrategyHandler}>
                      <option name="Strategy" className="form-control" >
                          TimeBased
                      </option>
                      <option name="Strategy" className="form-control" >
                          UsageBased
                      </option>
                      <option name="Strategy" className="form-control" >
                          ConditionBased
                      </option>
                      <option name="Strategy" className="form-control" >
                          Predictive
                      </option>
                      <option name="Strategy" className="form-control" >
                          Corrective
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMaintenancePlan}>Save</button>
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

export default UpdateMaintenancePlanComponent
