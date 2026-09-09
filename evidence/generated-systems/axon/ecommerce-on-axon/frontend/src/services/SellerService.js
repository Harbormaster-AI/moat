import axios from 'axios';

const SELLER_API_BASE_URL = "/Seller";

class SellerService {

    getSellers(){
        return axios.get(SELLER_API_BASE_URL + '/' );
    }

    createSeller(seller){
        return axios.post(SELLER_API_BASE_URL  + '/create', seller);
    }

    getSellerById(sellerId){
        return axios.get(SELLER_API_BASE_URL + '/load?sellerId=' + sellerId);
    }

    updateSeller(seller){
        return axios.put(SELLER_API_BASE_URL + '/update', seller);
    }

    deleteSeller(sellerId){
        return axios.delete(SELLER_API_BASE_URL + '/delete?sellerId=' + sellerId);
    }
}

export default new SellerService()