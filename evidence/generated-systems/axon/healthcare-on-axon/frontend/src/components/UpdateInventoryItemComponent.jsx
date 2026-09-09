import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class UpdateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                sku: '',
                name: '',
                quantityOnHand: '',
                quantityReserved: ''
        }
        this.updateInventoryItem = this.updateInventoryItem.bind(this);

        this.changeskuHandler = this.changeskuHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
    }

    componentDidMount(){
        InventoryItemService.getInventoryItemById(this.state.id).then( (res) =>{
            let inventoryItem = res.data;
            this.setState({
                sku: inventoryItem.sku,
                name: inventoryItem.name,
                quantityOnHand: inventoryItem.quantityOnHand,
                quantityReserved: inventoryItem.quantityReserved
            });
        });
    }

    updateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
            inventoryItemId: this.state.id,
            sku: this.state.sku,
            name: this.state.name,
            quantityOnHand: this.state.quantityOnHand,
            quantityReserved: this.state.quantityReserved
        };
        console.log('inventoryItem => ' + JSON.stringify(inventoryItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
            this.props.history.push('/inventoryItems');
        });
    }

    changeskuHandler= (event) => {
        this.setState({sku: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changequantityOnHandHandler= (event) => {
        this.setState({quantityOnHand: event.target.value});
    }
    changequantityReservedHandler= (event) => {
        this.setState({quantityReserved: event.target.value});
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
                                            <label> sku: </label>
                                                <input placeholder="sku" name="sku" className="form-control" value={this.state.sku} onChange={this.changeskuHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> quantityOnHand: </label>
                                                <input type="number" placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved: </label>
                                                <input type="number" placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

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
