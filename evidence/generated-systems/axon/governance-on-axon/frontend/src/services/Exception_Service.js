import axios from 'axios';

const EXCEPTION__API_BASE_URL = "/Exception_";

class Exception_Service {

    getException_s(){
        return axios.get(EXCEPTION__API_BASE_URL + '/' );
    }

    createException_(exception_){
        return axios.post(EXCEPTION__API_BASE_URL  + '/create', exception_);
    }

    getException_ById(exception_Id){
        return axios.get(EXCEPTION__API_BASE_URL + '/load?exception_Id=' + exception_Id);
    }

    updateException_(exception_){
        return axios.put(EXCEPTION__API_BASE_URL + '/update', exception_);
    }

    deleteException_(exception_Id){
        return axios.delete(EXCEPTION__API_BASE_URL + '/delete?exception_Id=' + exception_Id);
    }
}

export default new Exception_Service()