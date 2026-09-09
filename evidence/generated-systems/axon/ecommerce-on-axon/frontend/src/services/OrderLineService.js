import axios from 'axios';

const ORDERLINE_API_BASE_URL = "/OrderLine";

class OrderLineService {

    getOrderLines(){
        return axios.get(ORDERLINE_API_BASE_URL + '/' );
    }

    createOrderLine(orderLine){
        return axios.post(ORDERLINE_API_BASE_URL  + '/create', orderLine);
    }

    getOrderLineById(orderLineId){
        return axios.get(ORDERLINE_API_BASE_URL + '/load?orderLineId=' + orderLineId);
    }

    updateOrderLine(orderLine){
        return axios.put(ORDERLINE_API_BASE_URL + '/update', orderLine);
    }

    deleteOrderLine(orderLineId){
        return axios.delete(ORDERLINE_API_BASE_URL + '/delete?orderLineId=' + orderLineId);
    }
}

export default new OrderLineService()