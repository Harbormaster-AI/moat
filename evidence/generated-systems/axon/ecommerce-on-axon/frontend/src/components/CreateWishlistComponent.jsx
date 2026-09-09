import React, { Component } from 'react'
import WishlistService from '../services/WishlistService';

class CreateWishlistComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                asPublic: '',
                createdAt: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeasPublicHandler = this.changeasPublicHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WishlistService.getWishlistById(this.state.id).then( (res) =>{
                let wishlist = res.data;
                this.setState({
                    name: wishlist.name,
                    asPublic: wishlist.asPublic,
                    createdAt: wishlist.createdAt
                });
            });
        }        
    }
    saveOrUpdateWishlist = (e) => {
        e.preventDefault();
        let wishlist = {
                wishlistId: this.state.id,
                name: this.state.name,
                asPublic: this.state.asPublic,
                createdAt: this.state.createdAt
            };
        console.log('wishlist => ' + JSON.stringify(wishlist));

        // step 5
        if(this.state.id === '_add'){
            wishlist.wishlistId=''
            WishlistService.createWishlist(wishlist).then(res =>{
                this.props.history.push('/wishlists');
            });
        }else{
            WishlistService.updateWishlist(wishlist).then( res => {
                this.props.history.push('/wishlists');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Wishlist</h3>
        }else{
            return <h3 className="text-center">Update Wishlist</h3>
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

                                            <label> asPublic:&emsp; </label>
                                                <input type="checkbox" placeholder="asPublic" name="asPublic" className="form-control" value={this.state.asPublic} onChange={this.changeasPublicHandler}/>


                                            <label> createdAt:&emsp; </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWishlist}>Save</button>
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

export default CreateWishlistComponent
