import axios from 'axios';

const DEPENDENT_API_BASE_URL = "/Dependent";

class DependentService {

    getDependents(){
        return axios.get(DEPENDENT_API_BASE_URL + '/' );
    }

    createDependent(dependent){
        return axios.post(DEPENDENT_API_BASE_URL  + '/create', dependent);
    }

    getDependentById(dependentId){
        return axios.get(DEPENDENT_API_BASE_URL + '/load?dependentId=' + dependentId);
    }

    updateDependent(dependent){
        return axios.put(DEPENDENT_API_BASE_URL + '/update', dependent);
    }

    deleteDependent(dependentId){
        return axios.delete(DEPENDENT_API_BASE_URL + '/delete?dependentId=' + dependentId);
    }
}

export default new DependentService()