import axios from 'axios';

const SECURITY_API_BASE_URL = "/Security";

class SecurityService {

    getSecuritys(){
        return axios.get(SECURITY_API_BASE_URL + '/' );
    }

    createSecurity(security){
        return axios.post(SECURITY_API_BASE_URL  + '/create', security);
    }

    getSecurityById(securityId){
        return axios.get(SECURITY_API_BASE_URL + '/load?securityId=' + securityId);
    }

    updateSecurity(security){
        return axios.put(SECURITY_API_BASE_URL + '/update', security);
    }

    deleteSecurity(securityId){
        return axios.delete(SECURITY_API_BASE_URL + '/delete?securityId=' + securityId);
    }
}

export default new SecurityService()