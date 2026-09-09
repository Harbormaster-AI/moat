import axios from 'axios';

const RATECARD_API_BASE_URL = "/RateCard";

class RateCardService {

    getRateCards(){
        return axios.get(RATECARD_API_BASE_URL + '/' );
    }

    createRateCard(rateCard){
        return axios.post(RATECARD_API_BASE_URL  + '/create', rateCard);
    }

    getRateCardById(rateCardId){
        return axios.get(RATECARD_API_BASE_URL + '/load?rateCardId=' + rateCardId);
    }

    updateRateCard(rateCard){
        return axios.put(RATECARD_API_BASE_URL + '/update', rateCard);
    }

    deleteRateCard(rateCardId){
        return axios.delete(RATECARD_API_BASE_URL + '/delete?rateCardId=' + rateCardId);
    }
}

export default new RateCardService()