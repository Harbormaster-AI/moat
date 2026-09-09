import axios from 'axios';

const INSPECTIONRESULT_API_BASE_URL = "/InspectionResult";

class InspectionResultService {

    getInspectionResults(){
        return axios.get(INSPECTIONRESULT_API_BASE_URL + '/' );
    }

    createInspectionResult(inspectionResult){
        return axios.post(INSPECTIONRESULT_API_BASE_URL  + '/create', inspectionResult);
    }

    getInspectionResultById(inspectionResultId){
        return axios.get(INSPECTIONRESULT_API_BASE_URL + '/load?inspectionResultId=' + inspectionResultId);
    }

    updateInspectionResult(inspectionResult){
        return axios.put(INSPECTIONRESULT_API_BASE_URL + '/update', inspectionResult);
    }

    deleteInspectionResult(inspectionResultId){
        return axios.delete(INSPECTIONRESULT_API_BASE_URL + '/delete?inspectionResultId=' + inspectionResultId);
    }
}

export default new InspectionResultService()