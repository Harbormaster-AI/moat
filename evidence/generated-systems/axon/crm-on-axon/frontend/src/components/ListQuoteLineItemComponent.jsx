import React, { Component } from 'react'
import QuoteLineItemService from '../services/QuoteLineItemService'

class ListQuoteLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                quoteLineItems: []
        }
        this.addQuoteLineItem = this.addQuoteLineItem.bind(this);
        this.editQuoteLineItem = this.editQuoteLineItem.bind(this);
        this.deleteQuoteLineItem = this.deleteQuoteLineItem.bind(this);
    }

    deleteQuoteLineItem(id){
        QuoteLineItemService.deleteQuoteLineItem(id).then( res => {
            this.setState({quoteLineItems: this.state.quoteLineItems.filter(quoteLineItem => quoteLineItem.quoteLineItemId !== id)});
        });
    }
    viewQuoteLineItem(id){
        this.props.history.push(`/view-quoteLineItem/${id}`);
    }
    editQuoteLineItem(id){
        this.props.history.push(`/add-quoteLineItem/${id}`);
    }

    componentDidMount(){
        QuoteLineItemService.getQuoteLineItems().then((res) => {
            this.setState({ quoteLineItems: res.data});
        });
    }

    addQuoteLineItem(){
        this.props.history.push('/add-quoteLineItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">QuoteLineItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addQuoteLineItem}> Add QuoteLineItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> DiscountAmount </th>
                                    <th> TaxAmount </th>
                                    <th> TotalAmount </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.quoteLineItems.map(
                                        quoteLineItem => 
                                        <tr key = {quoteLineItem.quoteLineItemId}>
                                             <td> { quoteLineItem.quantity } </td>
                                             <td> { quoteLineItem.unitPrice } </td>
                                             <td> { quoteLineItem.discountAmount } </td>
                                             <td> { quoteLineItem.taxAmount } </td>
                                             <td> { quoteLineItem.totalAmount } </td>
                                             <td>
                                                 <button onClick={ () => this.editQuoteLineItem(quoteLineItem.quoteLineItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteQuoteLineItem(quoteLineItem.quoteLineItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewQuoteLineItem(quoteLineItem.quoteLineItemId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListQuoteLineItemComponent
