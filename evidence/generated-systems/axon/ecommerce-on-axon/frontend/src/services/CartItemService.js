import axios from 'axios';

const CARTITEM_API_BASE_URL = "/CartItem";

class CartItemService {

    getCartItems(){
        return axios.get(CARTITEM_API_BASE_URL + '/' );
    }

    createCartItem(cartItem){
        return axios.post(CARTITEM_API_BASE_URL  + '/create', cartItem);
    }

    getCartItemById(cartItemId){
        return axios.get(CARTITEM_API_BASE_URL + '/load?cartItemId=' + cartItemId);
    }

    updateCartItem(cartItem){
        return axios.put(CARTITEM_API_BASE_URL + '/update', cartItem);
    }

    deleteCartItem(cartItemId){
        return axios.delete(CARTITEM_API_BASE_URL + '/delete?cartItemId=' + cartItemId);
    }
}

export default new CartItemService()