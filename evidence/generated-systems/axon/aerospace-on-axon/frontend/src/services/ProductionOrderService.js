import axios from 'axios';

const PRODUCTIONORDER_API_BASE_URL = "/ProductionOrder";

class ProductionOrderService {

    getProductionOrders(){
        return axios.get(PRODUCTIONORDER_API_BASE_URL + '/' );
    }

    createProductionOrder(productionOrder){
        return axios.post(PRODUCTIONORDER_API_BASE_URL  + '/create', productionOrder);
    }

    getProductionOrderById(productionOrderId){
        return axios.get(PRODUCTIONORDER_API_BASE_URL + '/load?productionOrderId=' + productionOrderId);
    }

    updateProductionOrder(productionOrder){
        return axios.put(PRODUCTIONORDER_API_BASE_URL + '/update', productionOrder);
    }

    deleteProductionOrder(productionOrderId){
        return axios.delete(PRODUCTIONORDER_API_BASE_URL + '/delete?productionOrderId=' + productionOrderId);
    }
}

export default new ProductionOrderService()