import axios from 'axios';

const PLANNEDORDER_API_BASE_URL = "/PlannedOrder";

class PlannedOrderService {

    getPlannedOrders(){
        return axios.get(PLANNEDORDER_API_BASE_URL + '/' );
    }

    createPlannedOrder(plannedOrder){
        return axios.post(PLANNEDORDER_API_BASE_URL  + '/create', plannedOrder);
    }

    getPlannedOrderById(plannedOrderId){
        return axios.get(PLANNEDORDER_API_BASE_URL + '/load?plannedOrderId=' + plannedOrderId);
    }

    updatePlannedOrder(plannedOrder){
        return axios.put(PLANNEDORDER_API_BASE_URL + '/update', plannedOrder);
    }

    deletePlannedOrder(plannedOrderId){
        return axios.delete(PLANNEDORDER_API_BASE_URL + '/delete?plannedOrderId=' + plannedOrderId);
    }
}

export default new PlannedOrderService()