import axios from 'axios';

const CREATIVEAPPROVAL_API_BASE_URL = "/CreativeApproval";

class CreativeApprovalService {

    getCreativeApprovals(){
        return axios.get(CREATIVEAPPROVAL_API_BASE_URL + '/' );
    }

    createCreativeApproval(creativeApproval){
        return axios.post(CREATIVEAPPROVAL_API_BASE_URL  + '/create', creativeApproval);
    }

    getCreativeApprovalById(creativeApprovalId){
        return axios.get(CREATIVEAPPROVAL_API_BASE_URL + '/load?creativeApprovalId=' + creativeApprovalId);
    }

    updateCreativeApproval(creativeApproval){
        return axios.put(CREATIVEAPPROVAL_API_BASE_URL + '/update', creativeApproval);
    }

    deleteCreativeApproval(creativeApprovalId){
        return axios.delete(CREATIVEAPPROVAL_API_BASE_URL + '/delete?creativeApprovalId=' + creativeApprovalId);
    }
}

export default new CreativeApprovalService()