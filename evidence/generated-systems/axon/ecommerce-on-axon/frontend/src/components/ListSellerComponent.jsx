import React, { Component } from 'react'
import SellerService from '../services/SellerService'

class ListSellerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                sellers: []
        }
        this.addSeller = this.addSeller.bind(this);
        this.editSeller = this.editSeller.bind(this);
        this.deleteSeller = this.deleteSeller.bind(this);
    }

    deleteSeller(id){
        SellerService.deleteSeller(id).then( res => {
            this.setState({sellers: this.state.sellers.filter(seller => seller.sellerId !== id)});
        });
    }
    viewSeller(id){
        this.props.history.push(`/view-seller/${id}`);
    }
    editSeller(id){
        this.props.history.push(`/add-seller/${id}`);
    }

    componentDidMount(){
        SellerService.getSellers().then((res) => {
            this.setState({ sellers: res.data});
        });
    }

    addSeller(){
        this.props.history.push('/add-seller/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Seller List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSeller}> Add Seller</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> SellerCode </th>
                                    <th> ContactEmail </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.sellers.map(
                                        seller => 
                                        <tr key = {seller.sellerId}>
                                             <td> { seller.name } </td>
                                             <td> { seller.sellerCode } </td>
                                             <td> { seller.contactEmail } </td>
                                             <td> { seller.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSeller(seller.sellerId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSeller(seller.sellerId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSeller(seller.sellerId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSellerComponent
