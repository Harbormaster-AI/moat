import axios from 'axios';

const CONTROLTEST__API_BASE_URL = "/ControlTest_";

class ControlTest_Service {

    getControlTest_s(){
        return axios.get(CONTROLTEST__API_BASE_URL + '/' );
    }

    createControlTest_(controlTest_){
        return axios.post(CONTROLTEST__API_BASE_URL  + '/create', controlTest_);
    }

    getControlTest_ById(controlTest_Id){
        return axios.get(CONTROLTEST__API_BASE_URL + '/load?controlTest_Id=' + controlTest_Id);
    }

    updateControlTest_(controlTest_){
        return axios.put(CONTROLTEST__API_BASE_URL + '/update', controlTest_);
    }

    deleteControlTest_(controlTest_Id){
        return axios.delete(CONTROLTEST__API_BASE_URL + '/delete?controlTest_Id=' + controlTest_Id);
    }
}

export default new ControlTest_Service()