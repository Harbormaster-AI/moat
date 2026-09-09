import axios from 'axios';

const CLINICIAN_API_BASE_URL = "/Clinician";

class ClinicianService {

    getClinicians(){
        return axios.get(CLINICIAN_API_BASE_URL + '/' );
    }

    createClinician(clinician){
        return axios.post(CLINICIAN_API_BASE_URL  + '/create', clinician);
    }

    getClinicianById(clinicianId){
        return axios.get(CLINICIAN_API_BASE_URL + '/load?clinicianId=' + clinicianId);
    }

    updateClinician(clinician){
        return axios.put(CLINICIAN_API_BASE_URL + '/update', clinician);
    }

    deleteClinician(clinicianId){
        return axios.delete(CLINICIAN_API_BASE_URL + '/delete?clinicianId=' + clinicianId);
    }
}

export default new ClinicianService()