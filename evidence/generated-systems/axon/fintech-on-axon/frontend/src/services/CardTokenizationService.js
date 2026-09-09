import axios from 'axios';

const CARDTOKENIZATION_API_BASE_URL = "/CardTokenization";

class CardTokenizationService {

    getCardTokenizations(){
        return axios.get(CARDTOKENIZATION_API_BASE_URL + '/' );
    }

    createCardTokenization(cardTokenization){
        return axios.post(CARDTOKENIZATION_API_BASE_URL  + '/create', cardTokenization);
    }

    getCardTokenizationById(cardTokenizationId){
        return axios.get(CARDTOKENIZATION_API_BASE_URL + '/load?cardTokenizationId=' + cardTokenizationId);
    }

    updateCardTokenization(cardTokenization){
        return axios.put(CARDTOKENIZATION_API_BASE_URL + '/update', cardTokenization);
    }

    deleteCardTokenization(cardTokenizationId){
        return axios.delete(CARDTOKENIZATION_API_BASE_URL + '/delete?cardTokenizationId=' + cardTokenizationId);
    }
}

export default new CardTokenizationService()