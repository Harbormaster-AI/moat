import React, { Component } from 'react'
import StockAdjustmentService from '../services/StockAdjustmentService'

class ListStockAdjustmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                stockAdjustments: []
        }
        this.addStockAdjustment = this.addStockAdjustment.bind(this);
        this.editStockAdjustment = this.editStockAdjustment.bind(this);
        this.deleteStockAdjustment = this.deleteStockAdjustment.bind(this);
    }

    deleteStockAdjustment(id){
        StockAdjustmentService.deleteStockAdjustment(id).then( res => {
            this.setState({stockAdjustments: this.state.stockAdjustments.filter(stockAdjustment => stockAdjustment.stockAdjustmentId !== id)});
        });
    }
    viewStockAdjustment(id){
        this.props.history.push(`/view-stockAdjustment/${id}`);
    }
    editStockAdjustment(id){
        this.props.history.push(`/add-stockAdjustment/${id}`);
    }

    componentDidMount(){
        StockAdjustmentService.getStockAdjustments().then((res) => {
            this.setState({ stockAdjustments: res.data});
        });
    }

    addStockAdjustment(){
        this.props.history.push('/add-stockAdjustment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">StockAdjustment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addStockAdjustment}> Add StockAdjustment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AdjustmentNumber </th>
                                    <th> Reason </th>
                                    <th> AdjustmentDate </th>
                                    <th> AdjustmentType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.stockAdjustments.map(
                                        stockAdjustment => 
                                        <tr key = {stockAdjustment.stockAdjustmentId}>
                                             <td> { stockAdjustment.adjustmentNumber } </td>
                                             <td> { stockAdjustment.reason } </td>
                                             <td> { stockAdjustment.adjustmentDate } </td>
                                             <td> { stockAdjustment.adjustmentType } </td>
                                             <td> { stockAdjustment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editStockAdjustment(stockAdjustment.stockAdjustmentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteStockAdjustment(stockAdjustment.stockAdjustmentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewStockAdjustment(stockAdjustment.stockAdjustmentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListStockAdjustmentComponent
