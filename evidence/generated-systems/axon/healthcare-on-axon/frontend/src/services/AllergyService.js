import axios from 'axios';

const ALLERGY_API_BASE_URL = "/Allergy";

class AllergyService {

    getAllergys(){
        return axios.get(ALLERGY_API_BASE_URL + '/' );
    }

    createAllergy(allergy){
        return axios.post(ALLERGY_API_BASE_URL  + '/create', allergy);
    }

    getAllergyById(allergyId){
        return axios.get(ALLERGY_API_BASE_URL + '/load?allergyId=' + allergyId);
    }

    updateAllergy(allergy){
        return axios.put(ALLERGY_API_BASE_URL + '/update', allergy);
    }

    deleteAllergy(allergyId){
        return axios.delete(ALLERGY_API_BASE_URL + '/delete?allergyId=' + allergyId);
    }
}

export default new AllergyService()