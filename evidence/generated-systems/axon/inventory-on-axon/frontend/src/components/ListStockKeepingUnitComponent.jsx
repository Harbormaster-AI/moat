import React, { Component } from 'react'
import StockKeepingUnitService from '../services/StockKeepingUnitService'

class ListStockKeepingUnitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                stockKeepingUnits: []
        }
        this.addStockKeepingUnit = this.addStockKeepingUnit.bind(this);
        this.editStockKeepingUnit = this.editStockKeepingUnit.bind(this);
        this.deleteStockKeepingUnit = this.deleteStockKeepingUnit.bind(this);
    }

    deleteStockKeepingUnit(id){
        StockKeepingUnitService.deleteStockKeepingUnit(id).then( res => {
            this.setState({stockKeepingUnits: this.state.stockKeepingUnits.filter(stockKeepingUnit => stockKeepingUnit.stockKeepingUnitId !== id)});
        });
    }
    viewStockKeepingUnit(id){
        this.props.history.push(`/view-stockKeepingUnit/${id}`);
    }
    editStockKeepingUnit(id){
        this.props.history.push(`/add-stockKeepingUnit/${id}`);
    }

    componentDidMount(){
        StockKeepingUnitService.getStockKeepingUnits().then((res) => {
            this.setState({ stockKeepingUnits: res.data});
        });
    }

    addStockKeepingUnit(){
        this.props.history.push('/add-stockKeepingUnit/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">StockKeepingUnit List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addStockKeepingUnit}> Add StockKeepingUnit</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> SkuCode </th>
                                    <th> Name </th>
                                    <th> Weight </th>
                                    <th> WeightUnit </th>
                                    <th> Volume </th>
                                    <th> VolumeUnit </th>
                                    <th> ShelfLifeDays </th>
                                    <th> HazardousMaterial </th>
                                    <th> ItemType </th>
                                    <th> UnitOfMeasure </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.stockKeepingUnits.map(
                                        stockKeepingUnit => 
                                        <tr key = {stockKeepingUnit.stockKeepingUnitId}>
                                             <td> { stockKeepingUnit.skuCode } </td>
                                             <td> { stockKeepingUnit.name } </td>
                                             <td> { stockKeepingUnit.weight } </td>
                                             <td> { stockKeepingUnit.weightUnit } </td>
                                             <td> { stockKeepingUnit.volume } </td>
                                             <td> { stockKeepingUnit.volumeUnit } </td>
                                             <td> { stockKeepingUnit.shelfLifeDays } </td>
                                             <td> { stockKeepingUnit.hazardousMaterial } </td>
                                             <td> { stockKeepingUnit.itemType } </td>
                                             <td> { stockKeepingUnit.unitOfMeasure } </td>
                                             <td>
                                                 <button onClick={ () => this.editStockKeepingUnit(stockKeepingUnit.stockKeepingUnitId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteStockKeepingUnit(stockKeepingUnit.stockKeepingUnitId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewStockKeepingUnit(stockKeepingUnit.stockKeepingUnitId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListStockKeepingUnitComponent
