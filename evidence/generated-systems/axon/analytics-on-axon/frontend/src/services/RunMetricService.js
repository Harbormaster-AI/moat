import axios from 'axios';

const RUNMETRIC_API_BASE_URL = "/RunMetric";

class RunMetricService {

    getRunMetrics(){
        return axios.get(RUNMETRIC_API_BASE_URL + '/' );
    }

    createRunMetric(runMetric){
        return axios.post(RUNMETRIC_API_BASE_URL  + '/create', runMetric);
    }

    getRunMetricById(runMetricId){
        return axios.get(RUNMETRIC_API_BASE_URL + '/load?runMetricId=' + runMetricId);
    }

    updateRunMetric(runMetric){
        return axios.put(RUNMETRIC_API_BASE_URL + '/update', runMetric);
    }

    deleteRunMetric(runMetricId){
        return axios.delete(RUNMETRIC_API_BASE_URL + '/delete?runMetricId=' + runMetricId);
    }
}

export default new RunMetricService()