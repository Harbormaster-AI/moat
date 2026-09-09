import axios from 'axios';

const QUOTELINEITEM_API_BASE_URL = "/QuoteLineItem";

class QuoteLineItemService {

    getQuoteLineItems(){
        return axios.get(QUOTELINEITEM_API_BASE_URL + '/' );
    }

    createQuoteLineItem(quoteLineItem){
        return axios.post(QUOTELINEITEM_API_BASE_URL  + '/create', quoteLineItem);
    }

    getQuoteLineItemById(quoteLineItemId){
        return axios.get(QUOTELINEITEM_API_BASE_URL + '/load?quoteLineItemId=' + quoteLineItemId);
    }

    updateQuoteLineItem(quoteLineItem){
        return axios.put(QUOTELINEITEM_API_BASE_URL + '/update', quoteLineItem);
    }

    deleteQuoteLineItem(quoteLineItemId){
        return axios.delete(QUOTELINEITEM_API_BASE_URL + '/delete?quoteLineItemId=' + quoteLineItemId);
    }
}

export default new QuoteLineItemService()