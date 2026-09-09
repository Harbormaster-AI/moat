import React, { Component } from 'react'
import PurchaseOrderLineService from '../services/PurchaseOrderLineService';

class UpdatePurchaseOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                unitPrice: '',
                dueDate: ''
        }
        this.updatePurchaseOrderLine = this.updatePurchaseOrderLine.bind(this);

        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
    }

    componentDidMount(){
        PurchaseOrderLineService.getPurchaseOrderLineById(this.state.id).then( (res) =>{
            let purchaseOrderLine = res.data;
            this.setState({
                lineNumber: purchaseOrderLine.lineNumber,
                quantity: purchaseOrderLine.quantity,
                unitPrice: purchaseOrderLine.unitPrice,
                dueDate: purchaseOrderLine.dueDate
            });
        });
    }

    updatePurchaseOrderLine = (e) => {
        e.preventDefault();
        let purchaseOrderLine = {
            purchaseOrderLineId: this.state.id,
            lineNumber: this.state.lineNumber,
            quantity: this.state.quantity,
            unitPrice: this.state.unitPrice,
            dueDate: this.state.dueDate
        };
        console.log('purchaseOrderLine => ' + JSON.stringify(purchaseOrderLine));
        console.log('id => ' + JSON.stringify(this.state.id));
        PurchaseOrderLineService.updatePurchaseOrderLine(purchaseOrderLine).then( res => {
            this.props.history.push('/purchaseOrderLines');
        });
    }

    changelineNumberHandler= (event) => {
        this.setState({lineNumber: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changedueDateHandler= (event) => {
        this.setState({dueDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/purchaseOrderLines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PurchaseOrderLine</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber: </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> dueDate: </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePurchaseOrderLine}>Save</button>
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

export default UpdatePurchaseOrderLineComponent
