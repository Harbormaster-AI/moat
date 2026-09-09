import React, { Component } from 'react'
import PurchaseOrderLineService from '../services/PurchaseOrderLineService';

class CreatePurchaseOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                unitPrice: '',
                dueDate: ''
        }
        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePurchaseOrderLine = (e) => {
        e.preventDefault();
        let purchaseOrderLine = {
                purchaseOrderLineId: this.state.id,
                lineNumber: this.state.lineNumber,
                quantity: this.state.quantity,
                unitPrice: this.state.unitPrice,
                dueDate: this.state.dueDate
            };
        console.log('purchaseOrderLine => ' + JSON.stringify(purchaseOrderLine));

        // step 5
        if(this.state.id === '_add'){
            purchaseOrderLine.purchaseOrderLineId=''
            PurchaseOrderLineService.createPurchaseOrderLine(purchaseOrderLine).then(res =>{
                this.props.history.push('/purchaseOrderLines');
            });
        }else{
            PurchaseOrderLineService.updatePurchaseOrderLine(purchaseOrderLine).then( res => {
                this.props.history.push('/purchaseOrderLines');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PurchaseOrderLine</h3>
        }else{
            return <h3 className="text-center">Update PurchaseOrderLine</h3>
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
                                            <label> lineNumber:&emsp; </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice:&emsp; </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> dueDate:&emsp; </label>
                                                <input type="date" placeholder="dueDate" name="dueDate" className="form-control" value={this.state.dueDate} onChange={this.changedueDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePurchaseOrderLine}>Save</button>
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

export default CreatePurchaseOrderLineComponent
