import axios from 'axios';

const BILLINGACCOUNT_API_BASE_URL = "/BillingAccount";

class BillingAccountService {

    getBillingAccounts(){
        return axios.get(BILLINGACCOUNT_API_BASE_URL + '/' );
    }

    createBillingAccount(billingAccount){
        return axios.post(BILLINGACCOUNT_API_BASE_URL  + '/create', billingAccount);
    }

    getBillingAccountById(billingAccountId){
        return axios.get(BILLINGACCOUNT_API_BASE_URL + '/load?billingAccountId=' + billingAccountId);
    }

    updateBillingAccount(billingAccount){
        return axios.put(BILLINGACCOUNT_API_BASE_URL + '/update', billingAccount);
    }

    deleteBillingAccount(billingAccountId){
        return axios.delete(BILLINGACCOUNT_API_BASE_URL + '/delete?billingAccountId=' + billingAccountId);
    }
}

export default new BillingAccountService()