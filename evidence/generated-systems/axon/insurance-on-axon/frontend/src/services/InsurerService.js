import axios from 'axios';

const INSURER_API_BASE_URL = "/Insurer";

class InsurerService {

    getInsurers(){
        return axios.get(INSURER_API_BASE_URL + '/' );
    }

    createInsurer(insurer){
        return axios.post(INSURER_API_BASE_URL  + '/create', insurer);
    }

    getInsurerById(insurerId){
        return axios.get(INSURER_API_BASE_URL + '/load?insurerId=' + insurerId);
    }

    updateInsurer(insurer){
        return axios.put(INSURER_API_BASE_URL + '/update', insurer);
    }

    deleteInsurer(insurerId){
        return axios.delete(INSURER_API_BASE_URL + '/delete?insurerId=' + insurerId);
    }
}

export default new InsurerService()