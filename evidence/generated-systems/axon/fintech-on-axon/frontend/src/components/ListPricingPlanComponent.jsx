import React, { Component } from 'react'
import PricingPlanService from '../services/PricingPlanService'

class ListPricingPlanComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                pricingPlans: []
        }
        this.addPricingPlan = this.addPricingPlan.bind(this);
        this.editPricingPlan = this.editPricingPlan.bind(this);
        this.deletePricingPlan = this.deletePricingPlan.bind(this);
    }

    deletePricingPlan(id){
        PricingPlanService.deletePricingPlan(id).then( res => {
            this.setState({pricingPlans: this.state.pricingPlans.filter(pricingPlan => pricingPlan.pricingPlanId !== id)});
        });
    }
    viewPricingPlan(id){
        this.props.history.push(`/view-pricingPlan/${id}`);
    }
    editPricingPlan(id){
        this.props.history.push(`/add-pricingPlan/${id}`);
    }

    componentDidMount(){
        PricingPlanService.getPricingPlans().then((res) => {
            this.setState({ pricingPlans: res.data});
        });
    }

    addPricingPlan(){
        this.props.history.push('/add-pricingPlan/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PricingPlan List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPricingPlan}> Add PricingPlan</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> PlanCode </th>
                                    <th> BaseCurrency </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.pricingPlans.map(
                                        pricingPlan => 
                                        <tr key = {pricingPlan.pricingPlanId}>
                                             <td> { pricingPlan.name } </td>
                                             <td> { pricingPlan.planCode } </td>
                                             <td> { pricingPlan.baseCurrency } </td>
                                             <td> { pricingPlan.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPricingPlan(pricingPlan.pricingPlanId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePricingPlan(pricingPlan.pricingPlanId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPricingPlan(pricingPlan.pricingPlanId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPricingPlanComponent
