import axios from 'axios';

const BACKGROUNDCHECK_API_BASE_URL = "/BackgroundCheck";

class BackgroundCheckService {

    getBackgroundChecks(){
        return axios.get(BACKGROUNDCHECK_API_BASE_URL + '/' );
    }

    createBackgroundCheck(backgroundCheck){
        return axios.post(BACKGROUNDCHECK_API_BASE_URL  + '/create', backgroundCheck);
    }

    getBackgroundCheckById(backgroundCheckId){
        return axios.get(BACKGROUNDCHECK_API_BASE_URL + '/load?backgroundCheckId=' + backgroundCheckId);
    }

    updateBackgroundCheck(backgroundCheck){
        return axios.put(BACKGROUNDCHECK_API_BASE_URL + '/update', backgroundCheck);
    }

    deleteBackgroundCheck(backgroundCheckId){
        return axios.delete(BACKGROUNDCHECK_API_BASE_URL + '/delete?backgroundCheckId=' + backgroundCheckId);
    }
}

export default new BackgroundCheckService()