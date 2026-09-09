import React, { Component } from 'react'
import StockAdjustmentLineService from '../services/StockAdjustmentLineService';

class CreateStockAdjustmentLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                unitOfMeasure: '',
                stockStatus: ''
        }
        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeUnitOfMeasureHandler = this.changeUnitOfMeasureHandler.bind(this);
        this.changeStockStatusHandler = this.changeStockStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            StockAdjustmentLineService.getStockAdjustmentLineById(this.state.id).then( (res) =>{
                let stockAdjustmentLine = res.data;
                this.setState({
                    lineNumber: stockAdjustmentLine.lineNumber,
                    quantity: stockAdjustmentLine.quantity,
                    unitOfMeasure: stockAdjustmentLine.unitOfMeasure,
                    stockStatus: stockAdjustmentLine.stockStatus
                });
            });
        }        
    }
    saveOrUpdateStockAdjustmentLine = (e) => {
        e.preventDefault();
        let stockAdjustmentLine = {
                stockAdjustmentLineId: this.state.id,
                lineNumber: this.state.lineNumber,
                quantity: this.state.quantity,
                unitOfMeasure: this.state.unitOfMeasure,
                stockStatus: this.state.stockStatus
            };
        console.log('stockAdjustmentLine => ' + JSON.stringify(stockAdjustmentLine));

        // step 5
        if(this.state.id === '_add'){
            stockAdjustmentLine.stockAdjustmentLineId=''
            StockAdjustmentLineService.createStockAdjustmentLine(stockAdjustmentLine).then(res =>{
                this.props.history.push('/stockAdjustmentLines');
            });
        }else{
            StockAdjustmentLineService.updateStockAdjustmentLine(stockAdjustmentLine).then( res => {
                this.props.history.push('/stockAdjustmentLines');
            });
        }
    }
    
    changelineNumberHandler= (event) => {
        this.setState({lineNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeUnitOfMeasureHandler= (event) => {
        this.setState({unitOfMeasure: event.target.value});
    }
    changeStockStatusHandler= (event) => {
        this.setState({stockStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/stockAdjustmentLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add StockAdjustmentLine</h3>
        }else{
            return <h3 className="text-center">Update StockAdjustmentLine</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber:&emsp; </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> UnitOfMeasure:&emsp; </label>
                                                <select value={this.state.unitOfMeasure} onChange={this.changeUnitOfMeasureHandler}>
                      <option name="UnitOfMeasure" className="form-control" >
                          Each
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Case
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pallet
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Dozen
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Gram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Kilogram
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Pound
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Ounce
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Milliliter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Liter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          CubicMeter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Meter
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          Foot
                      </option>
                      <option name="UnitOfMeasure" className="form-control" >
                          SquareMeter
                      </option>
                    </select>

                                            <label> StockStatus:&emsp; </label>
                                                <select value={this.state.stockStatus} onChange={this.changeStockStatusHandler}>
                      <option name="StockStatus" className="form-control" >
                          Available
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Reserved
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Damaged
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Hold
                      </option>
                      <option name="StockStatus" className="form-control" >
                          Quarantined
                      </option>
                      <option name="StockStatus" className="form-control" >
                          InTransit
                      </option>
                      <option name="StockStatus" className="form-control" >
                          PendingInspection
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateStockAdjustmentLine}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateStockAdjustmentLineComponent
