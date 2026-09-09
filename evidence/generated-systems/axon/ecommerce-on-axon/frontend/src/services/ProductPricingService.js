import axios from 'axios';

const PRODUCTPRICING_API_BASE_URL = "/ProductPricing";

class ProductPricingService {

    getProductPricings(){
        return axios.get(PRODUCTPRICING_API_BASE_URL + '/' );
    }

    createProductPricing(productPricing){
        return axios.post(PRODUCTPRICING_API_BASE_URL  + '/create', productPricing);
    }

    getProductPricingById(productPricingId){
        return axios.get(PRODUCTPRICING_API_BASE_URL + '/load?productPricingId=' + productPricingId);
    }

    updateProductPricing(productPricing){
        return axios.put(PRODUCTPRICING_API_BASE_URL + '/update', productPricing);
    }

    deleteProductPricing(productPricingId){
        return axios.delete(PRODUCTPRICING_API_BASE_URL + '/delete?productPricingId=' + productPricingId);
    }
}

export default new ProductPricingService()