import axios from 'axios';

const LABRESULT_API_BASE_URL = "/LabResult";

class LabResultService {

    getLabResults(){
        return axios.get(LABRESULT_API_BASE_URL + '/' );
    }

    createLabResult(labResult){
        return axios.post(LABRESULT_API_BASE_URL  + '/create', labResult);
    }

    getLabResultById(labResultId){
        return axios.get(LABRESULT_API_BASE_URL + '/load?labResultId=' + labResultId);
    }

    updateLabResult(labResult){
        return axios.put(LABRESULT_API_BASE_URL + '/update', labResult);
    }

    deleteLabResult(labResultId){
        return axios.delete(LABRESULT_API_BASE_URL + '/delete?labResultId=' + labResultId);
    }
}

export default new LabResultService()