import axios from 'axios';

const TRANSFERORDER_API_BASE_URL = "/TransferOrder";

class TransferOrderService {

    getTransferOrders(){
        return axios.get(TRANSFERORDER_API_BASE_URL + '/' );
    }

    createTransferOrder(transferOrder){
        return axios.post(TRANSFERORDER_API_BASE_URL  + '/create', transferOrder);
    }

    getTransferOrderById(transferOrderId){
        return axios.get(TRANSFERORDER_API_BASE_URL + '/load?transferOrderId=' + transferOrderId);
    }

    updateTransferOrder(transferOrder){
        return axios.put(TRANSFERORDER_API_BASE_URL + '/update', transferOrder);
    }

    deleteTransferOrder(transferOrderId){
        return axios.delete(TRANSFERORDER_API_BASE_URL + '/delete?transferOrderId=' + transferOrderId);
    }
}

export default new TransferOrderService()