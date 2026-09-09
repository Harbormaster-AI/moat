import axios from 'axios';

const ORDER_API_BASE_URL = "/Order";

class OrderService {

    getOrders(){
        return axios.get(ORDER_API_BASE_URL + '/' );
    }

    createOrder(order){
        return axios.post(ORDER_API_BASE_URL  + '/create', order);
    }

    getOrderById(orderId){
        return axios.get(ORDER_API_BASE_URL + '/load?orderId=' + orderId);
    }

    updateOrder(order){
        return axios.put(ORDER_API_BASE_URL + '/update', order);
    }

    deleteOrder(orderId){
        return axios.delete(ORDER_API_BASE_URL + '/delete?orderId=' + orderId);
    }
}

export default new OrderService()