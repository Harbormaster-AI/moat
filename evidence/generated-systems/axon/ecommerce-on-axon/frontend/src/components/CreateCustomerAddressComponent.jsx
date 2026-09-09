import React, { Component } from 'react'
import CustomerAddressService from '../services/CustomerAddressService';

class CreateCustomerAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                label: '',
                address: '',
                asDefaultShipping: '',
                asDefaultBilling: ''
        }
        this.changelabelHandler = this.changelabelHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeasDefaultShippingHandler = this.changeasDefaultShippingHandler.bind(this);
        this.changeasDefaultBillingHandler = this.changeasDefaultBillingHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCustomerAddress = (e) => {
        e.preventDefault();
        let customerAddress = {
                customerAddressId: this.state.id,
                label: this.state.label,
                address: this.state.address,
                asDefaultShipping: this.state.asDefaultShipping,
                asDefaultBilling: this.state.asDefaultBilling
            };
        console.log('customerAddress => ' + JSON.stringify(customerAddress));

        // step 5
        if(this.state.id === '_add'){
            customerAddress.customerAddressId=''
            CustomerAddressService.createCustomerAddress(customerAddress).then(res =>{
                this.props.history.push('/customerAddresss');
            });
        }else{
            CustomerAddressService.updateCustomerAddress(customerAddress).then( res => {
                this.props.history.push('/customerAddresss');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CustomerAddress</h3>
        }else{
            return <h3 className="text-center">Update CustomerAddress</h3>
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
                                            <label> label:&emsp; </label>
                                                <input placeholder="label" name="label" className="form-control" value={this.state.label} onChange={this.changelabelHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> asDefaultShipping:&emsp; </label>
                                                <input type="checkbox" placeholder="asDefaultShipping" name="asDefaultShipping" className="form-control" value={this.state.asDefaultShipping} onChange={this.changeasDefaultShippingHandler}/>


                                            <label> asDefaultBilling:&emsp; </label>
                                                <input type="checkbox" placeholder="asDefaultBilling" name="asDefaultBilling" className="form-control" value={this.state.asDefaultBilling} onChange={this.changeasDefaultBillingHandler}/>


                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCustomerAddress}>Save</button>
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

export default CreateCustomerAddressComponent
