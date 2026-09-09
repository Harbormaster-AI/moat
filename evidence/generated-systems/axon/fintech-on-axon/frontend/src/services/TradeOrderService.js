import axios from 'axios';

const TRADEORDER_API_BASE_URL = "/TradeOrder";

class TradeOrderService {

    getTradeOrders(){
        return axios.get(TRADEORDER_API_BASE_URL + '/' );
    }

    createTradeOrder(tradeOrder){
        return axios.post(TRADEORDER_API_BASE_URL  + '/create', tradeOrder);
    }

    getTradeOrderById(tradeOrderId){
        return axios.get(TRADEORDER_API_BASE_URL + '/load?tradeOrderId=' + tradeOrderId);
    }

    updateTradeOrder(tradeOrder){
        return axios.put(TRADEORDER_API_BASE_URL + '/update', tradeOrder);
    }

    deleteTradeOrder(tradeOrderId){
        return axios.delete(TRADEORDER_API_BASE_URL + '/delete?tradeOrderId=' + tradeOrderId);
    }
}

export default new TradeOrderService()