import axios from 'axios';

const POLICYCOVERAGE_API_BASE_URL = "/PolicyCoverage";

class PolicyCoverageService {

    getPolicyCoverages(){
        return axios.get(POLICYCOVERAGE_API_BASE_URL + '/' );
    }

    createPolicyCoverage(policyCoverage){
        return axios.post(POLICYCOVERAGE_API_BASE_URL  + '/create', policyCoverage);
    }

    getPolicyCoverageById(policyCoverageId){
        return axios.get(POLICYCOVERAGE_API_BASE_URL + '/load?policyCoverageId=' + policyCoverageId);
    }

    updatePolicyCoverage(policyCoverage){
        return axios.put(POLICYCOVERAGE_API_BASE_URL + '/update', policyCoverage);
    }

    deletePolicyCoverage(policyCoverageId){
        return axios.delete(POLICYCOVERAGE_API_BASE_URL + '/delete?policyCoverageId=' + policyCoverageId);
    }
}

export default new PolicyCoverageService()