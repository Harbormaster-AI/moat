import axios from 'axios';

const MEDICATIONORDER_API_BASE_URL = "/MedicationOrder";

class MedicationOrderService {

    getMedicationOrders(){
        return axios.get(MEDICATIONORDER_API_BASE_URL + '/' );
    }

    createMedicationOrder(medicationOrder){
        return axios.post(MEDICATIONORDER_API_BASE_URL  + '/create', medicationOrder);
    }

    getMedicationOrderById(medicationOrderId){
        return axios.get(MEDICATIONORDER_API_BASE_URL + '/load?medicationOrderId=' + medicationOrderId);
    }

    updateMedicationOrder(medicationOrder){
        return axios.put(MEDICATIONORDER_API_BASE_URL + '/update', medicationOrder);
    }

    deleteMedicationOrder(medicationOrderId){
        return axios.delete(MEDICATIONORDER_API_BASE_URL + '/delete?medicationOrderId=' + medicationOrderId);
    }
}

export default new MedicationOrderService()