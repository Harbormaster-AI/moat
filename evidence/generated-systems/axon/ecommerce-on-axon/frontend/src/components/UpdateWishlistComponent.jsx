import React, { Component } from 'react'
import WishlistService from '../services/WishlistService';

class UpdateWishlistComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                asPublic: '',
                createdAt: ''
        }
        this.updateWishlist = this.updateWishlist.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeasPublicHandler = this.changeasPublicHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
    }

    componentDidMount(){
        WishlistService.getWishlistById(this.state.id).then( (res) =>{
            let wishlist = res.data;
            this.setState({
                name: wishlist.name,
                asPublic: wishlist.asPublic,
                createdAt: wishlist.createdAt
            });
        });
    }

    updateWishlist = (e) => {
        e.preventDefault();
        let wishlist = {
            wishlistId: this.state.id,
            name: this.state.name,
            asPublic: this.state.asPublic,
            createdAt: this.state.createdAt
        };
        console.log('wishlist => ' + JSON.stringify(wishlist));
        console.log('id => ' + JSON.stringify(this.state.id));
        WishlistService.updateWishlist(wishlist).then( res => {
            this.props.history.push('/wishlists');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeasPublicHandler= (event) => {
        this.setState({asPublic: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }

    cancel(){
        this.props.history.push('/wishlists');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Wishlist</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> asPublic: </label>
                                                <input type="checkbox" placeholder="asPublic" name="asPublic" className="form-control" value={this.state.asPublic} onChange={this.changeasPublicHandler}/>


                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWishlist}>Save</button>
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

export default UpdateWishlistComponent
