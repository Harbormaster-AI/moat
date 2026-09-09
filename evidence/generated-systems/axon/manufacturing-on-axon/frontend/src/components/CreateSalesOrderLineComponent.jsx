import React, { Component } from 'react'
import SalesOrderLineService from '../services/SalesOrderLineService';

class CreateSalesOrderLineComponent extends Component {
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
            SalesOrderLineService.getSalesOrderLineById(this.state.id).then( (res) =>{
                let salesOrderLine = res.data;
                this.setState({
                    lineNumber: salesOrderLine.lineNumber,
                    quantity: salesOrderLine.quantity,
                    unitPrice: salesOrderLine.unitPrice,
                    dueDate: salesOrderLine.dueDate
                });
            });
        }        
    }
    saveOrUpdateSalesOrderLine = (e) => {
        e.preventDefault();
        let salesOrderLine = {
                salesOrderLineId: this.state.id,
                lineNumber: this.state.lineNumber,
                quantity: this.state.quantity,
                unitPrice: this.state.unitPrice,
                dueDate: this.state.dueDate
            };
        console.log('salesOrderLine => ' + JSON.stringify(salesOrderLine));

        // step 5
        if(this.state.id === '_add'){
            salesOrderLine.salesOrderLineId=''
            SalesOrderLineService.createSalesOrderLine(salesOrderLine).then(res =>{
                this.props.history.push('/salesOrderLines');
            });
        }else{
            SalesOrderLineService.updateSalesOrderLine(salesOrderLine).then( res => {
                this.props.history.push('/salesOrderLines');
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
        this.props.history.push('/salesOrderLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SalesOrderLine</h3>
        }else{
            return <h3 className="text-center">Update SalesOrderLine</h3>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSalesOrderLine}>Save</button>
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

export default CreateSalesOrderLineComponent
