import React, { Component } from 'react'
import InventoryTransactionService from '../services/InventoryTransactionService'

class ViewInventoryTransactionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inventoryTransaction: {}
        }
    }

    componentDidMount(){
        InventoryTransactionService.getInventoryTransactionById(this.state.id).then( res => {
            this.setState({inventoryTransaction: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InventoryTransaction Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> transactionNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryTransaction.transactionNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryTransaction.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> transactionDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryTransaction.transactionDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> referenceDocument:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryTransaction.referenceDocument }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TransactionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryTransaction.transactionType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInventoryTransactionComponent
