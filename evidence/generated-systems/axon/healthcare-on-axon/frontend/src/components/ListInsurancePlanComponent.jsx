import React, { Component } from 'react'
import InsurancePlanService from '../services/InsurancePlanService'

class ListInsurancePlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insurancePlans: []
        }
        this.addInsurancePlan = this.addInsurancePlan.bind(this);
        this.editInsurancePlan = this.editInsurancePlan.bind(this);
        this.deleteInsurancePlan = this.deleteInsurancePlan.bind(this);
    }

    deleteInsurancePlan(id){
        InsurancePlanService.deleteInsurancePlan(id).then( res => {
            this.setState({insurancePlans: this.state.insurancePlans.filter(insurancePlan => insurancePlan.insurancePlanId !== id)});
        });
    }
    viewInsurancePlan(id){
        this.props.history.push(`/view-insurancePlan/${id}`);
    }
    editInsurancePlan(id){
        this.props.history.push(`/add-insurancePlan/${id}`);
    }

    componentDidMount(){
        InsurancePlanService.getInsurancePlans().then((res) => {
            this.setState({ insurancePlans: res.data});
        });
    }

    addInsurancePlan(){
        this.props.history.push('/add-insurancePlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InsurancePlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsurancePlan}> Add InsurancePlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> PlanCode </th>
                                    <th> PlanType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insurancePlans.map(
                                        insurancePlan => 
                                        <tr key = {insurancePlan.insurancePlanId}>
                                             <td> { insurancePlan.name } </td>
                                             <td> { insurancePlan.planCode } </td>
                                             <td> { insurancePlan.planType } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsurancePlan(insurancePlan.insurancePlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsurancePlan(insurancePlan.insurancePlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsurancePlan(insurancePlan.insurancePlanId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsurancePlanComponent
