import React, { Component } from 'react'
import WishlistItemService from '../services/WishlistItemService';

class CreateWishlistItemComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                addedDate: ''
        }
        this.changeaddedDateHandler = this.changeaddedDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            WishlistItemService.getWishlistItemById(this.state.id).then( (res) =>{
                let wishlistItem = res.data;
                this.setState({
                    addedDate: wishlistItem.addedDate
                });
            });
        }        
    }
    saveOrUpdateWishlistItem = (e) => {
        e.preventDefault();
        let wishlistItem = {
                wishlistItemId: this.state.id,
                addedDate: this.state.addedDate
            };
        console.log('wishlistItem => ' + JSON.stringify(wishlistItem));

        // step 5
        if(this.state.id === '_add'){
            wishlistItem.wishlistItemId=''
            WishlistItemService.createWishlistItem(wishlistItem).then(res =>{
                this.props.history.push('/wishlistItems');
            });
        }else{
            WishlistItemService.updateWishlistItem(wishlistItem).then( res => {
                this.props.history.push('/wishlistItems');
            });
        }
    }
    
    changeaddedDateHandler= (event) => {
        this.setState({addedDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/wishlistItems');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add WishlistItem</h3>
        }else{
            return <h3 className="text-center">Update WishlistItem</h3>
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
                                            <label> addedDate:&emsp; </label>
                                                <input type="date" placeholder="addedDate" name="addedDate" className="form-control" value={this.state.addedDate} onChange={this.changeaddedDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateWishlistItem}>Save</button>
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

export default CreateWishlistItemComponent
