import React, { Component } from 'react'
import ShipmentItemService from '../services/ShipmentItemService';

class UpdateShipmentItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: ''
        }
        this.updateShipmentItem = this.updateShipmentItem.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
    }

    componentDidMount(){
        ShipmentItemService.getShipmentItemById(this.state.id).then( (res) =>{
            let shipmentItem = res.data;
            this.setState({
                quantity: shipmentItem.quantity
            });
        });
    }

    updateShipmentItem = (e) => {
        e.preventDefault();
        let shipmentItem = {
            shipmentItemId: this.state.id,
            quantity: this.state.quantity
        };
        console.log('shipmentItem => ' + JSON.stringify(shipmentItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        ShipmentItemService.updateShipmentItem(shipmentItem).then( res => {
            this.props.history.push('/shipmentItems');
        });
    }

    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }

    cancel(){
        this.props.history.push('/shipmentItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ShipmentItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input type="number" placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateShipmentItem}>Save</button>
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

export default UpdateShipmentItemComponent
