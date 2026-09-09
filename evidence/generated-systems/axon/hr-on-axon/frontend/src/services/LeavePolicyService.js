import axios from 'axios';

const LEAVEPOLICY_API_BASE_URL = "/LeavePolicy";

class LeavePolicyService {

    getLeavePolicys(){
        return axios.get(LEAVEPOLICY_API_BASE_URL + '/' );
    }

    createLeavePolicy(leavePolicy){
        return axios.post(LEAVEPOLICY_API_BASE_URL  + '/create', leavePolicy);
    }

    getLeavePolicyById(leavePolicyId){
        return axios.get(LEAVEPOLICY_API_BASE_URL + '/load?leavePolicyId=' + leavePolicyId);
    }

    updateLeavePolicy(leavePolicy){
        return axios.put(LEAVEPOLICY_API_BASE_URL + '/update', leavePolicy);
    }

    deleteLeavePolicy(leavePolicyId){
        return axios.delete(LEAVEPOLICY_API_BASE_URL + '/delete?leavePolicyId=' + leavePolicyId);
    }
}

export default new LeavePolicyService()