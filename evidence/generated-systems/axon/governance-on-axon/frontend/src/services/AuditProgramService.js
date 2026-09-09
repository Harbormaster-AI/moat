import axios from 'axios';

const AUDITPROGRAM_API_BASE_URL = "/AuditProgram";

class AuditProgramService {

    getAuditPrograms(){
        return axios.get(AUDITPROGRAM_API_BASE_URL + '/' );
    }

    createAuditProgram(auditProgram){
        return axios.post(AUDITPROGRAM_API_BASE_URL  + '/create', auditProgram);
    }

    getAuditProgramById(auditProgramId){
        return axios.get(AUDITPROGRAM_API_BASE_URL + '/load?auditProgramId=' + auditProgramId);
    }

    updateAuditProgram(auditProgram){
        return axios.put(AUDITPROGRAM_API_BASE_URL + '/update', auditProgram);
    }

    deleteAuditProgram(auditProgramId){
        return axios.delete(AUDITPROGRAM_API_BASE_URL + '/delete?auditProgramId=' + auditProgramId);
    }
}

export default new AuditProgramService()