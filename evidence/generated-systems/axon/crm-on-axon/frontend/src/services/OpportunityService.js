import axios from 'axios';

const OPPORTUNITY_API_BASE_URL = "/Opportunity";

class OpportunityService {

    getOpportunitys(){
        return axios.get(OPPORTUNITY_API_BASE_URL + '/' );
    }

    createOpportunity(opportunity){
        return axios.post(OPPORTUNITY_API_BASE_URL  + '/create', opportunity);
    }

    getOpportunityById(opportunityId){
        return axios.get(OPPORTUNITY_API_BASE_URL + '/load?opportunityId=' + opportunityId);
    }

    updateOpportunity(opportunity){
        return axios.put(OPPORTUNITY_API_BASE_URL + '/update', opportunity);
    }

    deleteOpportunity(opportunityId){
        return axios.delete(OPPORTUNITY_API_BASE_URL + '/delete?opportunityId=' + opportunityId);
    }
}

export default new OpportunityService()