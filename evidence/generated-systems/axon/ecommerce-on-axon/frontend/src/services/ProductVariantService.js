import axios from 'axios';

const PRODUCTVARIANT_API_BASE_URL = "/ProductVariant";

class ProductVariantService {

    getProductVariants(){
        return axios.get(PRODUCTVARIANT_API_BASE_URL + '/' );
    }

    createProductVariant(productVariant){
        return axios.post(PRODUCTVARIANT_API_BASE_URL  + '/create', productVariant);
    }

    getProductVariantById(productVariantId){
        return axios.get(PRODUCTVARIANT_API_BASE_URL + '/load?productVariantId=' + productVariantId);
    }

    updateProductVariant(productVariant){
        return axios.put(PRODUCTVARIANT_API_BASE_URL + '/update', productVariant);
    }

    deleteProductVariant(productVariantId){
        return axios.delete(PRODUCTVARIANT_API_BASE_URL + '/delete?productVariantId=' + productVariantId);
    }
}

export default new ProductVariantService()