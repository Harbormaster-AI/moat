import axios from 'axios';

const BIQUERY_API_BASE_URL = "/BIQuery";

class BIQueryService {

    getBIQuerys(){
        return axios.get(BIQUERY_API_BASE_URL + '/' );
    }

    createBIQuery(bIQuery){
        return axios.post(BIQUERY_API_BASE_URL  + '/create', bIQuery);
    }

    getBIQueryById(bIQueryId){
        return axios.get(BIQUERY_API_BASE_URL + '/load?bIQueryId=' + bIQueryId);
    }

    updateBIQuery(bIQuery){
        return axios.put(BIQUERY_API_BASE_URL + '/update', bIQuery);
    }

    deleteBIQuery(bIQueryId){
        return axios.delete(BIQUERY_API_BASE_URL + '/delete?bIQueryId=' + bIQueryId);
    }
}

export default new BIQueryService()