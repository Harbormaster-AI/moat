import axios from 'axios';

const AUDITENGAGEMENT_API_BASE_URL = "/AuditEngagement";

class AuditEngagementService {

    getAuditEngagements(){
        return axios.get(AUDITENGAGEMENT_API_BASE_URL + '/' );
    }

    createAuditEngagement(auditEngagement){
        return axios.post(AUDITENGAGEMENT_API_BASE_URL  + '/create', auditEngagement);
    }

    getAuditEngagementById(auditEngagementId){
        return axios.get(AUDITENGAGEMENT_API_BASE_URL + '/load?auditEngagementId=' + auditEngagementId);
    }

    updateAuditEngagement(auditEngagement){
        return axios.put(AUDITENGAGEMENT_API_BASE_URL + '/update', auditEngagement);
    }

    deleteAuditEngagement(auditEngagementId){
        return axios.delete(AUDITENGAGEMENT_API_BASE_URL + '/delete?auditEngagementId=' + auditEngagementId);
    }
}

export default new AuditEngagementService()