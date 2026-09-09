import axios from 'axios';

const SHIPPINGMETHOD_API_BASE_URL = "/ShippingMethod";

class ShippingMethodService {

    getShippingMethods(){
        return axios.get(SHIPPINGMETHOD_API_BASE_URL + '/' );
    }

    createShippingMethod(shippingMethod){
        return axios.post(SHIPPINGMETHOD_API_BASE_URL  + '/create', shippingMethod);
    }

    getShippingMethodById(shippingMethodId){
        return axios.get(SHIPPINGMETHOD_API_BASE_URL + '/load?shippingMethodId=' + shippingMethodId);
    }

    updateShippingMethod(shippingMethod){
        return axios.put(SHIPPINGMETHOD_API_BASE_URL + '/update', shippingMethod);
    }

    deleteShippingMethod(shippingMethodId){
        return axios.delete(SHIPPINGMETHOD_API_BASE_URL + '/delete?shippingMethodId=' + shippingMethodId);
    }
}

export default new ShippingMethodService()