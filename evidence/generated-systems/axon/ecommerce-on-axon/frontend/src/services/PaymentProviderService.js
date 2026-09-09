import axios from 'axios';

const PAYMENTPROVIDER_API_BASE_URL = "/PaymentProvider";

class PaymentProviderService {

    getPaymentProviders(){
        return axios.get(PAYMENTPROVIDER_API_BASE_URL + '/' );
    }

    createPaymentProvider(paymentProvider){
        return axios.post(PAYMENTPROVIDER_API_BASE_URL  + '/create', paymentProvider);
    }

    getPaymentProviderById(paymentProviderId){
        return axios.get(PAYMENTPROVIDER_API_BASE_URL + '/load?paymentProviderId=' + paymentProviderId);
    }

    updatePaymentProvider(paymentProvider){
        return axios.put(PAYMENTPROVIDER_API_BASE_URL + '/update', paymentProvider);
    }

    deletePaymentProvider(paymentProviderId){
        return axios.delete(PAYMENTPROVIDER_API_BASE_URL + '/delete?paymentProviderId=' + paymentProviderId);
    }
}

export default new PaymentProviderService()