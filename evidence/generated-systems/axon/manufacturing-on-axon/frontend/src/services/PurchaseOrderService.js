import axios from 'axios';

const PURCHASEORDER_API_BASE_URL = "/PurchaseOrder";

class PurchaseOrderService {

    getPurchaseOrders(){
        return axios.get(PURCHASEORDER_API_BASE_URL + '/' );
    }

    createPurchaseOrder(purchaseOrder){
        return axios.post(PURCHASEORDER_API_BASE_URL  + '/create', purchaseOrder);
    }

    getPurchaseOrderById(purchaseOrderId){
        return axios.get(PURCHASEORDER_API_BASE_URL + '/load?purchaseOrderId=' + purchaseOrderId);
    }

    updatePurchaseOrder(purchaseOrder){
        return axios.put(PURCHASEORDER_API_BASE_URL + '/update', purchaseOrder);
    }

    deletePurchaseOrder(purchaseOrderId){
        return axios.delete(PURCHASEORDER_API_BASE_URL + '/delete?purchaseOrderId=' + purchaseOrderId);
    }
}

export default new PurchaseOrderService()