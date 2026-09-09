import React, { Component } from 'react'
import OpportunityLineItemService from '../services/OpportunityLineItemService';

class UpdateOpportunityLineItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                quantity: '',
                unitPrice: '',
                discountPercent: '',
                totalPrice: ''
        }
        this.updateOpportunityLineItem = this.updateOpportunityLineItem.bind(this);

        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changediscountPercentHandler = this.changediscountPercentHandler.bind(this);
        this.changetotalPriceHandler = this.changetotalPriceHandler.bind(this);
    }

    componentDidMount(){
        OpportunityLineItemService.getOpportunityLineItemById(this.state.id).then( (res) =>{
            let opportunityLineItem = res.data;
            this.setState({
                quantity: opportunityLineItem.quantity,
                unitPrice: opportunityLineItem.unitPrice,
                discountPercent: opportunityLineItem.discountPercent,
                totalPrice: opportunityLineItem.totalPrice
            });
        });
    }

    updateOpportunityLineItem = (e) => {
        e.preventDefault();
        let opportunityLineItem = {
            opportunityLineItemId: this.state.id,
            quantity: this.state.quantity,
            unitPrice: this.state.unitPrice,
            discountPercent: this.state.discountPercent,
            totalPrice: this.state.totalPrice
        };
        console.log('opportunityLineItem => ' + JSON.stringify(opportunityLineItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        OpportunityLineItemService.updateOpportunityLineItem(opportunityLineItem).then( res => {
            this.props.history.push('/opportunityLineItems');
        });
    }

    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changediscountPercentHandler= (event) => {
        this.setState({discountPercent: event.target.value});
    }
    changetotalPriceHandler= (event) => {
        this.setState({totalPrice: event.target.value});
    }

    cancel(){
        this.props.history.push('/opportunityLineItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update OpportunityLineItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> quantity: </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> discountPercent: </label>
                                                <input placeholder="discountPercent" name="discountPercent" className="form-control" value={this.state.discountPercent} onChange={this.changediscountPercentHandler}/>

                                            <label> totalPrice: </label>
                                                <input placeholder="totalPrice" name="totalPrice" className="form-control" value={this.state.totalPrice} onChange={this.changetotalPriceHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateOpportunityLineItem}>Save</button>
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

export default UpdateOpportunityLineItemComponent
