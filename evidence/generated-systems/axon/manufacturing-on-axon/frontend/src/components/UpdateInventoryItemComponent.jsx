import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class UpdateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityReserved: '',
                lotNumber: '',
                serialNumber: ''
        }
        this.updateInventoryItem = this.updateInventoryItem.bind(this);

        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changelotNumberHandler = this.changelotNumberHandler.bind(this);
        this.changeserialNumberHandler = this.changeserialNumberHandler.bind(this);
    }

    componentDidMount(){
        InventoryItemService.getInventoryItemById(this.state.id).then( (res) =>{
            let inventoryItem = res.data;
            this.setState({
                quantityOnHand: inventoryItem.quantityOnHand,
                quantityReserved: inventoryItem.quantityReserved,
                lotNumber: inventoryItem.lotNumber,
                serialNumber: inventoryItem.serialNumber
            });
        });
    }

    updateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
            inventoryItemId: this.state.id,
            quantityOnHand: this.state.quantityOnHand,
            quantityReserved: this.state.quantityReserved,
            lotNumber: this.state.lotNumber,
            serialNumber: this.state.serialNumber
        };
        console.log('inventoryItem => ' + JSON.stringify(inventoryItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
            this.props.history.push('/inventoryItems');
        });
    }

    changequantityOnHandHandler= (event) => {
        this.setState({quantityOnHand: event.target.value});
    }
    changequantityReservedHandler= (event) => {
        this.setState({quantityReserved: event.target.value});
    }
    changelotNumberHandler= (event) => {
        this.setState({lotNumber: event.target.value});
    }
    changeserialNumberHandler= (event) => {
        this.setState({serialNumber: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventoryItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InventoryItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantityOnHand: </label>
                                                <input placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved: </label>
                                                <input placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> lotNumber: </label>
                                                <input placeholder="lotNumber" name="lotNumber" className="form-control" value={this.state.lotNumber} onChange={this.changelotNumberHandler}/>

                                            <label> serialNumber: </label>
                                                <input placeholder="serialNumber" name="serialNumber" className="form-control" value={this.state.serialNumber} onChange={this.changeserialNumberHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInventoryItem}>Save</button>
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

export default UpdateInventoryItemComponent
