import React, { Component } from 'react'
import WishlistService from '../services/WishlistService'

class ListWishlistComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                wishlists: []
        }
        this.addWishlist = this.addWishlist.bind(this);
        this.editWishlist = this.editWishlist.bind(this);
        this.deleteWishlist = this.deleteWishlist.bind(this);
    }

    deleteWishlist(id){
        WishlistService.deleteWishlist(id).then( res => {
            this.setState({wishlists: this.state.wishlists.filter(wishlist => wishlist.wishlistId !== id)});
        });
    }
    viewWishlist(id){
        this.props.history.push(`/view-wishlist/${id}`);
    }
    editWishlist(id){
        this.props.history.push(`/add-wishlist/${id}`);
    }

    componentDidMount(){
        WishlistService.getWishlists().then((res) => {
            this.setState({ wishlists: res.data});
        });
    }

    addWishlist(){
        this.props.history.push('/add-wishlist/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Wishlist List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWishlist}> Add Wishlist</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> AsPublic </th>
                                    <th> CreatedAt </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.wishlists.map(
                                        wishlist => 
                                        <tr key = {wishlist.wishlistId}>
                                             <td> { wishlist.name } </td>
                                             <td> { wishlist.asPublic } </td>
                                             <td> { wishlist.createdAt } </td>
                                             <td>
                                                 <button onClick={ () => this.editWishlist(wishlist.wishlistId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWishlist(wishlist.wishlistId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWishlist(wishlist.wishlistId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWishlistComponent
