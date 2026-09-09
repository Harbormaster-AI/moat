import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService'

class ViewInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inventoryItem: {}
        }
    }

    componentDidMount(){
        InventoryItemService.getInventoryItemById(this.state.id).then( res => {
            this.setState({inventoryItem: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InventoryItem Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantityOnHand:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryItem.quantityOnHand }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantityReserved:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryItem.quantityReserved }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lotNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inventoryItem.lotNumber }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInventoryItemComponent
