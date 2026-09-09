import axios from 'axios';

const SALESORDER_API_BASE_URL = "/SalesOrder";

class SalesOrderService {

    getSalesOrders(){
        return axios.get(SALESORDER_API_BASE_URL + '/' );
    }

    createSalesOrder(salesOrder){
        return axios.post(SALESORDER_API_BASE_URL  + '/create', salesOrder);
    }

    getSalesOrderById(salesOrderId){
        return axios.get(SALESORDER_API_BASE_URL + '/load?salesOrderId=' + salesOrderId);
    }

    updateSalesOrder(salesOrder){
        return axios.put(SALESORDER_API_BASE_URL + '/update', salesOrder);
    }

    deleteSalesOrder(salesOrderId){
        return axios.delete(SALESORDER_API_BASE_URL + '/delete?salesOrderId=' + salesOrderId);
    }
}

export default new SalesOrderService()