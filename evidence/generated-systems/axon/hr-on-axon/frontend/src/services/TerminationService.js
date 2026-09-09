import axios from 'axios';

const TERMINATION_API_BASE_URL = "/Termination";

class TerminationService {

    getTerminations(){
        return axios.get(TERMINATION_API_BASE_URL + '/' );
    }

    createTermination(termination){
        return axios.post(TERMINATION_API_BASE_URL  + '/create', termination);
    }

    getTerminationById(terminationId){
        return axios.get(TERMINATION_API_BASE_URL + '/load?terminationId=' + terminationId);
    }

    updateTermination(termination){
        return axios.put(TERMINATION_API_BASE_URL + '/update', termination);
    }

    deleteTermination(terminationId){
        return axios.delete(TERMINATION_API_BASE_URL + '/delete?terminationId=' + terminationId);
    }
}

export default new TerminationService()