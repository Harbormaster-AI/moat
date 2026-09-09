import axios from 'axios';

const PURCHASEORDERLINE_API_BASE_URL = "/PurchaseOrderLine";

class PurchaseOrderLineService {

    getPurchaseOrderLines(){
        return axios.get(PURCHASEORDERLINE_API_BASE_URL + '/' );
    }

    createPurchaseOrderLine(purchaseOrderLine){
        return axios.post(PURCHASEORDERLINE_API_BASE_URL  + '/create', purchaseOrderLine);
    }

    getPurchaseOrderLineById(purchaseOrderLineId){
        return axios.get(PURCHASEORDERLINE_API_BASE_URL + '/load?purchaseOrderLineId=' + purchaseOrderLineId);
    }

    updatePurchaseOrderLine(purchaseOrderLine){
        return axios.put(PURCHASEORDERLINE_API_BASE_URL + '/update', purchaseOrderLine);
    }

    deletePurchaseOrderLine(purchaseOrderLineId){
        return axios.delete(PURCHASEORDERLINE_API_BASE_URL + '/delete?purchaseOrderLineId=' + purchaseOrderLineId);
    }
}

export default new PurchaseOrderLineService()