import React, { Component } from 'react'
import WishlistItemService from '../services/WishlistItemService'

class ListWishlistItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                wishlistItems: []
        }
        this.addWishlistItem = this.addWishlistItem.bind(this);
        this.editWishlistItem = this.editWishlistItem.bind(this);
        this.deleteWishlistItem = this.deleteWishlistItem.bind(this);
    }

    deleteWishlistItem(id){
        WishlistItemService.deleteWishlistItem(id).then( res => {
            this.setState({wishlistItems: this.state.wishlistItems.filter(wishlistItem => wishlistItem.wishlistItemId !== id)});
        });
    }
    viewWishlistItem(id){
        this.props.history.push(`/view-wishlistItem/${id}`);
    }
    editWishlistItem(id){
        this.props.history.push(`/add-wishlistItem/${id}`);
    }

    componentDidMount(){
        WishlistItemService.getWishlistItems().then((res) => {
            this.setState({ wishlistItems: res.data});
        });
    }

    addWishlistItem(){
        this.props.history.push('/add-wishlistItem/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WishlistItem List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWishlistItem}> Add WishlistItem</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AddedDate </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.wishlistItems.map(
                                        wishlistItem => 
                                        <tr key = {wishlistItem.wishlistItemId}>
                                             <td> { wishlistItem.addedDate } </td>
                                             <td>
                                                 <button onClick={ () => this.editWishlistItem(wishlistItem.wishlistItemId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWishlistItem(wishlistItem.wishlistItemId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWishlistItem(wishlistItem.wishlistItemId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWishlistItemComponent
