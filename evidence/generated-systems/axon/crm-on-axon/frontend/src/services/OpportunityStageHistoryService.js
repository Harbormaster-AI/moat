import axios from 'axios';

const OPPORTUNITYSTAGEHISTORY_API_BASE_URL = "/OpportunityStageHistory";

class OpportunityStageHistoryService {

    getOpportunityStageHistorys(){
        return axios.get(OPPORTUNITYSTAGEHISTORY_API_BASE_URL + '/' );
    }

    createOpportunityStageHistory(opportunityStageHistory){
        return axios.post(OPPORTUNITYSTAGEHISTORY_API_BASE_URL  + '/create', opportunityStageHistory);
    }

    getOpportunityStageHistoryById(opportunityStageHistoryId){
        return axios.get(OPPORTUNITYSTAGEHISTORY_API_BASE_URL + '/load?opportunityStageHistoryId=' + opportunityStageHistoryId);
    }

    updateOpportunityStageHistory(opportunityStageHistory){
        return axios.put(OPPORTUNITYSTAGEHISTORY_API_BASE_URL + '/update', opportunityStageHistory);
    }

    deleteOpportunityStageHistory(opportunityStageHistoryId){
        return axios.delete(OPPORTUNITYSTAGEHISTORY_API_BASE_URL + '/delete?opportunityStageHistoryId=' + opportunityStageHistoryId);
    }
}

export default new OpportunityStageHistoryService()