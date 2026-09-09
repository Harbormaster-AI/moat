import axios from 'axios';

const RUNPARAMETER_API_BASE_URL = "/RunParameter";

class RunParameterService {

    getRunParameters(){
        return axios.get(RUNPARAMETER_API_BASE_URL + '/' );
    }

    createRunParameter(runParameter){
        return axios.post(RUNPARAMETER_API_BASE_URL  + '/create', runParameter);
    }

    getRunParameterById(runParameterId){
        return axios.get(RUNPARAMETER_API_BASE_URL + '/load?runParameterId=' + runParameterId);
    }

    updateRunParameter(runParameter){
        return axios.put(RUNPARAMETER_API_BASE_URL + '/update', runParameter);
    }

    deleteRunParameter(runParameterId){
        return axios.delete(RUNPARAMETER_API_BASE_URL + '/delete?runParameterId=' + runParameterId);
    }
}

export default new RunParameterService()