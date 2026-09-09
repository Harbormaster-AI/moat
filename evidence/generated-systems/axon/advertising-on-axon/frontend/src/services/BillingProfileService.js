import axios from 'axios';

const BILLINGPROFILE_API_BASE_URL = "/BillingProfile";

class BillingProfileService {

    getBillingProfiles(){
        return axios.get(BILLINGPROFILE_API_BASE_URL + '/' );
    }

    createBillingProfile(billingProfile){
        return axios.post(BILLINGPROFILE_API_BASE_URL  + '/create', billingProfile);
    }

    getBillingProfileById(billingProfileId){
        return axios.get(BILLINGPROFILE_API_BASE_URL + '/load?billingProfileId=' + billingProfileId);
    }

    updateBillingProfile(billingProfile){
        return axios.put(BILLINGPROFILE_API_BASE_URL + '/update', billingProfile);
    }

    deleteBillingProfile(billingProfileId){
        return axios.delete(BILLINGPROFILE_API_BASE_URL + '/delete?billingProfileId=' + billingProfileId);
    }
}

export default new BillingProfileService()