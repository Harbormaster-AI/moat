import React, { Component } from 'react'
import CustomerService from '../services/CustomerService';

class UpdateCustomerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                organizationName: '',
                taxId: '',
                dateOfBirth: '',
                primaryAddress: '',
                customerType: ''
        }
        this.updateCustomer = this.updateCustomer.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeorganizationNameHandler = this.changeorganizationNameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changedateOfBirthHandler = this.changedateOfBirthHandler.bind(this);
        this.changeprimaryAddressHandler = this.changeprimaryAddressHandler.bind(this);
        this.changeCustomerTypeHandler = this.changeCustomerTypeHandler.bind(this);
    }

    componentDidMount(){
        CustomerService.getCustomerById(this.state.id).then( (res) =>{
            let customer = res.data;
            this.setState({
                firstName: customer.firstName,
                lastName: customer.lastName,
                organizationName: customer.organizationName,
                taxId: customer.taxId,
                dateOfBirth: customer.dateOfBirth,
                primaryAddress: customer.primaryAddress,
                customerType: customer.customerType
            });
        });
    }

    updateCustomer = (e) => {
        e.preventDefault();
        let customer = {
            customerId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            organizationName: this.state.organizationName,
            taxId: this.state.taxId,
            dateOfBirth: this.state.dateOfBirth,
            primaryAddress: this.state.primaryAddress,
            customerType: this.state.customerType
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
    changeorganizationNameHandler= (event) => {
        this.setState({organizationName: event.target.value});
    }
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changedateOfBirthHandler= (event) => {
        this.setState({dateOfBirth: event.target.value});
    }
    changeprimaryAddressHandler= (event) => {
        this.setState({primaryAddress: event.target.value});
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
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> organizationName: </label>
                                                <input placeholder="organizationName" name="organizationName" className="form-control" value={this.state.organizationName} onChange={this.changeorganizationNameHandler}/>

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> dateOfBirth: </label>
                                                <input type="date" placeholder="dateOfBirth" name="dateOfBirth" className="form-control" value={this.state.dateOfBirth} onChange={this.changedateOfBirthHandler}/>

                                            <label> primaryAddress: </label>
                                                <input placeholder="primaryAddress" name="primaryAddress" className="form-control" value={this.state.primaryAddress} onChange={this.changeprimaryAddressHandler}/>

                                            <label> CustomerType: </label>
                                                <select value={this.state.customerType} onChange={this.changeCustomerTypeHandler}>
                      <option name="CustomerType" className="form-control" >
                          Individual
                      </option>
                      <option name="CustomerType" className="form-control" >
                          Organization
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
