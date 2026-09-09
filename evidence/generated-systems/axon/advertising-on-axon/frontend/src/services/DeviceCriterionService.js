import axios from 'axios';

const DEVICECRITERION_API_BASE_URL = "/DeviceCriterion";

class DeviceCriterionService {

    getDeviceCriterions(){
        return axios.get(DEVICECRITERION_API_BASE_URL + '/' );
    }

    createDeviceCriterion(deviceCriterion){
        return axios.post(DEVICECRITERION_API_BASE_URL  + '/create', deviceCriterion);
    }

    getDeviceCriterionById(deviceCriterionId){
        return axios.get(DEVICECRITERION_API_BASE_URL + '/load?deviceCriterionId=' + deviceCriterionId);
    }

    updateDeviceCriterion(deviceCriterion){
        return axios.put(DEVICECRITERION_API_BASE_URL + '/update', deviceCriterion);
    }

    deleteDeviceCriterion(deviceCriterionId){
        return axios.delete(DEVICECRITERION_API_BASE_URL + '/delete?deviceCriterionId=' + deviceCriterionId);
    }
}

export default new DeviceCriterionService()