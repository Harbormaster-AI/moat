import axios from 'axios';

const ADACCOUNT_API_BASE_URL = "/AdAccount";

class AdAccountService {

    getAdAccounts(){
        return axios.get(ADACCOUNT_API_BASE_URL + '/' );
    }

    createAdAccount(adAccount){
        return axios.post(ADACCOUNT_API_BASE_URL  + '/create', adAccount);
    }

    getAdAccountById(adAccountId){
        return axios.get(ADACCOUNT_API_BASE_URL + '/load?adAccountId=' + adAccountId);
    }

    updateAdAccount(adAccount){
        return axios.put(ADACCOUNT_API_BASE_URL + '/update', adAccount);
    }

    deleteAdAccount(adAccountId){
        return axios.delete(ADACCOUNT_API_BASE_URL + '/delete?adAccountId=' + adAccountId);
    }
}

export default new AdAccountService()