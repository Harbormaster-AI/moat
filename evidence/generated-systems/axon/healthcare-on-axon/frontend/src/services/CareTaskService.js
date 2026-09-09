import axios from 'axios';

const CARETASK_API_BASE_URL = "/CareTask";

class CareTaskService {

    getCareTasks(){
        return axios.get(CARETASK_API_BASE_URL + '/' );
    }

    createCareTask(careTask){
        return axios.post(CARETASK_API_BASE_URL  + '/create', careTask);
    }

    getCareTaskById(careTaskId){
        return axios.get(CARETASK_API_BASE_URL + '/load?careTaskId=' + careTaskId);
    }

    updateCareTask(careTask){
        return axios.put(CARETASK_API_BASE_URL + '/update', careTask);
    }

    deleteCareTask(careTaskId){
        return axios.delete(CARETASK_API_BASE_URL + '/delete?careTaskId=' + careTaskId);
    }
}

export default new CareTaskService()