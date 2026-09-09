import axios from 'axios';

const SALESORDERLINE_API_BASE_URL = "/SalesOrderLine";

class SalesOrderLineService {

    getSalesOrderLines(){
        return axios.get(SALESORDERLINE_API_BASE_URL + '/' );
    }

    createSalesOrderLine(salesOrderLine){
        return axios.post(SALESORDERLINE_API_BASE_URL  + '/create', salesOrderLine);
    }

    getSalesOrderLineById(salesOrderLineId){
        return axios.get(SALESORDERLINE_API_BASE_URL + '/load?salesOrderLineId=' + salesOrderLineId);
    }

    updateSalesOrderLine(salesOrderLine){
        return axios.put(SALESORDERLINE_API_BASE_URL + '/update', salesOrderLine);
    }

    deleteSalesOrderLine(salesOrderLineId){
        return axios.delete(SALESORDERLINE_API_BASE_URL + '/delete?salesOrderLineId=' + salesOrderLineId);
    }
}

export default new SalesOrderLineService()