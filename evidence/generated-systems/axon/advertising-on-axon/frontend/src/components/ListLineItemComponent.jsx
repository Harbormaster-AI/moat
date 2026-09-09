import React, { Component } from 'react'
import LineItemService from '../services/LineItemService'

class ListLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                lineItems: []
        }
        this.addLineItem = this.addLineItem.bind(this);
        this.editLineItem = this.editLineItem.bind(this);
        this.deleteLineItem = this.deleteLineItem.bind(this);
    }

    deleteLineItem(id){
        LineItemService.deleteLineItem(id).then( res => {
            this.setState({lineItems: this.state.lineItems.filter(lineItem => lineItem.lineItemId !== id)});
        });
    }
    viewLineItem(id){
        this.props.history.push(`/view-lineItem/${id}`);
    }
    editLineItem(id){
        this.props.history.push(`/add-lineItem/${id}`);
    }

    componentDidMount(){
        LineItemService.getLineItems().then((res) => {
            this.setState({ lineItems: res.data});
        });
    }

    addLineItem(){
        this.props.history.push('/add-lineItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LineItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLineItem}> Add LineItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> BidAmount </th>
                                    <th> DailyBudget </th>
                                    <th> FrequencyCap </th>
                                    <th> Status </th>
                                    <th> PricingModel </th>
                                    <th> BidStrategy </th>
                                    <th> Pacing </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.lineItems.map(
                                        lineItem => 
                                        <tr key = {lineItem.lineItemId}>
                                             <td> { lineItem.name } </td>
                                             <td> { lineItem.bidAmount } </td>
                                             <td> { lineItem.dailyBudget } </td>
                                             <td> { lineItem.frequencyCap } </td>
                                             <td> { lineItem.status } </td>
                                             <td> { lineItem.pricingModel } </td>
                                             <td> { lineItem.bidStrategy } </td>
                                             <td> { lineItem.pacing } </td>
                                             <td>
                                                 <button onClick={ () => this.editLineItem(lineItem.lineItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLineItem(lineItem.lineItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLineItem(lineItem.lineItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLineItemComponent
