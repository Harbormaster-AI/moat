import axios from 'axios';

const PRICEBOOKENTRY_API_BASE_URL = "/PriceBookEntry";

class PriceBookEntryService {

    getPriceBookEntrys(){
        return axios.get(PRICEBOOKENTRY_API_BASE_URL + '/' );
    }

    createPriceBookEntry(priceBookEntry){
        return axios.post(PRICEBOOKENTRY_API_BASE_URL  + '/create', priceBookEntry);
    }

    getPriceBookEntryById(priceBookEntryId){
        return axios.get(PRICEBOOKENTRY_API_BASE_URL + '/load?priceBookEntryId=' + priceBookEntryId);
    }

    updatePriceBookEntry(priceBookEntry){
        return axios.put(PRICEBOOKENTRY_API_BASE_URL + '/update', priceBookEntry);
    }

    deletePriceBookEntry(priceBookEntryId){
        return axios.delete(PRICEBOOKENTRY_API_BASE_URL + '/delete?priceBookEntryId=' + priceBookEntryId);
    }
}

export default new PriceBookEntryService()