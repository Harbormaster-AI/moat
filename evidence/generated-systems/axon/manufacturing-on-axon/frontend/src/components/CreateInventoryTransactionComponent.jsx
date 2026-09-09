import React, { Component } from 'react'
import InventoryTransactionService from '../services/InventoryTransactionService';

class CreateInventoryTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                transactionNumber: '',
                quantity: '',
                transactionDateTime: '',
                referenceDocument: '',
                transactionType: ''
        }
        this.changetransactionNumberHandler = this.changetransactionNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changetransactionDateTimeHandler = this.changetransactionDateTimeHandler.bind(this);
        this.changereferenceDocumentHandler = this.changereferenceDocumentHandler.bind(this);
        this.changeTransactionTypeHandler = this.changeTransactionTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InventoryTransactionService.getInventoryTransactionById(this.state.id).then( (res) =>{
                let inventoryTransaction = res.data;
                this.setState({
                    transactionNumber: inventoryTransaction.transactionNumber,
                    quantity: inventoryTransaction.quantity,
                    transactionDateTime: inventoryTransaction.transactionDateTime,
                    referenceDocument: inventoryTransaction.referenceDocument,
                    transactionType: inventoryTransaction.transactionType
                });
            });
        }        
    }
    saveOrUpdateInventoryTransaction = (e) => {
        e.preventDefault();
        let inventoryTransaction = {
                inventoryTransactionId: this.state.id,
                transactionNumber: this.state.transactionNumber,
                quantity: this.state.quantity,
                transactionDateTime: this.state.transactionDateTime,
                referenceDocument: this.state.referenceDocument,
                transactionType: this.state.transactionType
            };
        console.log('inventoryTransaction => ' + JSON.stringify(inventoryTransaction));

        // step 5
        if(this.state.id === '_add'){
            inventoryTransaction.inventoryTransactionId=''
            InventoryTransactionService.createInventoryTransaction(inventoryTransaction).then(res =>{
                this.props.history.push('/inventoryTransactions');
            });
        }else{
            InventoryTransactionService.updateInventoryTransaction(inventoryTransaction).then( res => {
                this.props.history.push('/inventoryTransactions');
            });
        }
    }
    
    changetransactionNumberHandler= (event) => {
        this.setState({transactionNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changetransactionDateTimeHandler= (event) => {
        this.setState({transactionDateTime: event.target.value});
    }
    changereferenceDocumentHandler= (event) => {
        this.setState({referenceDocument: event.target.value});
    }
    changeTransactionTypeHandler= (event) => {
        this.setState({transactionType: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventoryTransactions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InventoryTransaction</h3>
        }else{
            return <h3 className="text-center">Update InventoryTransaction</h3>
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
                                            <label> transactionNumber:&emsp; </label>
                                                <input placeholder="transactionNumber" name="transactionNumber" className="form-control" value={this.state.transactionNumber} onChange={this.changetransactionNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> transactionDateTime:&emsp; </label>
                                                <input type="time" placeholder="transactionDateTime" name="transactionDateTime" className="form-control" value={this.state.transactionDateTime} onChange={this.changetransactionDateTimeHandler}/>

                                            <label> referenceDocument:&emsp; </label>
                                                <input placeholder="referenceDocument" name="referenceDocument" className="form-control" value={this.state.referenceDocument} onChange={this.changereferenceDocumentHandler}/>

                                            <label> TransactionType:&emsp; </label>
                                                <select value={this.state.transactionType} onChange={this.changeTransactionTypeHandler}>
                      <option name="TransactionType" className="form-control" >
                          Receipt
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Issue
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Return
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Adjustment
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Transfer
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Consumption
                      </option>
                      <option name="TransactionType" className="form-control" >
                          ProductionReceipt
                      </option>
                      <option name="TransactionType" className="form-control" >
                          Scrap
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInventoryTransaction}>Save</button>
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

export default CreateInventoryTransactionComponent
