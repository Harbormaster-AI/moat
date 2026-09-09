import axios from 'axios';

const PRODUCTOFFERING_API_BASE_URL = "/ProductOffering";

class ProductOfferingService {

    getProductOfferings(){
        return axios.get(PRODUCTOFFERING_API_BASE_URL + '/' );
    }

    createProductOffering(productOffering){
        return axios.post(PRODUCTOFFERING_API_BASE_URL  + '/create', productOffering);
    }

    getProductOfferingById(productOfferingId){
        return axios.get(PRODUCTOFFERING_API_BASE_URL + '/load?productOfferingId=' + productOfferingId);
    }

    updateProductOffering(productOffering){
        return axios.put(PRODUCTOFFERING_API_BASE_URL + '/update', productOffering);
    }

    deleteProductOffering(productOfferingId){
        return axios.delete(PRODUCTOFFERING_API_BASE_URL + '/delete?productOfferingId=' + productOfferingId);
    }
}

export default new ProductOfferingService()