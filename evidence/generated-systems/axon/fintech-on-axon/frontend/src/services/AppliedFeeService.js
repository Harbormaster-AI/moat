import axios from 'axios';

const APPLIEDFEE_API_BASE_URL = "/AppliedFee";

class AppliedFeeService {

    getAppliedFees(){
        return axios.get(APPLIEDFEE_API_BASE_URL + '/' );
    }

    createAppliedFee(appliedFee){
        return axios.post(APPLIEDFEE_API_BASE_URL  + '/create', appliedFee);
    }

    getAppliedFeeById(appliedFeeId){
        return axios.get(APPLIEDFEE_API_BASE_URL + '/load?appliedFeeId=' + appliedFeeId);
    }

    updateAppliedFee(appliedFee){
        return axios.put(APPLIEDFEE_API_BASE_URL + '/update', appliedFee);
    }

    deleteAppliedFee(appliedFeeId){
        return axios.delete(APPLIEDFEE_API_BASE_URL + '/delete?appliedFeeId=' + appliedFeeId);
    }
}

export default new AppliedFeeService()