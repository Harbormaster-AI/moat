import axios from 'axios';

const COMPLIANCEPROGRAM_API_BASE_URL = "/ComplianceProgram";

class ComplianceProgramService {

    getCompliancePrograms(){
        return axios.get(COMPLIANCEPROGRAM_API_BASE_URL + '/' );
    }

    createComplianceProgram(complianceProgram){
        return axios.post(COMPLIANCEPROGRAM_API_BASE_URL  + '/create', complianceProgram);
    }

    getComplianceProgramById(complianceProgramId){
        return axios.get(COMPLIANCEPROGRAM_API_BASE_URL + '/load?complianceProgramId=' + complianceProgramId);
    }

    updateComplianceProgram(complianceProgram){
        return axios.put(COMPLIANCEPROGRAM_API_BASE_URL + '/update', complianceProgram);
    }

    deleteComplianceProgram(complianceProgramId){
        return axios.delete(COMPLIANCEPROGRAM_API_BASE_URL + '/delete?complianceProgramId=' + complianceProgramId);
    }
}

export default new ComplianceProgramService()