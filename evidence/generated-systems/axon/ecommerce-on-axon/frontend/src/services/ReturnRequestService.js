import axios from 'axios';

const RETURNREQUEST_API_BASE_URL = "/ReturnRequest";

class ReturnRequestService {

    getReturnRequests(){
        return axios.get(RETURNREQUEST_API_BASE_URL + '/' );
    }

    createReturnRequest(returnRequest){
        return axios.post(RETURNREQUEST_API_BASE_URL  + '/create', returnRequest);
    }

    getReturnRequestById(returnRequestId){
        return axios.get(RETURNREQUEST_API_BASE_URL + '/load?returnRequestId=' + returnRequestId);
    }

    updateReturnRequest(returnRequest){
        return axios.put(RETURNREQUEST_API_BASE_URL + '/update', returnRequest);
    }

    deleteReturnRequest(returnRequestId){
        return axios.delete(RETURNREQUEST_API_BASE_URL + '/delete?returnRequestId=' + returnRequestId);
    }
}

export default new ReturnRequestService()