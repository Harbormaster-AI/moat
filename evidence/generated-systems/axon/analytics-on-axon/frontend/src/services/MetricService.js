import axios from 'axios';

const METRIC_API_BASE_URL = "/Metric";

class MetricService {

    getMetrics(){
        return axios.get(METRIC_API_BASE_URL + '/' );
    }

    createMetric(metric){
        return axios.post(METRIC_API_BASE_URL  + '/create', metric);
    }

    getMetricById(metricId){
        return axios.get(METRIC_API_BASE_URL + '/load?metricId=' + metricId);
    }

    updateMetric(metric){
        return axios.put(METRIC_API_BASE_URL + '/update', metric);
    }

    deleteMetric(metricId){
        return axios.delete(METRIC_API_BASE_URL + '/delete?metricId=' + metricId);
    }
}

export default new MetricService()