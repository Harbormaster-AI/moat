import React, { Component } from 'react'
import TransferOrderLineService from '../services/TransferOrderLineService';

class CreateTransferOrderLineComponent extends Component {
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
            TransferOrderLineService.getTransferOrderLineById(this.state.id).then( (res) =>{
                let transferOrderLine = res.data;
                this.setState({
                    lineNumber: transferOrderLine.lineNumber,
                    quantity: transferOrderLine.quantity,
                    unitOfMeasure: transferOrderLine.unitOfMeasure,
                    stockStatus: transferOrderLine.stockStatus
                });
            });
        }        
    }
    saveOrUpdateTransferOrderLine = (e) => {
        e.preventDefault();
        let transferOrderLine = {
                transferOrderLineId: this.state.id,
                lineNumber: this.state.lineNumber,
                quantity: this.state.quantity,
                unitOfMeasure: this.state.unitOfMeasure,
                stockStatus: this.state.stockStatus
            };
        console.log('transferOrderLine => ' + JSON.stringify(transferOrderLine));

        // step 5
        if(this.state.id === '_add'){
            transferOrderLine.transferOrderLineId=''
            TransferOrderLineService.createTransferOrderLine(transferOrderLine).then(res =>{
                this.props.history.push('/transferOrderLines');
            });
        }else{
            TransferOrderLineService.updateTransferOrderLine(transferOrderLine).then( res => {
                this.props.history.push('/transferOrderLines');
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
        this.props.history.push('/transferOrderLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TransferOrderLine</h3>
        }else{
            return <h3 className="text-center">Update TransferOrderLine</h3>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTransferOrderLine}>Save</button>
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

export default CreateTransferOrderLineComponent
