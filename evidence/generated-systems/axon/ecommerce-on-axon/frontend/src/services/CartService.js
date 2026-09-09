import axios from 'axios';

const CART_API_BASE_URL = "/Cart";

class CartService {

    getCarts(){
        return axios.get(CART_API_BASE_URL + '/' );
    }

    createCart(cart){
        return axios.post(CART_API_BASE_URL  + '/create', cart);
    }

    getCartById(cartId){
        return axios.get(CART_API_BASE_URL + '/load?cartId=' + cartId);
    }

    updateCart(cart){
        return axios.put(CART_API_BASE_URL + '/update', cart);
    }

    deleteCart(cartId){
        return axios.delete(CART_API_BASE_URL + '/delete?cartId=' + cartId);
    }
}

export default new CartService()