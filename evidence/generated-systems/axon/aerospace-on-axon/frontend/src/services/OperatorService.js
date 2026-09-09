import axios from 'axios';

const OPERATOR_API_BASE_URL = "/Operator";

class OperatorService {

    getOperators(){
        return axios.get(OPERATOR_API_BASE_URL + '/' );
    }

    createOperator(operator){
        return axios.post(OPERATOR_API_BASE_URL  + '/create', operator);
    }

    getOperatorById(operatorId){
        return axios.get(OPERATOR_API_BASE_URL + '/load?operatorId=' + operatorId);
    }

    updateOperator(operator){
        return axios.put(OPERATOR_API_BASE_URL + '/update', operator);
    }

    deleteOperator(operatorId){
        return axios.delete(OPERATOR_API_BASE_URL + '/delete?operatorId=' + operatorId);
    }
}

export default new OperatorService()