import React, { Component } from 'react'
import MaintenancePlanService from '../services/MaintenancePlanService'

class ListMaintenancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                maintenancePlans: []
        }
        this.addMaintenancePlan = this.addMaintenancePlan.bind(this);
        this.editMaintenancePlan = this.editMaintenancePlan.bind(this);
        this.deleteMaintenancePlan = this.deleteMaintenancePlan.bind(this);
    }

    deleteMaintenancePlan(id){
        MaintenancePlanService.deleteMaintenancePlan(id).then( res => {
            this.setState({maintenancePlans: this.state.maintenancePlans.filter(maintenancePlan => maintenancePlan.maintenancePlanId !== id)});
        });
    }
    viewMaintenancePlan(id){
        this.props.history.push(`/view-maintenancePlan/${id}`);
    }
    editMaintenancePlan(id){
        this.props.history.push(`/add-maintenancePlan/${id}`);
    }

    componentDidMount(){
        MaintenancePlanService.getMaintenancePlans().then((res) => {
            this.setState({ maintenancePlans: res.data});
        });
    }

    addMaintenancePlan(){
        this.props.history.push('/add-maintenancePlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MaintenancePlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMaintenancePlan}> Add MaintenancePlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PlanNumber </th>
                                    <th> Interval </th>
                                    <th> LastServiceDate </th>
                                    <th> Strategy </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.maintenancePlans.map(
                                        maintenancePlan => 
                                        <tr key = {maintenancePlan.maintenancePlanId}>
                                             <td> { maintenancePlan.planNumber } </td>
                                             <td> { maintenancePlan.interval } </td>
                                             <td> { maintenancePlan.lastServiceDate } </td>
                                             <td> { maintenancePlan.strategy } </td>
                                             <td>
                                                 <button onClick={ () => this.editMaintenancePlan(maintenancePlan.maintenancePlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMaintenancePlan(maintenancePlan.maintenancePlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMaintenancePlan(maintenancePlan.maintenancePlanId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListMaintenancePlanComponent
