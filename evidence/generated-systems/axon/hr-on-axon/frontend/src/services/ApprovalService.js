import axios from 'axios';

const APPROVAL_API_BASE_URL = "/Approval";

class ApprovalService {

    getApprovals(){
        return axios.get(APPROVAL_API_BASE_URL + '/' );
    }

    createApproval(approval){
        return axios.post(APPROVAL_API_BASE_URL  + '/create', approval);
    }

    getApprovalById(approvalId){
        return axios.get(APPROVAL_API_BASE_URL + '/load?approvalId=' + approvalId);
    }

    updateApproval(approval){
        return axios.put(APPROVAL_API_BASE_URL + '/update', approval);
    }

    deleteApproval(approvalId){
        return axios.delete(APPROVAL_API_BASE_URL + '/delete?approvalId=' + approvalId);
    }
}

export default new ApprovalService()