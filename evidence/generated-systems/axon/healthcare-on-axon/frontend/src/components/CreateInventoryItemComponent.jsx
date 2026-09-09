import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class CreateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                sku: '',
                name: '',
                quantityOnHand: '',
                quantityReserved: ''
        }
        this.changeskuHandler = this.changeskuHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
                inventoryItemId: this.state.id,
                sku: this.state.sku,
                name: this.state.name,
                quantityOnHand: this.state.quantityOnHand,
                quantityReserved: this.state.quantityReserved
            };
        console.log('inventoryItem => ' + JSON.stringify(inventoryItem));

        // step 5
        if(this.state.id === '_add'){
            inventoryItem.inventoryItemId=''
            InventoryItemService.createInventoryItem(inventoryItem).then(res =>{
                this.props.history.push('/inventoryItems');
            });
        }else{
            InventoryItemService.updateInventoryItem(inventoryItem).then( res => {
                this.props.history.push('/inventoryItems');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InventoryItem</h3>
        }else{
            return <h3 className="text-center">Update InventoryItem</h3>
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
                                            <label> sku:&emsp; </label>
                                                <input placeholder="sku" name="sku" className="form-control" value={this.state.sku} onChange={this.changeskuHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> quantityOnHand:&emsp; </label>
                                                <input type="number" placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved:&emsp; </label>
                                                <input type="number" placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInventoryItem}>Save</button>
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

export default CreateInventoryItemComponent
