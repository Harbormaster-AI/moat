import axios from 'axios';

const QUOTE_API_BASE_URL = "/Quote";

class QuoteService {

    getQuotes(){
        return axios.get(QUOTE_API_BASE_URL + '/' );
    }

    createQuote(quote){
        return axios.post(QUOTE_API_BASE_URL  + '/create', quote);
    }

    getQuoteById(quoteId){
        return axios.get(QUOTE_API_BASE_URL + '/load?quoteId=' + quoteId);
    }

    updateQuote(quote){
        return axios.put(QUOTE_API_BASE_URL + '/update', quote);
    }

    deleteQuote(quoteId){
        return axios.delete(QUOTE_API_BASE_URL + '/delete?quoteId=' + quoteId);
    }
}

export default new QuoteService()