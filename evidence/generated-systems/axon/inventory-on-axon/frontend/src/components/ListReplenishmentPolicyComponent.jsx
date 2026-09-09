import React, { Component } from 'react'
import ReplenishmentPolicyService from '../services/ReplenishmentPolicyService'

class ListReplenishmentPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                replenishmentPolicys: []
        }
        this.addReplenishmentPolicy = this.addReplenishmentPolicy.bind(this);
        this.editReplenishmentPolicy = this.editReplenishmentPolicy.bind(this);
        this.deleteReplenishmentPolicy = this.deleteReplenishmentPolicy.bind(this);
    }

    deleteReplenishmentPolicy(id){
        ReplenishmentPolicyService.deleteReplenishmentPolicy(id).then( res => {
            this.setState({replenishmentPolicys: this.state.replenishmentPolicys.filter(replenishmentPolicy => replenishmentPolicy.replenishmentPolicyId !== id)});
        });
    }
    viewReplenishmentPolicy(id){
        this.props.history.push(`/view-replenishmentPolicy/${id}`);
    }
    editReplenishmentPolicy(id){
        this.props.history.push(`/add-replenishmentPolicy/${id}`);
    }

    componentDidMount(){
        ReplenishmentPolicyService.getReplenishmentPolicys().then((res) => {
            this.setState({ replenishmentPolicys: res.data});
        });
    }

    addReplenishmentPolicy(){
        this.props.history.push('/add-replenishmentPolicy/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ReplenishmentPolicy List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReplenishmentPolicy}> Add ReplenishmentPolicy</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> MinLevel </th>
                                    <th> MaxLevel </th>
                                    <th> ReorderPoint </th>
                                    <th> ReorderQuantity </th>
                                    <th> LeadTimeDays </th>
                                    <th> ReviewPeriodDays </th>
                                    <th> PolicyType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.replenishmentPolicys.map(
                                        replenishmentPolicy => 
                                        <tr key = {replenishmentPolicy.replenishmentPolicyId}>
                                             <td> { replenishmentPolicy.minLevel } </td>
                                             <td> { replenishmentPolicy.maxLevel } </td>
                                             <td> { replenishmentPolicy.reorderPoint } </td>
                                             <td> { replenishmentPolicy.reorderQuantity } </td>
                                             <td> { replenishmentPolicy.leadTimeDays } </td>
                                             <td> { replenishmentPolicy.reviewPeriodDays } </td>
                                             <td> { replenishmentPolicy.policyType } </td>
                                             <td>
                                                 <button onClick={ () => this.editReplenishmentPolicy(replenishmentPolicy.replenishmentPolicyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReplenishmentPolicy(replenishmentPolicy.replenishmentPolicyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReplenishmentPolicy(replenishmentPolicy.replenishmentPolicyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListReplenishmentPolicyComponent
