import axios from 'axios';

const CASE__API_BASE_URL = "/Case_";

class Case_Service {

    getCase_s(){
        return axios.get(CASE__API_BASE_URL + '/' );
    }

    createCase_(case_){
        return axios.post(CASE__API_BASE_URL  + '/create', case_);
    }

    getCase_ById(case_Id){
        return axios.get(CASE__API_BASE_URL + '/load?case_Id=' + case_Id);
    }

    updateCase_(case_){
        return axios.put(CASE__API_BASE_URL + '/update', case_);
    }

    deleteCase_(case_Id){
        return axios.delete(CASE__API_BASE_URL + '/delete?case_Id=' + case_Id);
    }
}

export default new Case_Service()