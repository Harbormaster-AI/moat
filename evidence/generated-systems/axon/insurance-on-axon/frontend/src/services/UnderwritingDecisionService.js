import axios from 'axios';

const UNDERWRITINGDECISION_API_BASE_URL = "/UnderwritingDecision";

class UnderwritingDecisionService {

    getUnderwritingDecisions(){
        return axios.get(UNDERWRITINGDECISION_API_BASE_URL + '/' );
    }

    createUnderwritingDecision(underwritingDecision){
        return axios.post(UNDERWRITINGDECISION_API_BASE_URL  + '/create', underwritingDecision);
    }

    getUnderwritingDecisionById(underwritingDecisionId){
        return axios.get(UNDERWRITINGDECISION_API_BASE_URL + '/load?underwritingDecisionId=' + underwritingDecisionId);
    }

    updateUnderwritingDecision(underwritingDecision){
        return axios.put(UNDERWRITINGDECISION_API_BASE_URL + '/update', underwritingDecision);
    }

    deleteUnderwritingDecision(underwritingDecisionId){
        return axios.delete(UNDERWRITINGDECISION_API_BASE_URL + '/delete?underwritingDecisionId=' + underwritingDecisionId);
    }
}

export default new UnderwritingDecisionService()