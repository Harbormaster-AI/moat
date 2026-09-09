import axios from 'axios';

const PERFORMANCEMETRIC_API_BASE_URL = "/PerformanceMetric";

class PerformanceMetricService {

    getPerformanceMetrics(){
        return axios.get(PERFORMANCEMETRIC_API_BASE_URL + '/' );
    }

    createPerformanceMetric(performanceMetric){
        return axios.post(PERFORMANCEMETRIC_API_BASE_URL  + '/create', performanceMetric);
    }

    getPerformanceMetricById(performanceMetricId){
        return axios.get(PERFORMANCEMETRIC_API_BASE_URL + '/load?performanceMetricId=' + performanceMetricId);
    }

    updatePerformanceMetric(performanceMetric){
        return axios.put(PERFORMANCEMETRIC_API_BASE_URL + '/update', performanceMetric);
    }

    deletePerformanceMetric(performanceMetricId){
        return axios.delete(PERFORMANCEMETRIC_API_BASE_URL + '/delete?performanceMetricId=' + performanceMetricId);
    }
}

export default new PerformanceMetricService()