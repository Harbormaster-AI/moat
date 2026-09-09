import axios from 'axios';

const PURCHASEAGREEMENT_API_BASE_URL = "/PurchaseAgreement";

class PurchaseAgreementService {

    getPurchaseAgreements(){
        return axios.get(PURCHASEAGREEMENT_API_BASE_URL + '/' );
    }

    createPurchaseAgreement(purchaseAgreement){
        return axios.post(PURCHASEAGREEMENT_API_BASE_URL  + '/create', purchaseAgreement);
    }

    getPurchaseAgreementById(purchaseAgreementId){
        return axios.get(PURCHASEAGREEMENT_API_BASE_URL + '/load?purchaseAgreementId=' + purchaseAgreementId);
    }

    updatePurchaseAgreement(purchaseAgreement){
        return axios.put(PURCHASEAGREEMENT_API_BASE_URL + '/update', purchaseAgreement);
    }

    deletePurchaseAgreement(purchaseAgreementId){
        return axios.delete(PURCHASEAGREEMENT_API_BASE_URL + '/delete?purchaseAgreementId=' + purchaseAgreementId);
    }
}

export default new PurchaseAgreementService()