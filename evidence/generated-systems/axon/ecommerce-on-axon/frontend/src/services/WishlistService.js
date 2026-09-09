import axios from 'axios';

const WISHLIST_API_BASE_URL = "/Wishlist";

class WishlistService {

    getWishlists(){
        return axios.get(WISHLIST_API_BASE_URL + '/' );
    }

    createWishlist(wishlist){
        return axios.post(WISHLIST_API_BASE_URL  + '/create', wishlist);
    }

    getWishlistById(wishlistId){
        return axios.get(WISHLIST_API_BASE_URL + '/load?wishlistId=' + wishlistId);
    }

    updateWishlist(wishlist){
        return axios.put(WISHLIST_API_BASE_URL + '/update', wishlist);
    }

    deleteWishlist(wishlistId){
        return axios.delete(WISHLIST_API_BASE_URL + '/delete?wishlistId=' + wishlistId);
    }
}

export default new WishlistService()