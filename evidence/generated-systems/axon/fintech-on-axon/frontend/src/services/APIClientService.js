import axios from 'axios';

const APICLIENT_API_BASE_URL = "/APIClient";

class APIClientService {

    getAPIClients(){
        return axios.get(APICLIENT_API_BASE_URL + '/' );
    }

    createAPIClient(aPIClient){
        return axios.post(APICLIENT_API_BASE_URL  + '/create', aPIClient);
    }

    getAPIClientById(aPIClientId){
        return axios.get(APICLIENT_API_BASE_URL + '/load?aPIClientId=' + aPIClientId);
    }

    updateAPIClient(aPIClient){
        return axios.put(APICLIENT_API_BASE_URL + '/update', aPIClient);
    }

    deleteAPIClient(aPIClientId){
        return axios.delete(APICLIENT_API_BASE_URL + '/delete?aPIClientId=' + aPIClientId);
    }
}

export default new APIClientService()