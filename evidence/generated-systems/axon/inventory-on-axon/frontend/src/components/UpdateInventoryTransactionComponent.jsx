import React, { Component } from 'react'
import InventoryTransactionService from '../services/InventoryTransactionService';

class UpdateInventoryTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                transactionNumber: '',
                quantity: '',
                unitCost: '',
                transactionDate: '',
                reasonCode: '',
                transactionType: '',
                unitOfMeasure: '',
                status: ''
        }
        this.updateInventoryTransaction = this.updateInventoryTransaction.bind(this);

        this.changetransactionNumberHandler = this.changetransactionNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitCostHandler = this.changeunitCostHandler.bind(this);
        this.changetransactionDateHandler = this.changetransactionDateHandler.bind(this);
        this.changereasonCodeHandler = this.changereasonCodeHandler.bind(this);
        this.changeTransactionTypeHandler = this.changeTransactionTypeHandler.bind(this);
        this.changeUnitOfMeasureHandler = this.changeUnitOfMeasureHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        InventoryTransactionService.getInventoryTransactionById(this.state.id).then( (res) =>{
            let inventoryTransaction = res.data;
            this.setState({
                transactionNumber: inventoryTransaction.transactionNumber,
                quantity: inventoryTransaction.quantity,
                unitCost: inventoryTransaction.unitCost,
                transactionDate: inventoryTransaction.transactionDate,
                reasonCode: inventoryTransaction.reasonCode,
                transactionType: inventoryTransaction.transactionType,
                unitOfMeasure: inventoryTransaction.unitOfMeasure,
                status: inventoryTransaction.status
            });
        });
    }

    updateInventoryTransaction = (e) => {
        e.preventDefault();
        let inventoryTransaction = {
            inventoryTransactionId: this.state.id,
            transactionNumber: this.state.transactionNumber,
            quantity: this.state.quantity,
            unitCost: this.state.unitCost,
            transactionDate: this.state.transactionDate,
            reasonCode: this.state.reasonCode,
            transactionType: this.state.transactionType,
            unitOfMeasure: this.state.unitOfMeasure,
            status: this.state.status
        };
        console.log('inventoryTransaction => ' + JSON.stringify(inventoryTransaction));
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryTransactionService.updateInventoryTransaction(inventoryTransaction).then( res => {
            this.props.history.push('/inventoryTransactions');
        });
    }

    changetransactionNumberHandler= (event) => {
        this.setState({transactionNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeunitCostHandler= (event) => {
        this.setState({unitCost: event.target.value});
    }
    changetransactionDateHandler= (event) => {
        this.setState({transactionDate: event.target.value});
    }
    changereasonCodeHandler= (event) => {
        this.setState({reasonCode: event.target.value});
    }
    changeTransactionTypeHandler= (event) => {
        this.setState({transactionType: event.target.value});
    }
    changeUnitOfMeasureHandler= (event) => {
        this.setState({unitOfMeasure: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventoryTransactions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InventoryTransaction</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> transactionNumber: </label>
                                                <input placeholder="transactionNumber" name="transactionNumber" className="form-control" value={this.state.transactionNumber} onChange={this.changetransactionNumberHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitCost: </label>
                                                <input placeholder="unitCost" name="unitCost" className="form-control" value={this.state.unitCost} onChange={this.changeunitCostHandler}/>

                                            <label> transactionDate: </label>
                                                <input type="date" placeholder="transactionDate" name="transactionDate" className="form-control" value={this.state.transactionDate} onChange={this.changetransactionDateHandler}/>

                                            <label> reasonCode: </label>
                                                <input placeholder="reasonCode" name="reasonCode" className="form-control" value={this.state.reasonCode} onChange={this.changereasonCodeHandler}/>

                                            <label> TransactionType: </label>
                                                <select value={this.state.transactionType} onChange={this.changeTransactionTypeHandler}>
                      <option name="TransactionType" className="form-control" >
                          Receipt
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Issue
                      </option>
                      <option name="TransactionType" className="form-control" >
                          AdjustmentIncrease
                      </option>
                      <option name="TransactionType" className="form-control" >
                          AdjustmentDecrease
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Reclassification
                      </option>
                      <option name="TransactionType" className="form-control" >
                          TransferOut
                      </option>
                      <option name="TransactionType" className="form-control" >
                          TransferIn
                      </option>
                      <option name="TransactionType" className="form-control" >
                          CountIncrease
                      </option>
                      <option name="TransactionType" className="form-control" >
                          CountDecrease
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Putaway
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Pick
                      </option>
                    </select>

                                            <label> UnitOfMeasure: </label>
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

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Posted
                      </option>
                      <option name="Status" className="form-control" >
                          Voided
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInventoryTransaction}>Save</button>
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

export default UpdateInventoryTransactionComponent
