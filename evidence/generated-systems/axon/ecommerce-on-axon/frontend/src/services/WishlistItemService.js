import axios from 'axios';

const WISHLISTITEM_API_BASE_URL = "/WishlistItem";

class WishlistItemService {

    getWishlistItems(){
        return axios.get(WISHLISTITEM_API_BASE_URL + '/' );
    }

    createWishlistItem(wishlistItem){
        return axios.post(WISHLISTITEM_API_BASE_URL  + '/create', wishlistItem);
    }

    getWishlistItemById(wishlistItemId){
        return axios.get(WISHLISTITEM_API_BASE_URL + '/load?wishlistItemId=' + wishlistItemId);
    }

    updateWishlistItem(wishlistItem){
        return axios.put(WISHLISTITEM_API_BASE_URL + '/update', wishlistItem);
    }

    deleteWishlistItem(wishlistItemId){
        return axios.delete(WISHLISTITEM_API_BASE_URL + '/delete?wishlistItemId=' + wishlistItemId);
    }
}

export default new WishlistItemService()