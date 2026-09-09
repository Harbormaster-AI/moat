import axios from 'axios';

const OPERATION_API_BASE_URL = "/Operation";

class OperationService {

    getOperations(){
        return axios.get(OPERATION_API_BASE_URL + '/' );
    }

    createOperation(operation){
        return axios.post(OPERATION_API_BASE_URL  + '/create', operation);
    }

    getOperationById(operationId){
        return axios.get(OPERATION_API_BASE_URL + '/load?operationId=' + operationId);
    }

    updateOperation(operation){
        return axios.put(OPERATION_API_BASE_URL + '/update', operation);
    }

    deleteOperation(operationId){
        return axios.delete(OPERATION_API_BASE_URL + '/delete?operationId=' + operationId);
    }
}

export default new OperationService()