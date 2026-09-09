import axios from 'axios';

const CORRECTIVEACTION_API_BASE_URL = "/CorrectiveAction";

class CorrectiveActionService {

    getCorrectiveActions(){
        return axios.get(CORRECTIVEACTION_API_BASE_URL + '/' );
    }

    createCorrectiveAction(correctiveAction){
        return axios.post(CORRECTIVEACTION_API_BASE_URL  + '/create', correctiveAction);
    }

    getCorrectiveActionById(correctiveActionId){
        return axios.get(CORRECTIVEACTION_API_BASE_URL + '/load?correctiveActionId=' + correctiveActionId);
    }

    updateCorrectiveAction(correctiveAction){
        return axios.put(CORRECTIVEACTION_API_BASE_URL + '/update', correctiveAction);
    }

    deleteCorrectiveAction(correctiveActionId){
        return axios.delete(CORRECTIVEACTION_API_BASE_URL + '/delete?correctiveActionId=' + correctiveActionId);
    }
}

export default new CorrectiveActionService()