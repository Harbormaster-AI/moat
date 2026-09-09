import axios from 'axios';

const DATAPROCESSINGACTIVITY_API_BASE_URL = "/DataProcessingActivity";

class DataProcessingActivityService {

    getDataProcessingActivitys(){
        return axios.get(DATAPROCESSINGACTIVITY_API_BASE_URL + '/' );
    }

    createDataProcessingActivity(dataProcessingActivity){
        return axios.post(DATAPROCESSINGACTIVITY_API_BASE_URL  + '/create', dataProcessingActivity);
    }

    getDataProcessingActivityById(dataProcessingActivityId){
        return axios.get(DATAPROCESSINGACTIVITY_API_BASE_URL + '/load?dataProcessingActivityId=' + dataProcessingActivityId);
    }

    updateDataProcessingActivity(dataProcessingActivity){
        return axios.put(DATAPROCESSINGACTIVITY_API_BASE_URL + '/update', dataProcessingActivity);
    }

    deleteDataProcessingActivity(dataProcessingActivityId){
        return axios.delete(DATAPROCESSINGACTIVITY_API_BASE_URL + '/delete?dataProcessingActivityId=' + dataProcessingActivityId);
    }
}

export default new DataProcessingActivityService()