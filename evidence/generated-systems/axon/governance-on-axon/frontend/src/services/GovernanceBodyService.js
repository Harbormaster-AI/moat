import axios from 'axios';

const GOVERNANCEBODY_API_BASE_URL = "/GovernanceBody";

class GovernanceBodyService {

    getGovernanceBodys(){
        return axios.get(GOVERNANCEBODY_API_BASE_URL + '/' );
    }

    createGovernanceBody(governanceBody){
        return axios.post(GOVERNANCEBODY_API_BASE_URL  + '/create', governanceBody);
    }

    getGovernanceBodyById(governanceBodyId){
        return axios.get(GOVERNANCEBODY_API_BASE_URL + '/load?governanceBodyId=' + governanceBodyId);
    }

    updateGovernanceBody(governanceBody){
        return axios.put(GOVERNANCEBODY_API_BASE_URL + '/update', governanceBody);
    }

    deleteGovernanceBody(governanceBodyId){
        return axios.delete(GOVERNANCEBODY_API_BASE_URL + '/delete?governanceBodyId=' + governanceBodyId);
    }
}

export default new GovernanceBodyService()