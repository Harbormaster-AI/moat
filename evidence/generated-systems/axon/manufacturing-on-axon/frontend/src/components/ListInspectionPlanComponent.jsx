import React, { Component } from 'react'
import InspectionPlanService from '../services/InspectionPlanService'

class ListInspectionPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inspectionPlans: []
        }
        this.addInspectionPlan = this.addInspectionPlan.bind(this);
        this.editInspectionPlan = this.editInspectionPlan.bind(this);
        this.deleteInspectionPlan = this.deleteInspectionPlan.bind(this);
    }

    deleteInspectionPlan(id){
        InspectionPlanService.deleteInspectionPlan(id).then( res => {
            this.setState({inspectionPlans: this.state.inspectionPlans.filter(inspectionPlan => inspectionPlan.inspectionPlanId !== id)});
        });
    }
    viewInspectionPlan(id){
        this.props.history.push(`/view-inspectionPlan/${id}`);
    }
    editInspectionPlan(id){
        this.props.history.push(`/add-inspectionPlan/${id}`);
    }

    componentDidMount(){
        InspectionPlanService.getInspectionPlans().then((res) => {
            this.setState({ inspectionPlans: res.data});
        });
    }

    addInspectionPlan(){
        this.props.history.push('/add-inspectionPlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InspectionPlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInspectionPlan}> Add InspectionPlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PlanNumber </th>
                                    <th> Revision </th>
                                    <th> SamplingPlan </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inspectionPlans.map(
                                        inspectionPlan => 
                                        <tr key = {inspectionPlan.inspectionPlanId}>
                                             <td> { inspectionPlan.planNumber } </td>
                                             <td> { inspectionPlan.revision } </td>
                                             <td> { inspectionPlan.samplingPlan } </td>
                                             <td> { inspectionPlan.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInspectionPlan(inspectionPlan.inspectionPlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInspectionPlan(inspectionPlan.inspectionPlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInspectionPlan(inspectionPlan.inspectionPlanId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInspectionPlanComponent
