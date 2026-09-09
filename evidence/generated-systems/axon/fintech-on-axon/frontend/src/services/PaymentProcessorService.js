import axios from 'axios';

const PAYMENTPROCESSOR_API_BASE_URL = "/PaymentProcessor";

class PaymentProcessorService {

    getPaymentProcessors(){
        return axios.get(PAYMENTPROCESSOR_API_BASE_URL + '/' );
    }

    createPaymentProcessor(paymentProcessor){
        return axios.post(PAYMENTPROCESSOR_API_BASE_URL  + '/create', paymentProcessor);
    }

    getPaymentProcessorById(paymentProcessorId){
        return axios.get(PAYMENTPROCESSOR_API_BASE_URL + '/load?paymentProcessorId=' + paymentProcessorId);
    }

    updatePaymentProcessor(paymentProcessor){
        return axios.put(PAYMENTPROCESSOR_API_BASE_URL + '/update', paymentProcessor);
    }

    deletePaymentProcessor(paymentProcessorId){
        return axios.delete(PAYMENTPROCESSOR_API_BASE_URL + '/delete?paymentProcessorId=' + paymentProcessorId);
    }
}

export default new PaymentProcessorService()