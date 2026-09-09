import React, { Component } from 'react'
import CarePlanService from '../services/CarePlanService'

class ListCarePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                carePlans: []
        }
        this.addCarePlan = this.addCarePlan.bind(this);
        this.editCarePlan = this.editCarePlan.bind(this);
        this.deleteCarePlan = this.deleteCarePlan.bind(this);
    }

    deleteCarePlan(id){
        CarePlanService.deleteCarePlan(id).then( res => {
            this.setState({carePlans: this.state.carePlans.filter(carePlan => carePlan.carePlanId !== id)});
        });
    }
    viewCarePlan(id){
        this.props.history.push(`/view-carePlan/${id}`);
    }
    editCarePlan(id){
        this.props.history.push(`/add-carePlan/${id}`);
    }

    componentDidMount(){
        CarePlanService.getCarePlans().then((res) => {
            this.setState({ carePlans: res.data});
        });
    }

    addCarePlan(){
        this.props.history.push('/add-carePlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CarePlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCarePlan}> Add CarePlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PlanNumber </th>
                                    <th> GoalSummary </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.carePlans.map(
                                        carePlan => 
                                        <tr key = {carePlan.carePlanId}>
                                             <td> { carePlan.planNumber } </td>
                                             <td> { carePlan.goalSummary } </td>
                                             <td> { carePlan.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCarePlan(carePlan.carePlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCarePlan(carePlan.carePlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCarePlan(carePlan.carePlanId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCarePlanComponent
