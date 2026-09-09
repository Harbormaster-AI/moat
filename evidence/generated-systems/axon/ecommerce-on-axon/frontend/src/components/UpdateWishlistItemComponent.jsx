import React, { Component } from 'react'
import WishlistItemService from '../services/WishlistItemService';

class UpdateWishlistItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                addedDate: ''
        }
        this.updateWishlistItem = this.updateWishlistItem.bind(this);

        this.changeaddedDateHandler = this.changeaddedDateHandler.bind(this);
    }

    componentDidMount(){
        WishlistItemService.getWishlistItemById(this.state.id).then( (res) =>{
            let wishlistItem = res.data;
            this.setState({
                addedDate: wishlistItem.addedDate
            });
        });
    }

    updateWishlistItem = (e) => {
        e.preventDefault();
        let wishlistItem = {
            wishlistItemId: this.state.id,
            addedDate: this.state.addedDate
        };
        console.log('wishlistItem => ' + JSON.stringify(wishlistItem));
        console.log('id => ' + JSON.stringify(this.state.id));
        WishlistItemService.updateWishlistItem(wishlistItem).then( res => {
            this.props.history.push('/wishlistItems');
        });
    }

    changeaddedDateHandler= (event) => {
        this.setState({addedDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/wishlistItems');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WishlistItem</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> addedDate: </label>
                                                <input type="date" placeholder="addedDate" name="addedDate" className="form-control" value={this.state.addedDate} onChange={this.changeaddedDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWishlistItem}>Save</button>
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

export default UpdateWishlistItemComponent
