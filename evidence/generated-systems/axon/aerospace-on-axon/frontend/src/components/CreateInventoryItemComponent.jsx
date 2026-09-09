import React, { Component } from 'react'
import InventoryItemService from '../services/InventoryItemService';

class CreateInventoryItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantityOnHand: '',
                quantityReserved: '',
                lotNumber: ''
        }
        this.changequantityOnHandHandler = this.changequantityOnHandHandler.bind(this);
        this.changequantityReservedHandler = this.changequantityReservedHandler.bind(this);
        this.changelotNumberHandler = this.changelotNumberHandler.bind(this);
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
                    quantityOnHand: inventoryItem.quantityOnHand,
                    quantityReserved: inventoryItem.quantityReserved,
                    lotNumber: inventoryItem.lotNumber
                });
            });
        }        
    }
    saveOrUpdateInventoryItem = (e) => {
        e.preventDefault();
        let inventoryItem = {
                inventoryItemId: this.state.id,
                quantityOnHand: this.state.quantityOnHand,
                quantityReserved: this.state.quantityReserved,
                lotNumber: this.state.lotNumber
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
    
    changequantityOnHandHandler= (event) => {
        this.setState({quantityOnHand: event.target.value});
    }
    changequantityReservedHandler= (event) => {
        this.setState({quantityReserved: event.target.value});
    }
    changelotNumberHandler= (event) => {
        this.setState({lotNumber: event.target.value});
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
                                            <label> quantityOnHand:&emsp; </label>
                                                <input type="number" placeholder="quantityOnHand" name="quantityOnHand" className="form-control" value={this.state.quantityOnHand} onChange={this.changequantityOnHandHandler}/>

                                            <label> quantityReserved:&emsp; </label>
                                                <input type="number" placeholder="quantityReserved" name="quantityReserved" className="form-control" value={this.state.quantityReserved} onChange={this.changequantityReservedHandler}/>

                                            <label> lotNumber:&emsp; </label>
                                                <input placeholder="lotNumber" name="lotNumber" className="form-control" value={this.state.lotNumber} onChange={this.changelotNumberHandler}/>

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
