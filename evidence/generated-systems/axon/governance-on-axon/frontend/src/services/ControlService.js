import axios from 'axios';

const CONTROL_API_BASE_URL = "/Control";

class ControlService {

    getControls(){
        return axios.get(CONTROL_API_BASE_URL + '/' );
    }

    createControl(control){
        return axios.post(CONTROL_API_BASE_URL  + '/create', control);
    }

    getControlById(controlId){
        return axios.get(CONTROL_API_BASE_URL + '/load?controlId=' + controlId);
    }

    updateControl(control){
        return axios.put(CONTROL_API_BASE_URL + '/update', control);
    }

    deleteControl(controlId){
        return axios.delete(CONTROL_API_BASE_URL + '/delete?controlId=' + controlId);
    }
}

export default new ControlService()