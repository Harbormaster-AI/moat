import axios from 'axios';

const USAGELIMIT_API_BASE_URL = "/UsageLimit";

class UsageLimitService {

    getUsageLimits(){
        return axios.get(USAGELIMIT_API_BASE_URL + '/' );
    }

    createUsageLimit(usageLimit){
        return axios.post(USAGELIMIT_API_BASE_URL  + '/create', usageLimit);
    }

    getUsageLimitById(usageLimitId){
        return axios.get(USAGELIMIT_API_BASE_URL + '/load?usageLimitId=' + usageLimitId);
    }

    updateUsageLimit(usageLimit){
        return axios.put(USAGELIMIT_API_BASE_URL + '/update', usageLimit);
    }

    deleteUsageLimit(usageLimitId){
        return axios.delete(USAGELIMIT_API_BASE_URL + '/delete?usageLimitId=' + usageLimitId);
    }
}

export default new UsageLimitService()