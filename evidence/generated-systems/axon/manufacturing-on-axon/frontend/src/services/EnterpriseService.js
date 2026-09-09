import axios from 'axios';

const ENTERPRISE_API_BASE_URL = "/Enterprise";

class EnterpriseService {

    getEnterprises(){
        return axios.get(ENTERPRISE_API_BASE_URL + '/' );
    }

    createEnterprise(enterprise){
        return axios.post(ENTERPRISE_API_BASE_URL  + '/create', enterprise);
    }

    getEnterpriseById(enterpriseId){
        return axios.get(ENTERPRISE_API_BASE_URL + '/load?enterpriseId=' + enterpriseId);
    }

    updateEnterprise(enterprise){
        return axios.put(ENTERPRISE_API_BASE_URL + '/update', enterprise);
    }

    deleteEnterprise(enterpriseId){
        return axios.delete(ENTERPRISE_API_BASE_URL + '/delete?enterpriseId=' + enterpriseId);
    }
}

export default new EnterpriseService()