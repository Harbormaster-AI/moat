import axios from 'axios';

const WORKORDER_API_BASE_URL = "/WorkOrder";

class WorkOrderService {

    getWorkOrders(){
        return axios.get(WORKORDER_API_BASE_URL + '/' );
    }

    createWorkOrder(workOrder){
        return axios.post(WORKORDER_API_BASE_URL  + '/create', workOrder);
    }

    getWorkOrderById(workOrderId){
        return axios.get(WORKORDER_API_BASE_URL + '/load?workOrderId=' + workOrderId);
    }

    updateWorkOrder(workOrder){
        return axios.put(WORKORDER_API_BASE_URL + '/update', workOrder);
    }

    deleteWorkOrder(workOrderId){
        return axios.delete(WORKORDER_API_BASE_URL + '/delete?workOrderId=' + workOrderId);
    }
}

export default new WorkOrderService()