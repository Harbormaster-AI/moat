import axios from 'axios';

const PAYMENT_API_BASE_URL = "/Payment";

class PaymentService {

    getPayments(){
        return axios.get(PAYMENT_API_BASE_URL + '/' );
    }

    createPayment(payment){
        return axios.post(PAYMENT_API_BASE_URL  + '/create', payment);
    }

    getPaymentById(paymentId){
        return axios.get(PAYMENT_API_BASE_URL + '/load?paymentId=' + paymentId);
    }

    updatePayment(payment){
        return axios.put(PAYMENT_API_BASE_URL + '/update', payment);
    }

    deletePayment(paymentId){
        return axios.delete(PAYMENT_API_BASE_URL + '/delete?paymentId=' + paymentId);
    }
}

export default new PaymentService()