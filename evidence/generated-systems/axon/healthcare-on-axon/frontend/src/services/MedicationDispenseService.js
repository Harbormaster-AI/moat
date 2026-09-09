import axios from 'axios';

const MEDICATIONDISPENSE_API_BASE_URL = "/MedicationDispense";

class MedicationDispenseService {

    getMedicationDispenses(){
        return axios.get(MEDICATIONDISPENSE_API_BASE_URL + '/' );
    }

    createMedicationDispense(medicationDispense){
        return axios.post(MEDICATIONDISPENSE_API_BASE_URL  + '/create', medicationDispense);
    }

    getMedicationDispenseById(medicationDispenseId){
        return axios.get(MEDICATIONDISPENSE_API_BASE_URL + '/load?medicationDispenseId=' + medicationDispenseId);
    }

    updateMedicationDispense(medicationDispense){
        return axios.put(MEDICATIONDISPENSE_API_BASE_URL + '/update', medicationDispense);
    }

    deleteMedicationDispense(medicationDispenseId){
        return axios.delete(MEDICATIONDISPENSE_API_BASE_URL + '/delete?medicationDispenseId=' + medicationDispenseId);
    }
}

export default new MedicationDispenseService()