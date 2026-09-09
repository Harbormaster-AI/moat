import axios from 'axios';

const QUARANTINE_API_BASE_URL = "/Quarantine";

class QuarantineService {

    getQuarantines(){
        return axios.get(QUARANTINE_API_BASE_URL + '/' );
    }

    createQuarantine(quarantine){
        return axios.post(QUARANTINE_API_BASE_URL  + '/create', quarantine);
    }

    getQuarantineById(quarantineId){
        return axios.get(QUARANTINE_API_BASE_URL + '/load?quarantineId=' + quarantineId);
    }

    updateQuarantine(quarantine){
        return axios.put(QUARANTINE_API_BASE_URL + '/update', quarantine);
    }

    deleteQuarantine(quarantineId){
        return axios.delete(QUARANTINE_API_BASE_URL + '/delete?quarantineId=' + quarantineId);
    }
}

export default new QuarantineService()