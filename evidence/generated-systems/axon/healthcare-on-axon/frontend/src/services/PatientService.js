import axios from 'axios';

const PATIENT_API_BASE_URL = "/Patient";

class PatientService {

    getPatients(){
        return axios.get(PATIENT_API_BASE_URL + '/' );
    }

    createPatient(patient){
        return axios.post(PATIENT_API_BASE_URL  + '/create', patient);
    }

    getPatientById(patientId){
        return axios.get(PATIENT_API_BASE_URL + '/load?patientId=' + patientId);
    }

    updatePatient(patient){
        return axios.put(PATIENT_API_BASE_URL + '/update', patient);
    }

    deletePatient(patientId){
        return axios.delete(PATIENT_API_BASE_URL + '/delete?patientId=' + patientId);
    }
}

export default new PatientService()