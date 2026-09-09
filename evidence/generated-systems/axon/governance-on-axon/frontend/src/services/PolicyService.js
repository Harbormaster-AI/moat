import axios from 'axios';

const POLICY_API_BASE_URL = "/Policy";

class PolicyService {

    getPolicys(){
        return axios.get(POLICY_API_BASE_URL + '/' );
    }

    createPolicy(policy){
        return axios.post(POLICY_API_BASE_URL  + '/create', policy);
    }

    getPolicyById(policyId){
        return axios.get(POLICY_API_BASE_URL + '/load?policyId=' + policyId);
    }

    updatePolicy(policy){
        return axios.put(POLICY_API_BASE_URL + '/update', policy);
    }

    deletePolicy(policyId){
        return axios.delete(POLICY_API_BASE_URL + '/delete?policyId=' + policyId);
    }
}

export default new PolicyService()