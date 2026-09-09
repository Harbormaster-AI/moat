import React, { Component } from 'react'
import CustomerService from '../services/CustomerService';

class CreateCustomerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                customerCode: '',
                address: '',
                customerType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecustomerCodeHandler = this.changecustomerCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeCustomerTypeHandler = this.changeCustomerTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCustomer = (e) => {
        e.preventDefault();
        let customer = {
                customerId: this.state.id,
                name: this.state.name,
                customerCode: this.state.customerCode,
                address: this.state.address,
                customerType: this.state.customerType
            };
        console.log('customer => ' + JSON.stringify(customer));

        // step 5
        if(this.state.id === '_add'){
            customer.customerId=''
            CustomerService.createCustomer(customer).then(res =>{
                this.props.history.push('/customers');
            });
        }else{
            CustomerService.updateCustomer(customer).then( res => {
                this.props.history.push('/customers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Customer</h3>
        }else{
            return <h3 className="text-center">Update Customer</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> customerCode:&emsp; </label>
                                                <input placeholder="customerCode" name="customerCode" className="form-control" value={this.state.customerCode} onChange={this.changecustomerCodeHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> CustomerType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCustomer}>Save</button>
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

export default CreateCustomerComponent
