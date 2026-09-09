import axios from 'axios';

const REPLENISHMENTPOLICY_API_BASE_URL = "/ReplenishmentPolicy";

class ReplenishmentPolicyService {

    getReplenishmentPolicys(){
        return axios.get(REPLENISHMENTPOLICY_API_BASE_URL + '/' );
    }

    createReplenishmentPolicy(replenishmentPolicy){
        return axios.post(REPLENISHMENTPOLICY_API_BASE_URL  + '/create', replenishmentPolicy);
    }

    getReplenishmentPolicyById(replenishmentPolicyId){
        return axios.get(REPLENISHMENTPOLICY_API_BASE_URL + '/load?replenishmentPolicyId=' + replenishmentPolicyId);
    }

    updateReplenishmentPolicy(replenishmentPolicy){
        return axios.put(REPLENISHMENTPOLICY_API_BASE_URL + '/update', replenishmentPolicy);
    }

    deleteReplenishmentPolicy(replenishmentPolicyId){
        return axios.delete(REPLENISHMENTPOLICY_API_BASE_URL + '/delete?replenishmentPolicyId=' + replenishmentPolicyId);
    }
}

export default new ReplenishmentPolicyService()