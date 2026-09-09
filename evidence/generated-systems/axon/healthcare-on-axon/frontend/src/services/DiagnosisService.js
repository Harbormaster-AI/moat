import axios from 'axios';

const DIAGNOSIS_API_BASE_URL = "/Diagnosis";

class DiagnosisService {

    getDiagnosiss(){
        return axios.get(DIAGNOSIS_API_BASE_URL + '/' );
    }

    createDiagnosis(diagnosis){
        return axios.post(DIAGNOSIS_API_BASE_URL  + '/create', diagnosis);
    }

    getDiagnosisById(diagnosisId){
        return axios.get(DIAGNOSIS_API_BASE_URL + '/load?diagnosisId=' + diagnosisId);
    }

    updateDiagnosis(diagnosis){
        return axios.put(DIAGNOSIS_API_BASE_URL + '/update', diagnosis);
    }

    deleteDiagnosis(diagnosisId){
        return axios.delete(DIAGNOSIS_API_BASE_URL + '/delete?diagnosisId=' + diagnosisId);
    }
}

export default new DiagnosisService()