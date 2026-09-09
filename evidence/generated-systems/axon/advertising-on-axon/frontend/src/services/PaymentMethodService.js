import axios from 'axios';

const PAYMENTMETHOD_API_BASE_URL = "/PaymentMethod";

class PaymentMethodService {

    getPaymentMethods(){
        return axios.get(PAYMENTMETHOD_API_BASE_URL + '/' );
    }

    createPaymentMethod(paymentMethod){
        return axios.post(PAYMENTMETHOD_API_BASE_URL  + '/create', paymentMethod);
    }

    getPaymentMethodById(paymentMethodId){
        return axios.get(PAYMENTMETHOD_API_BASE_URL + '/load?paymentMethodId=' + paymentMethodId);
    }

    updatePaymentMethod(paymentMethod){
        return axios.put(PAYMENTMETHOD_API_BASE_URL + '/update', paymentMethod);
    }

    deletePaymentMethod(paymentMethodId){
        return axios.delete(PAYMENTMETHOD_API_BASE_URL + '/delete?paymentMethodId=' + paymentMethodId);
    }
}

export default new PaymentMethodService()