import axios from 'axios';

const COMPLIANCEPOLICY_API_BASE_URL = "/CompliancePolicy";

class CompliancePolicyService {

    getCompliancePolicys(){
        return axios.get(COMPLIANCEPOLICY_API_BASE_URL + '/' );
    }

    createCompliancePolicy(compliancePolicy){
        return axios.post(COMPLIANCEPOLICY_API_BASE_URL  + '/create', compliancePolicy);
    }

    getCompliancePolicyById(compliancePolicyId){
        return axios.get(COMPLIANCEPOLICY_API_BASE_URL + '/load?compliancePolicyId=' + compliancePolicyId);
    }

    updateCompliancePolicy(compliancePolicy){
        return axios.put(COMPLIANCEPOLICY_API_BASE_URL + '/update', compliancePolicy);
    }

    deleteCompliancePolicy(compliancePolicyId){
        return axios.delete(COMPLIANCEPOLICY_API_BASE_URL + '/delete?compliancePolicyId=' + compliancePolicyId);
    }
}

export default new CompliancePolicyService()