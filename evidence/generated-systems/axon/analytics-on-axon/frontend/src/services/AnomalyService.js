import axios from 'axios';

const ANOMALY_API_BASE_URL = "/Anomaly";

class AnomalyService {

    getAnomalys(){
        return axios.get(ANOMALY_API_BASE_URL + '/' );
    }

    createAnomaly(anomaly){
        return axios.post(ANOMALY_API_BASE_URL  + '/create', anomaly);
    }

    getAnomalyById(anomalyId){
        return axios.get(ANOMALY_API_BASE_URL + '/load?anomalyId=' + anomalyId);
    }

    updateAnomaly(anomaly){
        return axios.put(ANOMALY_API_BASE_URL + '/update', anomaly);
    }

    deleteAnomaly(anomalyId){
        return axios.delete(ANOMALY_API_BASE_URL + '/delete?anomalyId=' + anomalyId);
    }
}

export default new AnomalyService()