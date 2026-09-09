import axios from 'axios';

const PERFORMANCECYCLE_API_BASE_URL = "/PerformanceCycle";

class PerformanceCycleService {

    getPerformanceCycles(){
        return axios.get(PERFORMANCECYCLE_API_BASE_URL + '/' );
    }

    createPerformanceCycle(performanceCycle){
        return axios.post(PERFORMANCECYCLE_API_BASE_URL  + '/create', performanceCycle);
    }

    getPerformanceCycleById(performanceCycleId){
        return axios.get(PERFORMANCECYCLE_API_BASE_URL + '/load?performanceCycleId=' + performanceCycleId);
    }

    updatePerformanceCycle(performanceCycle){
        return axios.put(PERFORMANCECYCLE_API_BASE_URL + '/update', performanceCycle);
    }

    deletePerformanceCycle(performanceCycleId){
        return axios.delete(PERFORMANCECYCLE_API_BASE_URL + '/delete?performanceCycleId=' + performanceCycleId);
    }
}

export default new PerformanceCycleService()