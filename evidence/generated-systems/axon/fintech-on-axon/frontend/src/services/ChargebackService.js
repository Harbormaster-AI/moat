import axios from 'axios';

const CHARGEBACK_API_BASE_URL = "/Chargeback";

class ChargebackService {

    getChargebacks(){
        return axios.get(CHARGEBACK_API_BASE_URL + '/' );
    }

    createChargeback(chargeback){
        return axios.post(CHARGEBACK_API_BASE_URL  + '/create', chargeback);
    }

    getChargebackById(chargebackId){
        return axios.get(CHARGEBACK_API_BASE_URL + '/load?chargebackId=' + chargebackId);
    }

    updateChargeback(chargeback){
        return axios.put(CHARGEBACK_API_BASE_URL + '/update', chargeback);
    }

    deleteChargeback(chargebackId){
        return axios.delete(CHARGEBACK_API_BASE_URL + '/delete?chargebackId=' + chargebackId);
    }
}

export default new ChargebackService()