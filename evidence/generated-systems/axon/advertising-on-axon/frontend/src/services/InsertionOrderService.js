import axios from 'axios';

const INSERTIONORDER_API_BASE_URL = "/InsertionOrder";

class InsertionOrderService {

    getInsertionOrders(){
        return axios.get(INSERTIONORDER_API_BASE_URL + '/' );
    }

    createInsertionOrder(insertionOrder){
        return axios.post(INSERTIONORDER_API_BASE_URL  + '/create', insertionOrder);
    }

    getInsertionOrderById(insertionOrderId){
        return axios.get(INSERTIONORDER_API_BASE_URL + '/load?insertionOrderId=' + insertionOrderId);
    }

    updateInsertionOrder(insertionOrder){
        return axios.put(INSERTIONORDER_API_BASE_URL + '/update', insertionOrder);
    }

    deleteInsertionOrder(insertionOrderId){
        return axios.delete(INSERTIONORDER_API_BASE_URL + '/delete?insertionOrderId=' + insertionOrderId);
    }
}

export default new InsertionOrderService()