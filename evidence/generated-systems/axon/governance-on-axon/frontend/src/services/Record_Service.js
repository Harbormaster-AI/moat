import axios from 'axios';

const RECORD__API_BASE_URL = "/Record_";

class Record_Service {

    getRecord_s(){
        return axios.get(RECORD__API_BASE_URL + '/' );
    }

    createRecord_(record_){
        return axios.post(RECORD__API_BASE_URL  + '/create', record_);
    }

    getRecord_ById(record_Id){
        return axios.get(RECORD__API_BASE_URL + '/load?record_Id=' + record_Id);
    }

    updateRecord_(record_){
        return axios.put(RECORD__API_BASE_URL + '/update', record_);
    }

    deleteRecord_(record_Id){
        return axios.delete(RECORD__API_BASE_URL + '/delete?record_Id=' + record_Id);
    }
}

export default new Record_Service()