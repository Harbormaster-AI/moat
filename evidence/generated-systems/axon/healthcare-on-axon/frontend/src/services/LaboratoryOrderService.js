import axios from 'axios';

const LABORATORYORDER_API_BASE_URL = "/LaboratoryOrder";

class LaboratoryOrderService {

    getLaboratoryOrders(){
        return axios.get(LABORATORYORDER_API_BASE_URL + '/' );
    }

    createLaboratoryOrder(laboratoryOrder){
        return axios.post(LABORATORYORDER_API_BASE_URL  + '/create', laboratoryOrder);
    }

    getLaboratoryOrderById(laboratoryOrderId){
        return axios.get(LABORATORYORDER_API_BASE_URL + '/load?laboratoryOrderId=' + laboratoryOrderId);
    }

    updateLaboratoryOrder(laboratoryOrder){
        return axios.put(LABORATORYORDER_API_BASE_URL + '/update', laboratoryOrder);
    }

    deleteLaboratoryOrder(laboratoryOrderId){
        return axios.delete(LABORATORYORDER_API_BASE_URL + '/delete?laboratoryOrderId=' + laboratoryOrderId);
    }
}

export default new LaboratoryOrderService()