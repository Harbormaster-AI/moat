import axios from 'axios';

const DATASUBJECTREQUEST_API_BASE_URL = "/DataSubjectRequest";

class DataSubjectRequestService {

    getDataSubjectRequests(){
        return axios.get(DATASUBJECTREQUEST_API_BASE_URL + '/' );
    }

    createDataSubjectRequest(dataSubjectRequest){
        return axios.post(DATASUBJECTREQUEST_API_BASE_URL  + '/create', dataSubjectRequest);
    }

    getDataSubjectRequestById(dataSubjectRequestId){
        return axios.get(DATASUBJECTREQUEST_API_BASE_URL + '/load?dataSubjectRequestId=' + dataSubjectRequestId);
    }

    updateDataSubjectRequest(dataSubjectRequest){
        return axios.put(DATASUBJECTREQUEST_API_BASE_URL + '/update', dataSubjectRequest);
    }

    deleteDataSubjectRequest(dataSubjectRequestId){
        return axios.delete(DATASUBJECTREQUEST_API_BASE_URL + '/delete?dataSubjectRequestId=' + dataSubjectRequestId);
    }
}

export default new DataSubjectRequestService()