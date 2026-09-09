import React, { Component } from 'react'
import CustomerService from '../services/CustomerService';

class UpdateCustomerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                email: '',
                phone: '',
                marketingOptIn: '',
                customerGroup: ''
        }
        this.updateCustomer = this.updateCustomer.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changemarketingOptInHandler = this.changemarketingOptInHandler.bind(this);
        this.changeCustomerGroupHandler = this.changeCustomerGroupHandler.bind(this);
    }

    componentDidMount(){
        CustomerService.getCustomerById(this.state.id).then( (res) =>{
            let customer = res.data;
            this.setState({
                firstName: customer.firstName,
                lastName: customer.lastName,
                email: customer.email,
                phone: customer.phone,
                marketingOptIn: customer.marketingOptIn,
                customerGroup: customer.customerGroup
            });
        });
    }

    updateCustomer = (e) => {
        e.preventDefault();
        let customer = {
            customerId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            email: this.state.email,
            phone: this.state.phone,
            marketingOptIn: this.state.marketingOptIn,
            customerGroup: this.state.customerGroup
        };
        console.log('customer => ' + JSON.stringify(customer));
        console.log('id => ' + JSON.stringify(this.state.id));
        CustomerService.updateCustomer(customer).then( res => {
            this.props.history.push('/customers');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changemarketingOptInHandler= (event) => {
        this.setState({marketingOptIn: event.target.value});
    }
    changeCustomerGroupHandler= (event) => {
        this.setState({customerGroup: event.target.value});
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
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone: </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> marketingOptIn: </label>
                                                <input type="checkbox" placeholder="marketingOptIn" name="marketingOptIn" className="form-control" value={this.state.marketingOptIn} onChange={this.changemarketingOptInHandler}/>


                                            <label> CustomerGroup: </label>
                                                <select value={this.state.customerGroup} onChange={this.changeCustomerGroupHandler}>
                      <option name="CustomerGroup" className="form-control" >
                          Retail
                      </option>
                      <option name="CustomerGroup" className="form-control" >
                          Wholesale
                      </option>
                      <option name="CustomerGroup" className="form-control" >
                          VIP
                      </option>
                      <option name="CustomerGroup" className="form-control" >
                          Employee
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
