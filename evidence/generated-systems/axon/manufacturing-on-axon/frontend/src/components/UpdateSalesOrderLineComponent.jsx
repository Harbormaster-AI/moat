import React, { Component } from 'react'
import SalesOrderLineService from '../services/SalesOrderLineService';

class UpdateSalesOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                lineNumber: '',
                quantity: '',
                unitPrice: '',
                dueDate: ''
        }
        this.updateSalesOrderLine = this.updateSalesOrderLine.bind(this);

        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changedueDateHandler = this.changedueDateHandler.bind(this);
    }

    componentDidMount(){
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

    updateSalesOrderLine = (e) => {
        e.preventDefault();
        let salesOrderLine = {
            salesOrderLineId: this.state.id,
            lineNumber: this.state.lineNumber,
            quantity: this.state.quantity,
            unitPrice: this.state.unitPrice,
            dueDate: this.state.dueDate
        };
        console.log('salesOrderLine => ' + JSON.stringify(salesOrderLine));
        console.log('id => ' + JSON.stringify(this.state.id));
        SalesOrderLineService.updateSalesOrderLine(salesOrderLine).then( res => {
            this.props.history.push('/salesOrderLines');
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
        this.props.history.push('/salesOrderLines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SalesOrderLine</h3>
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
                                        <button className="btn btn-success" onClick={this.updateSalesOrderLine}>Save</button>
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

export default UpdateSalesOrderLineComponent
