import React, { Component } from 'react'
import StockAdjustmentLineService from '../services/StockAdjustmentLineService'

class ListStockAdjustmentLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                stockAdjustmentLines: []
        }
        this.addStockAdjustmentLine = this.addStockAdjustmentLine.bind(this);
        this.editStockAdjustmentLine = this.editStockAdjustmentLine.bind(this);
        this.deleteStockAdjustmentLine = this.deleteStockAdjustmentLine.bind(this);
    }

    deleteStockAdjustmentLine(id){
        StockAdjustmentLineService.deleteStockAdjustmentLine(id).then( res => {
            this.setState({stockAdjustmentLines: this.state.stockAdjustmentLines.filter(stockAdjustmentLine => stockAdjustmentLine.stockAdjustmentLineId !== id)});
        });
    }
    viewStockAdjustmentLine(id){
        this.props.history.push(`/view-stockAdjustmentLine/${id}`);
    }
    editStockAdjustmentLine(id){
        this.props.history.push(`/add-stockAdjustmentLine/${id}`);
    }

    componentDidMount(){
        StockAdjustmentLineService.getStockAdjustmentLines().then((res) => {
            this.setState({ stockAdjustmentLines: res.data});
        });
    }

    addStockAdjustmentLine(){
        this.props.history.push('/add-stockAdjustmentLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">StockAdjustmentLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addStockAdjustmentLine}> Add StockAdjustmentLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> Quantity </th>
                                    <th> UnitOfMeasure </th>
                                    <th> StockStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.stockAdjustmentLines.map(
                                        stockAdjustmentLine => 
                                        <tr key = {stockAdjustmentLine.stockAdjustmentLineId}>
                                             <td> { stockAdjustmentLine.lineNumber } </td>
                                             <td> { stockAdjustmentLine.quantity } </td>
                                             <td> { stockAdjustmentLine.unitOfMeasure } </td>
                                             <td> { stockAdjustmentLine.stockStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editStockAdjustmentLine(stockAdjustmentLine.stockAdjustmentLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteStockAdjustmentLine(stockAdjustmentLine.stockAdjustmentLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewStockAdjustmentLine(stockAdjustmentLine.stockAdjustmentLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListStockAdjustmentLineComponent
