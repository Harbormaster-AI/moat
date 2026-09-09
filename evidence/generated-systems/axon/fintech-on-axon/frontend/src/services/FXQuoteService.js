import axios from 'axios';

const FXQUOTE_API_BASE_URL = "/FXQuote";

class FXQuoteService {

    getFXQuotes(){
        return axios.get(FXQUOTE_API_BASE_URL + '/' );
    }

    createFXQuote(fXQuote){
        return axios.post(FXQUOTE_API_BASE_URL  + '/create', fXQuote);
    }

    getFXQuoteById(fXQuoteId){
        return axios.get(FXQUOTE_API_BASE_URL + '/load?fXQuoteId=' + fXQuoteId);
    }

    updateFXQuote(fXQuote){
        return axios.put(FXQUOTE_API_BASE_URL + '/update', fXQuote);
    }

    deleteFXQuote(fXQuoteId){
        return axios.delete(FXQUOTE_API_BASE_URL + '/delete?fXQuoteId=' + fXQuoteId);
    }
}

export default new FXQuoteService()