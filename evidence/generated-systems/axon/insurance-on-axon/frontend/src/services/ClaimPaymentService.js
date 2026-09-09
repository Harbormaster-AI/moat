import axios from 'axios';

const CLAIMPAYMENT_API_BASE_URL = "/ClaimPayment";

class ClaimPaymentService {

    getClaimPayments(){
        return axios.get(CLAIMPAYMENT_API_BASE_URL + '/' );
    }

    createClaimPayment(claimPayment){
        return axios.post(CLAIMPAYMENT_API_BASE_URL  + '/create', claimPayment);
    }

    getClaimPaymentById(claimPaymentId){
        return axios.get(CLAIMPAYMENT_API_BASE_URL + '/load?claimPaymentId=' + claimPaymentId);
    }

    updateClaimPayment(claimPayment){
        return axios.put(CLAIMPAYMENT_API_BASE_URL + '/update', claimPayment);
    }

    deleteClaimPayment(claimPaymentId){
        return axios.delete(CLAIMPAYMENT_API_BASE_URL + '/delete?claimPaymentId=' + claimPaymentId);
    }
}

export default new ClaimPaymentService()