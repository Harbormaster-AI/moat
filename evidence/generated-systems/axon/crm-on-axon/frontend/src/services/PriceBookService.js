import axios from 'axios';

const PRICEBOOK_API_BASE_URL = "/PriceBook";

class PriceBookService {

    getPriceBooks(){
        return axios.get(PRICEBOOK_API_BASE_URL + '/' );
    }

    createPriceBook(priceBook){
        return axios.post(PRICEBOOK_API_BASE_URL  + '/create', priceBook);
    }

    getPriceBookById(priceBookId){
        return axios.get(PRICEBOOK_API_BASE_URL + '/load?priceBookId=' + priceBookId);
    }

    updatePriceBook(priceBook){
        return axios.put(PRICEBOOK_API_BASE_URL + '/update', priceBook);
    }

    deletePriceBook(priceBookId){
        return axios.delete(PRICEBOOK_API_BASE_URL + '/delete?priceBookId=' + priceBookId);
    }
}

export default new PriceBookService()