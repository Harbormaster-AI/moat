import axios from 'axios';

const INFERENCEENDPOINT_API_BASE_URL = "/InferenceEndpoint";

class InferenceEndpointService {

    getInferenceEndpoints(){
        return axios.get(INFERENCEENDPOINT_API_BASE_URL + '/' );
    }

    createInferenceEndpoint(inferenceEndpoint){
        return axios.post(INFERENCEENDPOINT_API_BASE_URL  + '/create', inferenceEndpoint);
    }

    getInferenceEndpointById(inferenceEndpointId){
        return axios.get(INFERENCEENDPOINT_API_BASE_URL + '/load?inferenceEndpointId=' + inferenceEndpointId);
    }

    updateInferenceEndpoint(inferenceEndpoint){
        return axios.put(INFERENCEENDPOINT_API_BASE_URL + '/update', inferenceEndpoint);
    }

    deleteInferenceEndpoint(inferenceEndpointId){
        return axios.delete(INFERENCEENDPOINT_API_BASE_URL + '/delete?inferenceEndpointId=' + inferenceEndpointId);
    }
}

export default new InferenceEndpointService()