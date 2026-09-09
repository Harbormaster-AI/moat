import axios from 'axios';

const LEAVEREQUEST_API_BASE_URL = "/LeaveRequest";

class LeaveRequestService {

    getLeaveRequests(){
        return axios.get(LEAVEREQUEST_API_BASE_URL + '/' );
    }

    createLeaveRequest(leaveRequest){
        return axios.post(LEAVEREQUEST_API_BASE_URL  + '/create', leaveRequest);
    }

    getLeaveRequestById(leaveRequestId){
        return axios.get(LEAVEREQUEST_API_BASE_URL + '/load?leaveRequestId=' + leaveRequestId);
    }

    updateLeaveRequest(leaveRequest){
        return axios.put(LEAVEREQUEST_API_BASE_URL + '/update', leaveRequest);
    }

    deleteLeaveRequest(leaveRequestId){
        return axios.delete(LEAVEREQUEST_API_BASE_URL + '/delete?leaveRequestId=' + leaveRequestId);
    }
}

export default new LeaveRequestService()