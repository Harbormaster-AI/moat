import React, { Component } from 'react'
import CustomerService from '../services/CustomerService';

class UpdateCustomerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                customerCode: '',
                address: '',
                customerType: ''
        }
        this.updateCustomer = this.updateCustomer.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecustomerCodeHandler = this.changecustomerCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeCustomerTypeHandler = this.changeCustomerTypeHandler.bind(this);
    }

    componentDidMount(){
        CustomerService.getCustomerById(this.state.id).then( (res) =>{
            let customer = res.data;
            this.setState({
                name: customer.name,
                customerCode: customer.customerCode,
                address: customer.address,
                customerType: customer.customerType
            });
        });
    }

    updateCustomer = (e) => {
        e.preventDefault();
        let customer = {
            customerId: this.state.id,
            name: this.state.name,
            customerCode: this.state.customerCode,
            address: this.state.address,
            customerType: this.state.customerType
        };
        console.log('customer => ' + JSON.stringify(customer));
        console.log('id => ' + JSON.stringify(this.state.id));
        CustomerService.updateCustomer(customer).then( res => {
            this.props.history.push('/customers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecustomerCodeHandler= (event) => {
        this.setState({customerCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeCustomerTypeHandler= (event) => {
        this.setState({customerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/customers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Customer</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> customerCode: </label>
                                                <input placeholder="customerCode" name="customerCode" className="form-control" value={this.state.customerCode} onChange={this.changecustomerCodeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> CustomerType: </label>
                                                <select value={this.state.customerType} onChange={this.changeCustomerTypeHandler}>
                      <option name="CustomerType" className="form-control" >
                          Distributor
                      </option>
                      <option name="CustomerType" className="form-control" >
                          OEM
                      </option>
                      <option name="CustomerType" className="form-control" >
                          Retailer
                      </option>
                      <option name="CustomerType" className="form-control" >
                          Direct
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCustomer}>Save</button>
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

export default UpdateCustomerComponent
