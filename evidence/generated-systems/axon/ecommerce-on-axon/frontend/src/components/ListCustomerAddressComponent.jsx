import React, { Component } from 'react'
import CustomerAddressService from '../services/CustomerAddressService'

class ListCustomerAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                customerAddresss: []
        }
        this.addCustomerAddress = this.addCustomerAddress.bind(this);
        this.editCustomerAddress = this.editCustomerAddress.bind(this);
        this.deleteCustomerAddress = this.deleteCustomerAddress.bind(this);
    }

    deleteCustomerAddress(id){
        CustomerAddressService.deleteCustomerAddress(id).then( res => {
            this.setState({customerAddresss: this.state.customerAddresss.filter(customerAddress => customerAddress.customerAddressId !== id)});
        });
    }
    viewCustomerAddress(id){
        this.props.history.push(`/view-customerAddress/${id}`);
    }
    editCustomerAddress(id){
        this.props.history.push(`/add-customerAddress/${id}`);
    }

    componentDidMount(){
        CustomerAddressService.getCustomerAddresss().then((res) => {
            this.setState({ customerAddresss: res.data});
        });
    }

    addCustomerAddress(){
        this.props.history.push('/add-customerAddress/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CustomerAddress List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCustomerAddress}> Add CustomerAddress</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Label </th>
                                    <th> Address </th>
                                    <th> AsDefaultShipping </th>
                                    <th> AsDefaultBilling </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.customerAddresss.map(
                                        customerAddress => 
                                        <tr key = {customerAddress.customerAddressId}>
                                             <td> { customerAddress.label } </td>
                                             <td> { customerAddress.address } </td>
                                             <td> { customerAddress.asDefaultShipping } </td>
                                             <td> { customerAddress.asDefaultBilling } </td>
                                             <td>
                                                 <button onClick={ () => this.editCustomerAddress(customerAddress.customerAddressId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCustomerAddress(customerAddress.customerAddressId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCustomerAddress(customerAddress.customerAddressId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCustomerAddressComponent
