import axios from 'axios';

const HEALTHSYSTEM_API_BASE_URL = "/HealthSystem";

class HealthSystemService {

    getHealthSystems(){
        return axios.get(HEALTHSYSTEM_API_BASE_URL + '/' );
    }

    createHealthSystem(healthSystem){
        return axios.post(HEALTHSYSTEM_API_BASE_URL  + '/create', healthSystem);
    }

    getHealthSystemById(healthSystemId){
        return axios.get(HEALTHSYSTEM_API_BASE_URL + '/load?healthSystemId=' + healthSystemId);
    }

    updateHealthSystem(healthSystem){
        return axios.put(HEALTHSYSTEM_API_BASE_URL + '/update', healthSystem);
    }

    deleteHealthSystem(healthSystemId){
        return axios.delete(HEALTHSYSTEM_API_BASE_URL + '/delete?healthSystemId=' + healthSystemId);
    }
}

export default new HealthSystemService()