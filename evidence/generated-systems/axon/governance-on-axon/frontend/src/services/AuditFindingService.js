import axios from 'axios';

const AUDITFINDING_API_BASE_URL = "/AuditFinding";

class AuditFindingService {

    getAuditFindings(){
        return axios.get(AUDITFINDING_API_BASE_URL + '/' );
    }

    createAuditFinding(auditFinding){
        return axios.post(AUDITFINDING_API_BASE_URL  + '/create', auditFinding);
    }

    getAuditFindingById(auditFindingId){
        return axios.get(AUDITFINDING_API_BASE_URL + '/load?auditFindingId=' + auditFindingId);
    }

    updateAuditFinding(auditFinding){
        return axios.put(AUDITFINDING_API_BASE_URL + '/update', auditFinding);
    }

    deleteAuditFinding(auditFindingId){
        return axios.delete(AUDITFINDING_API_BASE_URL + '/delete?auditFindingId=' + auditFindingId);
    }
}

export default new AuditFindingService()