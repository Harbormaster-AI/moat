import React, { Component } from 'react'
import CustomerService from '../services/CustomerService';

class UpdateCustomerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                dateOfBirth: '',
                email: '',
                phone: '',
                address: '',
                taxId: '',
                riskScore: '',
                customerType: ''
        }
        this.updateCustomer = this.updateCustomer.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changedateOfBirthHandler = this.changedateOfBirthHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeriskScoreHandler = this.changeriskScoreHandler.bind(this);
        this.changeCustomerTypeHandler = this.changeCustomerTypeHandler.bind(this);
    }

    componentDidMount(){
        CustomerService.getCustomerById(this.state.id).then( (res) =>{
            let customer = res.data;
            this.setState({
                firstName: customer.firstName,
                lastName: customer.lastName,
                dateOfBirth: customer.dateOfBirth,
                email: customer.email,
                phone: customer.phone,
                address: customer.address,
                taxId: customer.taxId,
                riskScore: customer.riskScore,
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
            dateOfBirth: this.state.dateOfBirth,
            email: this.state.email,
            phone: this.state.phone,
            address: this.state.address,
            taxId: this.state.taxId,
            riskScore: this.state.riskScore,
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
    changedateOfBirthHandler= (event) => {
        this.setState({dateOfBirth: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changeriskScoreHandler= (event) => {
        this.setState({riskScore: event.target.value});
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

                                            <label> dateOfBirth: </label>
                                                <input type="date" placeholder="dateOfBirth" name="dateOfBirth" className="form-control" value={this.state.dateOfBirth} onChange={this.changedateOfBirthHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone: </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> riskScore: </label>
                                                <input placeholder="riskScore" name="riskScore" className="form-control" value={this.state.riskScore} onChange={this.changeriskScoreHandler}/>

                                            <label> CustomerType: </label>
                                                <select value={this.state.customerType} onChange={this.changeCustomerTypeHandler}>
                      <option name="CustomerType" className="form-control" >
                          Individual
                      </option>
                      <option name="CustomerType" className="form-control" >
                          Business
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
