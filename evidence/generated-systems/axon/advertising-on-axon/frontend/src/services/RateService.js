import axios from 'axios';

const RATE_API_BASE_URL = "/Rate";

class RateService {

    getRates(){
        return axios.get(RATE_API_BASE_URL + '/' );
    }

    createRate(rate){
        return axios.post(RATE_API_BASE_URL  + '/create', rate);
    }

    getRateById(rateId){
        return axios.get(RATE_API_BASE_URL + '/load?rateId=' + rateId);
    }

    updateRate(rate){
        return axios.put(RATE_API_BASE_URL + '/update', rate);
    }

    deleteRate(rateId){
        return axios.delete(RATE_API_BASE_URL + '/delete?rateId=' + rateId);
    }
}

export default new RateService()