import axios from 'axios';

const AGENCY_API_BASE_URL = "/Agency";

class AgencyService {

    getAgencys(){
        return axios.get(AGENCY_API_BASE_URL + '/' );
    }

    createAgency(agency){
        return axios.post(AGENCY_API_BASE_URL  + '/create', agency);
    }

    getAgencyById(agencyId){
        return axios.get(AGENCY_API_BASE_URL + '/load?agencyId=' + agencyId);
    }

    updateAgency(agency){
        return axios.put(AGENCY_API_BASE_URL + '/update', agency);
    }

    deleteAgency(agencyId){
        return axios.delete(AGENCY_API_BASE_URL + '/delete?agencyId=' + agencyId);
    }
}

export default new AgencyService()