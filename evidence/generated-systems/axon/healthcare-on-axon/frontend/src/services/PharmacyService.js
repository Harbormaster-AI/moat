import axios from 'axios';

const PHARMACY_API_BASE_URL = "/Pharmacy";

class PharmacyService {

    getPharmacys(){
        return axios.get(PHARMACY_API_BASE_URL + '/' );
    }

    createPharmacy(pharmacy){
        return axios.post(PHARMACY_API_BASE_URL  + '/create', pharmacy);
    }

    getPharmacyById(pharmacyId){
        return axios.get(PHARMACY_API_BASE_URL + '/load?pharmacyId=' + pharmacyId);
    }

    updatePharmacy(pharmacy){
        return axios.put(PHARMACY_API_BASE_URL + '/update', pharmacy);
    }

    deletePharmacy(pharmacyId){
        return axios.delete(PHARMACY_API_BASE_URL + '/delete?pharmacyId=' + pharmacyId);
    }
}

export default new PharmacyService()