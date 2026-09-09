import axios from 'axios';

const EQUITYGRANT_API_BASE_URL = "/EquityGrant";

class EquityGrantService {

    getEquityGrants(){
        return axios.get(EQUITYGRANT_API_BASE_URL + '/' );
    }

    createEquityGrant(equityGrant){
        return axios.post(EQUITYGRANT_API_BASE_URL  + '/create', equityGrant);
    }

    getEquityGrantById(equityGrantId){
        return axios.get(EQUITYGRANT_API_BASE_URL + '/load?equityGrantId=' + equityGrantId);
    }

    updateEquityGrant(equityGrant){
        return axios.put(EQUITYGRANT_API_BASE_URL + '/update', equityGrant);
    }

    deleteEquityGrant(equityGrantId){
        return axios.delete(EQUITYGRANT_API_BASE_URL + '/delete?equityGrantId=' + equityGrantId);
    }
}

export default new EquityGrantService()