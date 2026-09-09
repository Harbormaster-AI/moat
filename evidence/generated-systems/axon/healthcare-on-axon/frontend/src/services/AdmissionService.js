import axios from 'axios';

const ADMISSION_API_BASE_URL = "/Admission";

class AdmissionService {

    getAdmissions(){
        return axios.get(ADMISSION_API_BASE_URL + '/' );
    }

    createAdmission(admission){
        return axios.post(ADMISSION_API_BASE_URL  + '/create', admission);
    }

    getAdmissionById(admissionId){
        return axios.get(ADMISSION_API_BASE_URL + '/load?admissionId=' + admissionId);
    }

    updateAdmission(admission){
        return axios.put(ADMISSION_API_BASE_URL + '/update', admission);
    }

    deleteAdmission(admissionId){
        return axios.delete(ADMISSION_API_BASE_URL + '/delete?admissionId=' + admissionId);
    }
}

export default new AdmissionService()