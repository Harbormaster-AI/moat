import React, { Component } from 'react'
import CustomerAddressService from '../services/CustomerAddressService';

class UpdateCustomerAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                label: '',
                address: '',
                asDefaultShipping: '',
                asDefaultBilling: ''
        }
        this.updateCustomerAddress = this.updateCustomerAddress.bind(this);

        this.changelabelHandler = this.changelabelHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeasDefaultShippingHandler = this.changeasDefaultShippingHandler.bind(this);
        this.changeasDefaultBillingHandler = this.changeasDefaultBillingHandler.bind(this);
    }

    componentDidMount(){
        CustomerAddressService.getCustomerAddressById(this.state.id).then( (res) =>{
            let customerAddress = res.data;
            this.setState({
                label: customerAddress.label,
                address: customerAddress.address,
                asDefaultShipping: customerAddress.asDefaultShipping,
                asDefaultBilling: customerAddress.asDefaultBilling
            });
        });
    }

    updateCustomerAddress = (e) => {
        e.preventDefault();
        let customerAddress = {
            customerAddressId: this.state.id,
            label: this.state.label,
            address: this.state.address,
            asDefaultShipping: this.state.asDefaultShipping,
            asDefaultBilling: this.state.asDefaultBilling
        };
        console.log('customerAddress => ' + JSON.stringify(customerAddress));
        console.log('id => ' + JSON.stringify(this.state.id));
        CustomerAddressService.updateCustomerAddress(customerAddress).then( res => {
            this.props.history.push('/customerAddresss');
        });
    }

    changelabelHandler= (event) => {
        this.setState({label: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeasDefaultShippingHandler= (event) => {
        this.setState({asDefaultShipping: event.target.value});
    }
    changeasDefaultBillingHandler= (event) => {
        this.setState({asDefaultBilling: event.target.value});
    }

    cancel(){
        this.props.history.push('/customerAddresss');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CustomerAddress</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> label: </label>
                                                <input placeholder="label" name="label" className="form-control" value={this.state.label} onChange={this.changelabelHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> asDefaultShipping: </label>
                                                <input type="checkbox" placeholder="asDefaultShipping" name="asDefaultShipping" className="form-control" value={this.state.asDefaultShipping} onChange={this.changeasDefaultShippingHandler}/>


                                            <label> asDefaultBilling: </label>
                                                <input type="checkbox" placeholder="asDefaultBilling" name="asDefaultBilling" className="form-control" value={this.state.asDefaultBilling} onChange={this.changeasDefaultBillingHandler}/>


                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCustomerAddress}>Save</button>
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

export default UpdateCustomerAddressComponent
