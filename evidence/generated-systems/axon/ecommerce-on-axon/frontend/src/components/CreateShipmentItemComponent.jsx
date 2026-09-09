import React, { Component } from 'react'
import ShipmentItemService from '../services/ShipmentItemService';

class CreateShipmentItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quantity: ''
        }
        this.changequantityHandler = this.changequantityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ShipmentItemService.getShipmentItemById(this.state.id).then( (res) =>{
                let shipmentItem = res.data;
                this.setState({
                    quantity: shipmentItem.quantity
                });
            });
        }        
    }
    saveOrUpdateShipmentItem = (e) => {
        e.preventDefault();
        let shipmentItem = {
                shipmentItemId: this.state.id,
                quantity: this.state.quantity
            };
        console.log('shipmentItem => ' + JSON.stringify(shipmentItem));

        // step 5
        if(this.state.id === '_add'){
            shipmentItem.shipmentItemId=''
            ShipmentItemService.createShipmentItem(shipmentItem).then(res =>{
                this.props.history.push('/shipmentItems');
            });
        }else{
            ShipmentItemService.updateShipmentItem(shipmentItem).then( res => {
                this.props.history.push('/shipmentItems');
            });
        }
    }
    
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }

    cancel(){
        this.props.history.push('/shipmentItems');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ShipmentItem</h3>
        }else{
            return <h3 className="text-center">Update ShipmentItem</h3>
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
                                            <label> quantity:&emsp; </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateShipmentItem}>Save</button>
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

export default CreateShipmentItemComponent
