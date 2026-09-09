import axios from 'axios';

const MERCHANT_API_BASE_URL = "/Merchant";

class MerchantService {

    getMerchants(){
        return axios.get(MERCHANT_API_BASE_URL + '/' );
    }

    createMerchant(merchant){
        return axios.post(MERCHANT_API_BASE_URL  + '/create', merchant);
    }

    getMerchantById(merchantId){
        return axios.get(MERCHANT_API_BASE_URL + '/load?merchantId=' + merchantId);
    }

    updateMerchant(merchant){
        return axios.put(MERCHANT_API_BASE_URL + '/update', merchant);
    }

    deleteMerchant(merchantId){
        return axios.delete(MERCHANT_API_BASE_URL + '/delete?merchantId=' + merchantId);
    }
}

export default new MerchantService()