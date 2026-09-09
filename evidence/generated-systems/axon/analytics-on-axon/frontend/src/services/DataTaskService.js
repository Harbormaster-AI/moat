import axios from 'axios';

const DATATASK_API_BASE_URL = "/DataTask";

class DataTaskService {

    getDataTasks(){
        return axios.get(DATATASK_API_BASE_URL + '/' );
    }

    createDataTask(dataTask){
        return axios.post(DATATASK_API_BASE_URL  + '/create', dataTask);
    }

    getDataTaskById(dataTaskId){
        return axios.get(DATATASK_API_BASE_URL + '/load?dataTaskId=' + dataTaskId);
    }

    updateDataTask(dataTask){
        return axios.put(DATATASK_API_BASE_URL + '/update', dataTask);
    }

    deleteDataTask(dataTaskId){
        return axios.delete(DATATASK_API_BASE_URL + '/delete?dataTaskId=' + dataTaskId);
    }
}

export default new DataTaskService()