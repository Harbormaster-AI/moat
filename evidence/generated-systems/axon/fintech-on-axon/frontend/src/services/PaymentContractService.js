import axios from 'axios';

const PAYMENTCONTRACT_API_BASE_URL = "/PaymentContract";

class PaymentContractService {

    getPaymentContracts(){
        return axios.get(PAYMENTCONTRACT_API_BASE_URL + '/' );
    }

    createPaymentContract(paymentContract){
        return axios.post(PAYMENTCONTRACT_API_BASE_URL  + '/create', paymentContract);
    }

    getPaymentContractById(paymentContractId){
        return axios.get(PAYMENTCONTRACT_API_BASE_URL + '/load?paymentContractId=' + paymentContractId);
    }

    updatePaymentContract(paymentContract){
        return axios.put(PAYMENTCONTRACT_API_BASE_URL + '/update', paymentContract);
    }

    deletePaymentContract(paymentContractId){
        return axios.delete(PAYMENTCONTRACT_API_BASE_URL + '/delete?paymentContractId=' + paymentContractId);
    }
}

export default new PaymentContractService()