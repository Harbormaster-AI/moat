import axios from 'axios';

const WORKAUTHORIZATION_API_BASE_URL = "/WorkAuthorization";

class WorkAuthorizationService {

    getWorkAuthorizations(){
        return axios.get(WORKAUTHORIZATION_API_BASE_URL + '/' );
    }

    createWorkAuthorization(workAuthorization){
        return axios.post(WORKAUTHORIZATION_API_BASE_URL  + '/create', workAuthorization);
    }

    getWorkAuthorizationById(workAuthorizationId){
        return axios.get(WORKAUTHORIZATION_API_BASE_URL + '/load?workAuthorizationId=' + workAuthorizationId);
    }

    updateWorkAuthorization(workAuthorization){
        return axios.put(WORKAUTHORIZATION_API_BASE_URL + '/update', workAuthorization);
    }

    deleteWorkAuthorization(workAuthorizationId){
        return axios.delete(WORKAUTHORIZATION_API_BASE_URL + '/delete?workAuthorizationId=' + workAuthorizationId);
    }
}

export default new WorkAuthorizationService()