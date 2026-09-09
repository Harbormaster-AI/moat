import axios from 'axios';

const CONDITION_API_BASE_URL = "/Condition";

class ConditionService {

    getConditions(){
        return axios.get(CONDITION_API_BASE_URL + '/' );
    }

    createCondition(condition){
        return axios.post(CONDITION_API_BASE_URL  + '/create', condition);
    }

    getConditionById(conditionId){
        return axios.get(CONDITION_API_BASE_URL + '/load?conditionId=' + conditionId);
    }

    updateCondition(condition){
        return axios.put(CONDITION_API_BASE_URL + '/update', condition);
    }

    deleteCondition(conditionId){
        return axios.delete(CONDITION_API_BASE_URL + '/delete?conditionId=' + conditionId);
    }
}

export default new ConditionService()