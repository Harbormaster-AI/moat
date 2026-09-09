import axios from 'axios';

const COMPLIANCEALERT_API_BASE_URL = "/ComplianceAlert";

class ComplianceAlertService {

    getComplianceAlerts(){
        return axios.get(COMPLIANCEALERT_API_BASE_URL + '/' );
    }

    createComplianceAlert(complianceAlert){
        return axios.post(COMPLIANCEALERT_API_BASE_URL  + '/create', complianceAlert);
    }

    getComplianceAlertById(complianceAlertId){
        return axios.get(COMPLIANCEALERT_API_BASE_URL + '/load?complianceAlertId=' + complianceAlertId);
    }

    updateComplianceAlert(complianceAlert){
        return axios.put(COMPLIANCEALERT_API_BASE_URL + '/update', complianceAlert);
    }

    deleteComplianceAlert(complianceAlertId){
        return axios.delete(COMPLIANCEALERT_API_BASE_URL + '/delete?complianceAlertId=' + complianceAlertId);
    }
}

export default new ComplianceAlertService()