import axios from 'axios';

const TRADE_API_BASE_URL = "/Trade";

class TradeService {

    getTrades(){
        return axios.get(TRADE_API_BASE_URL + '/' );
    }

    createTrade(trade){
        return axios.post(TRADE_API_BASE_URL  + '/create', trade);
    }

    getTradeById(tradeId){
        return axios.get(TRADE_API_BASE_URL + '/load?tradeId=' + tradeId);
    }

    updateTrade(trade){
        return axios.put(TRADE_API_BASE_URL + '/update', trade);
    }

    deleteTrade(tradeId){
        return axios.delete(TRADE_API_BASE_URL + '/delete?tradeId=' + tradeId);
    }
}

export default new TradeService()