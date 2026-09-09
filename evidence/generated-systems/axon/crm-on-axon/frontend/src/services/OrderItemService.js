import axios from 'axios';

const ORDERITEM_API_BASE_URL = "/OrderItem";

class OrderItemService {

    getOrderItems(){
        return axios.get(ORDERITEM_API_BASE_URL + '/' );
    }

    createOrderItem(orderItem){
        return axios.post(ORDERITEM_API_BASE_URL  + '/create', orderItem);
    }

    getOrderItemById(orderItemId){
        return axios.get(ORDERITEM_API_BASE_URL + '/load?orderItemId=' + orderItemId);
    }

    updateOrderItem(orderItem){
        return axios.put(ORDERITEM_API_BASE_URL + '/update', orderItem);
    }

    deleteOrderItem(orderItemId){
        return axios.delete(ORDERITEM_API_BASE_URL + '/delete?orderItemId=' + orderItemId);
    }
}

export default new OrderItemService()