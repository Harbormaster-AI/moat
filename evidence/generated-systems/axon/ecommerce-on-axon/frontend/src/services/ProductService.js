import axios from 'axios';

const PRODUCT_API_BASE_URL = "/Product";

class ProductService {

    getProducts(){
        return axios.get(PRODUCT_API_BASE_URL + '/' );
    }

    createProduct(product){
        return axios.post(PRODUCT_API_BASE_URL  + '/create', product);
    }

    getProductById(productId){
        return axios.get(PRODUCT_API_BASE_URL + '/load?productId=' + productId);
    }

    updateProduct(product){
        return axios.put(PRODUCT_API_BASE_URL + '/update', product);
    }

    deleteProduct(productId){
        return axios.delete(PRODUCT_API_BASE_URL + '/delete?productId=' + productId);
    }
}

export default new ProductService()