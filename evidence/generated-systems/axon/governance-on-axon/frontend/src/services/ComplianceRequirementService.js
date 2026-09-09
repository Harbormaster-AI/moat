import axios from 'axios';

const COMPLIANCEREQUIREMENT_API_BASE_URL = "/ComplianceRequirement";

class ComplianceRequirementService {

    getComplianceRequirements(){
        return axios.get(COMPLIANCEREQUIREMENT_API_BASE_URL + '/' );
    }

    createComplianceRequirement(complianceRequirement){
        return axios.post(COMPLIANCEREQUIREMENT_API_BASE_URL  + '/create', complianceRequirement);
    }

    getComplianceRequirementById(complianceRequirementId){
        return axios.get(COMPLIANCEREQUIREMENT_API_BASE_URL + '/load?complianceRequirementId=' + complianceRequirementId);
    }

    updateComplianceRequirement(complianceRequirement){
        return axios.put(COMPLIANCEREQUIREMENT_API_BASE_URL + '/update', complianceRequirement);
    }

    deleteComplianceRequirement(complianceRequirementId){
        return axios.delete(COMPLIANCEREQUIREMENT_API_BASE_URL + '/delete?complianceRequirementId=' + complianceRequirementId);
    }
}

export default new ComplianceRequirementService()