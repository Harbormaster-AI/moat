import React, { Component } from 'react'
import CostCenterService from '../services/CostCenterService'

class ListCostCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                costCenters: []
        }
        this.addCostCenter = this.addCostCenter.bind(this);
        this.editCostCenter = this.editCostCenter.bind(this);
        this.deleteCostCenter = this.deleteCostCenter.bind(this);
    }

    deleteCostCenter(id){
        CostCenterService.deleteCostCenter(id).then( res => {
            this.setState({costCenters: this.state.costCenters.filter(costCenter => costCenter.costCenterId !== id)});
        });
    }
    viewCostCenter(id){
        this.props.history.push(`/view-costCenter/${id}`);
    }
    editCostCenter(id){
        this.props.history.push(`/add-costCenter/${id}`);
    }

    componentDidMount(){
        CostCenterService.getCostCenters().then((res) => {
            this.setState({ costCenters: res.data});
        });
    }

    addCostCenter(){
        this.props.history.push('/add-costCenter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CostCenter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCostCenter}> Add CostCenter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.costCenters.map(
                                        costCenter => 
                                        <tr key = {costCenter.costCenterId}>
                                             <td> { costCenter.code } </td>
                                             <td> { costCenter.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editCostCenter(costCenter.costCenterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCostCenter(costCenter.costCenterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCostCenter(costCenter.costCenterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCostCenterComponent
