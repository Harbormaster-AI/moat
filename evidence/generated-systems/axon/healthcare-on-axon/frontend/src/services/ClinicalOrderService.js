import axios from 'axios';

const CLINICALORDER_API_BASE_URL = "/ClinicalOrder";

class ClinicalOrderService {

    getClinicalOrders(){
        return axios.get(CLINICALORDER_API_BASE_URL + '/' );
    }

    createClinicalOrder(clinicalOrder){
        return axios.post(CLINICALORDER_API_BASE_URL  + '/create', clinicalOrder);
    }

    getClinicalOrderById(clinicalOrderId){
        return axios.get(CLINICALORDER_API_BASE_URL + '/load?clinicalOrderId=' + clinicalOrderId);
    }

    updateClinicalOrder(clinicalOrder){
        return axios.put(CLINICALORDER_API_BASE_URL + '/update', clinicalOrder);
    }

    deleteClinicalOrder(clinicalOrderId){
        return axios.delete(CLINICALORDER_API_BASE_URL + '/delete?clinicalOrderId=' + clinicalOrderId);
    }
}

export default new ClinicalOrderService()