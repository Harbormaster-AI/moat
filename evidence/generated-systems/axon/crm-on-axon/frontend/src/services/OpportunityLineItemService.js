import axios from 'axios';

const OPPORTUNITYLINEITEM_API_BASE_URL = "/OpportunityLineItem";

class OpportunityLineItemService {

    getOpportunityLineItems(){
        return axios.get(OPPORTUNITYLINEITEM_API_BASE_URL + '/' );
    }

    createOpportunityLineItem(opportunityLineItem){
        return axios.post(OPPORTUNITYLINEITEM_API_BASE_URL  + '/create', opportunityLineItem);
    }

    getOpportunityLineItemById(opportunityLineItemId){
        return axios.get(OPPORTUNITYLINEITEM_API_BASE_URL + '/load?opportunityLineItemId=' + opportunityLineItemId);
    }

    updateOpportunityLineItem(opportunityLineItem){
        return axios.put(OPPORTUNITYLINEITEM_API_BASE_URL + '/update', opportunityLineItem);
    }

    deleteOpportunityLineItem(opportunityLineItemId){
        return axios.delete(OPPORTUNITYLINEITEM_API_BASE_URL + '/delete?opportunityLineItemId=' + opportunityLineItemId);
    }
}

export default new OpportunityLineItemService()