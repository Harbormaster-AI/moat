import axios from 'axios';

const PROCEDUREORDER_API_BASE_URL = "/ProcedureOrder";

class ProcedureOrderService {

    getProcedureOrders(){
        return axios.get(PROCEDUREORDER_API_BASE_URL + '/' );
    }

    createProcedureOrder(procedureOrder){
        return axios.post(PROCEDUREORDER_API_BASE_URL  + '/create', procedureOrder);
    }

    getProcedureOrderById(procedureOrderId){
        return axios.get(PROCEDUREORDER_API_BASE_URL + '/load?procedureOrderId=' + procedureOrderId);
    }

    updateProcedureOrder(procedureOrder){
        return axios.put(PROCEDUREORDER_API_BASE_URL + '/update', procedureOrder);
    }

    deleteProcedureOrder(procedureOrderId){
        return axios.delete(PROCEDUREORDER_API_BASE_URL + '/delete?procedureOrderId=' + procedureOrderId);
    }
}

export default new ProcedureOrderService()