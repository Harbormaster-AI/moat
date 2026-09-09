import axios from 'axios';

const SYSTEM__API_BASE_URL = "/System_";

class System_Service {

    getSystem_s(){
        return axios.get(SYSTEM__API_BASE_URL + '/' );
    }

    createSystem_(system_){
        return axios.post(SYSTEM__API_BASE_URL  + '/create', system_);
    }

    getSystem_ById(system_Id){
        return axios.get(SYSTEM__API_BASE_URL + '/load?system_Id=' + system_Id);
    }

    updateSystem_(system_){
        return axios.put(SYSTEM__API_BASE_URL + '/update', system_);
    }

    deleteSystem_(system_Id){
        return axios.delete(SYSTEM__API_BASE_URL + '/delete?system_Id=' + system_Id);
    }
}

export default new System_Service()