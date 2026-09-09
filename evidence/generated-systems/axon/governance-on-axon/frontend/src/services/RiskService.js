import axios from 'axios';

const RISK_API_BASE_URL = "/Risk";

class RiskService {

    getRisks(){
        return axios.get(RISK_API_BASE_URL + '/' );
    }

    createRisk(risk){
        return axios.post(RISK_API_BASE_URL  + '/create', risk);
    }

    getRiskById(riskId){
        return axios.get(RISK_API_BASE_URL + '/load?riskId=' + riskId);
    }

    updateRisk(risk){
        return axios.put(RISK_API_BASE_URL + '/update', risk);
    }

    deleteRisk(riskId){
        return axios.delete(RISK_API_BASE_URL + '/delete?riskId=' + riskId);
    }
}

export default new RiskService()