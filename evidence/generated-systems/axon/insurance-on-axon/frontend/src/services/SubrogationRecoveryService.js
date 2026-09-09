import axios from 'axios';

const SUBROGATIONRECOVERY_API_BASE_URL = "/SubrogationRecovery";

class SubrogationRecoveryService {

    getSubrogationRecoverys(){
        return axios.get(SUBROGATIONRECOVERY_API_BASE_URL + '/' );
    }

    createSubrogationRecovery(subrogationRecovery){
        return axios.post(SUBROGATIONRECOVERY_API_BASE_URL  + '/create', subrogationRecovery);
    }

    getSubrogationRecoveryById(subrogationRecoveryId){
        return axios.get(SUBROGATIONRECOVERY_API_BASE_URL + '/load?subrogationRecoveryId=' + subrogationRecoveryId);
    }

    updateSubrogationRecovery(subrogationRecovery){
        return axios.put(SUBROGATIONRECOVERY_API_BASE_URL + '/update', subrogationRecovery);
    }

    deleteSubrogationRecovery(subrogationRecoveryId){
        return axios.delete(SUBROGATIONRECOVERY_API_BASE_URL + '/delete?subrogationRecoveryId=' + subrogationRecoveryId);
    }
}

export default new SubrogationRecoveryService()