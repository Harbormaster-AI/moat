import axios from 'axios';

const FRAUDSIGNAL_API_BASE_URL = "/FraudSignal";

class FraudSignalService {

    getFraudSignals(){
        return axios.get(FRAUDSIGNAL_API_BASE_URL + '/' );
    }

    createFraudSignal(fraudSignal){
        return axios.post(FRAUDSIGNAL_API_BASE_URL  + '/create', fraudSignal);
    }

    getFraudSignalById(fraudSignalId){
        return axios.get(FRAUDSIGNAL_API_BASE_URL + '/load?fraudSignalId=' + fraudSignalId);
    }

    updateFraudSignal(fraudSignal){
        return axios.put(FRAUDSIGNAL_API_BASE_URL + '/update', fraudSignal);
    }

    deleteFraudSignal(fraudSignalId){
        return axios.delete(FRAUDSIGNAL_API_BASE_URL + '/delete?fraudSignalId=' + fraudSignalId);
    }
}

export default new FraudSignalService()