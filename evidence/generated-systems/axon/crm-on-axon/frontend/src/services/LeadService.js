import axios from 'axios';

const LEAD_API_BASE_URL = "/Lead";

class LeadService {

    getLeads(){
        return axios.get(LEAD_API_BASE_URL + '/' );
    }

    createLead(lead){
        return axios.post(LEAD_API_BASE_URL  + '/create', lead);
    }

    getLeadById(leadId){
        return axios.get(LEAD_API_BASE_URL + '/load?leadId=' + leadId);
    }

    updateLead(lead){
        return axios.put(LEAD_API_BASE_URL + '/update', lead);
    }

    deleteLead(leadId){
        return axios.delete(LEAD_API_BASE_URL + '/delete?leadId=' + leadId);
    }
}

export default new LeadService()