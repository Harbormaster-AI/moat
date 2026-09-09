import React, { Component } from 'react'
import VerifiedAddressService from '../services/VerifiedAddressService'

class ListVerifiedAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                verifiedAddresss: []
        }
        this.addVerifiedAddress = this.addVerifiedAddress.bind(this);
        this.editVerifiedAddress = this.editVerifiedAddress.bind(this);
        this.deleteVerifiedAddress = this.deleteVerifiedAddress.bind(this);
    }

    deleteVerifiedAddress(id){
        VerifiedAddressService.deleteVerifiedAddress(id).then( res => {
            this.setState({verifiedAddresss: this.state.verifiedAddresss.filter(verifiedAddress => verifiedAddress.verifiedAddressId !== id)});
        });
    }
    viewVerifiedAddress(id){
        this.props.history.push(`/view-verifiedAddress/${id}`);
    }
    editVerifiedAddress(id){
        this.props.history.push(`/add-verifiedAddress/${id}`);
    }

    componentDidMount(){
        VerifiedAddressService.getVerifiedAddresss().then((res) => {
            this.setState({ verifiedAddresss: res.data});
        });
    }

    addVerifiedAddress(){
        this.props.history.push('/add-verifiedAddress/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">VerifiedAddress List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addVerifiedAddress}> Add VerifiedAddress</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Address </th>
                                    <th> VerifiedAt </th>
                                    <th> VerificationStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.verifiedAddresss.map(
                                        verifiedAddress => 
                                        <tr key = {verifiedAddress.verifiedAddressId}>
                                             <td> { verifiedAddress.address } </td>
                                             <td> { verifiedAddress.verifiedAt } </td>
                                             <td> { verifiedAddress.verificationStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editVerifiedAddress(verifiedAddress.verifiedAddressId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteVerifiedAddress(verifiedAddress.verifiedAddressId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewVerifiedAddress(verifiedAddress.verifiedAddressId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListVerifiedAddressComponent
