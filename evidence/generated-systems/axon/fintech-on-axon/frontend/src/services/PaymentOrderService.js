import axios from 'axios';

const PAYMENTORDER_API_BASE_URL = "/PaymentOrder";

class PaymentOrderService {

    getPaymentOrders(){
        return axios.get(PAYMENTORDER_API_BASE_URL + '/' );
    }

    createPaymentOrder(paymentOrder){
        return axios.post(PAYMENTORDER_API_BASE_URL  + '/create', paymentOrder);
    }

    getPaymentOrderById(paymentOrderId){
        return axios.get(PAYMENTORDER_API_BASE_URL + '/load?paymentOrderId=' + paymentOrderId);
    }

    updatePaymentOrder(paymentOrder){
        return axios.put(PAYMENTORDER_API_BASE_URL + '/update', paymentOrder);
    }

    deletePaymentOrder(paymentOrderId){
        return axios.delete(PAYMENTORDER_API_BASE_URL + '/delete?paymentOrderId=' + paymentOrderId);
    }
}

export default new PaymentOrderService()