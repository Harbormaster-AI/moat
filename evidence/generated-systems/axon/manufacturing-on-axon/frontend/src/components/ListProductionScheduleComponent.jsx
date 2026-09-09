import React, { Component } from 'react'
import ProductionScheduleService from '../services/ProductionScheduleService'

class ListProductionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                productionSchedules: []
        }
        this.addProductionSchedule = this.addProductionSchedule.bind(this);
        this.editProductionSchedule = this.editProductionSchedule.bind(this);
        this.deleteProductionSchedule = this.deleteProductionSchedule.bind(this);
    }

    deleteProductionSchedule(id){
        ProductionScheduleService.deleteProductionSchedule(id).then( res => {
            this.setState({productionSchedules: this.state.productionSchedules.filter(productionSchedule => productionSchedule.productionScheduleId !== id)});
        });
    }
    viewProductionSchedule(id){
        this.props.history.push(`/view-productionSchedule/${id}`);
    }
    editProductionSchedule(id){
        this.props.history.push(`/add-productionSchedule/${id}`);
    }

    componentDidMount(){
        ProductionScheduleService.getProductionSchedules().then((res) => {
            this.setState({ productionSchedules: res.data});
        });
    }

    addProductionSchedule(){
        this.props.history.push('/add-productionSchedule/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ProductionSchedule List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addProductionSchedule}> Add ProductionSchedule</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ScheduleNumber </th>
                                    <th> HorizonStart </th>
                                    <th> HorizonEnd </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.productionSchedules.map(
                                        productionSchedule => 
                                        <tr key = {productionSchedule.productionScheduleId}>
                                             <td> { productionSchedule.scheduleNumber } </td>
                                             <td> { productionSchedule.horizonStart } </td>
                                             <td> { productionSchedule.horizonEnd } </td>
                                             <td> { productionSchedule.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editProductionSchedule(productionSchedule.productionScheduleId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteProductionSchedule(productionSchedule.productionScheduleId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewProductionSchedule(productionSchedule.productionScheduleId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListProductionScheduleComponent
