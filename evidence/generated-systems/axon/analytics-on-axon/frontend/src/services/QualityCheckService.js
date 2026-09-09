import axios from 'axios';

const QUALITYCHECK_API_BASE_URL = "/QualityCheck";

class QualityCheckService {

    getQualityChecks(){
        return axios.get(QUALITYCHECK_API_BASE_URL + '/' );
    }

    createQualityCheck(qualityCheck){
        return axios.post(QUALITYCHECK_API_BASE_URL  + '/create', qualityCheck);
    }

    getQualityCheckById(qualityCheckId){
        return axios.get(QUALITYCHECK_API_BASE_URL + '/load?qualityCheckId=' + qualityCheckId);
    }

    updateQualityCheck(qualityCheck){
        return axios.put(QUALITYCHECK_API_BASE_URL + '/update', qualityCheck);
    }

    deleteQualityCheck(qualityCheckId){
        return axios.delete(QUALITYCHECK_API_BASE_URL + '/delete?qualityCheckId=' + qualityCheckId);
    }
}

export default new QualityCheckService()