import axios from 'axios';

const PAYOUT_API_BASE_URL = "/Payout";

class PayoutService {

    getPayouts(){
        return axios.get(PAYOUT_API_BASE_URL + '/' );
    }

    createPayout(payout){
        return axios.post(PAYOUT_API_BASE_URL  + '/create', payout);
    }

    getPayoutById(payoutId){
        return axios.get(PAYOUT_API_BASE_URL + '/load?payoutId=' + payoutId);
    }

    updatePayout(payout){
        return axios.put(PAYOUT_API_BASE_URL + '/update', payout);
    }

    deletePayout(payoutId){
        return axios.delete(PAYOUT_API_BASE_URL + '/delete?payoutId=' + payoutId);
    }
}

export default new PayoutService()