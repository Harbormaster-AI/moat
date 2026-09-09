import React, { Component } from 'react'
import OpportunityLineItemService from '../services/OpportunityLineItemService'

class ListOpportunityLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                opportunityLineItems: []
        }
        this.addOpportunityLineItem = this.addOpportunityLineItem.bind(this);
        this.editOpportunityLineItem = this.editOpportunityLineItem.bind(this);
        this.deleteOpportunityLineItem = this.deleteOpportunityLineItem.bind(this);
    }

    deleteOpportunityLineItem(id){
        OpportunityLineItemService.deleteOpportunityLineItem(id).then( res => {
            this.setState({opportunityLineItems: this.state.opportunityLineItems.filter(opportunityLineItem => opportunityLineItem.opportunityLineItemId !== id)});
        });
    }
    viewOpportunityLineItem(id){
        this.props.history.push(`/view-opportunityLineItem/${id}`);
    }
    editOpportunityLineItem(id){
        this.props.history.push(`/add-opportunityLineItem/${id}`);
    }

    componentDidMount(){
        OpportunityLineItemService.getOpportunityLineItems().then((res) => {
            this.setState({ opportunityLineItems: res.data});
        });
    }

    addOpportunityLineItem(){
        this.props.history.push('/add-opportunityLineItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">OpportunityLineItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOpportunityLineItem}> Add OpportunityLineItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Quantity </th>
                                    <th> UnitPrice </th>
                                    <th> DiscountPercent </th>
                                    <th> TotalPrice </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.opportunityLineItems.map(
                                        opportunityLineItem => 
                                        <tr key = {opportunityLineItem.opportunityLineItemId}>
                                             <td> { opportunityLineItem.quantity } </td>
                                             <td> { opportunityLineItem.unitPrice } </td>
                                             <td> { opportunityLineItem.discountPercent } </td>
                                             <td> { opportunityLineItem.totalPrice } </td>
                                             <td>
                                                 <button onClick={ () => this.editOpportunityLineItem(opportunityLineItem.opportunityLineItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOpportunityLineItem(opportunityLineItem.opportunityLineItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOpportunityLineItem(opportunityLineItem.opportunityLineItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOpportunityLineItemComponent
