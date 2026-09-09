import axios from 'axios';

const COMPONENT__API_BASE_URL = "/Component_";

class Component_Service {

    getComponent_s(){
        return axios.get(COMPONENT__API_BASE_URL + '/' );
    }

    createComponent_(component_){
        return axios.post(COMPONENT__API_BASE_URL  + '/create', component_);
    }

    getComponent_ById(component_Id){
        return axios.get(COMPONENT__API_BASE_URL + '/load?component_Id=' + component_Id);
    }

    updateComponent_(component_){
        return axios.put(COMPONENT__API_BASE_URL + '/update', component_);
    }

    deleteComponent_(component_Id){
        return axios.delete(COMPONENT__API_BASE_URL + '/delete?component_Id=' + component_Id);
    }
}

export default new Component_Service()