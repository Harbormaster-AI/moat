import axios from 'axios';

const STOCKADJUSTMENT_API_BASE_URL = "/StockAdjustment";

class StockAdjustmentService {

    getStockAdjustments(){
        return axios.get(STOCKADJUSTMENT_API_BASE_URL + '/' );
    }

    createStockAdjustment(stockAdjustment){
        return axios.post(STOCKADJUSTMENT_API_BASE_URL  + '/create', stockAdjustment);
    }

    getStockAdjustmentById(stockAdjustmentId){
        return axios.get(STOCKADJUSTMENT_API_BASE_URL + '/load?stockAdjustmentId=' + stockAdjustmentId);
    }

    updateStockAdjustment(stockAdjustment){
        return axios.put(STOCKADJUSTMENT_API_BASE_URL + '/update', stockAdjustment);
    }

    deleteStockAdjustment(stockAdjustmentId){
        return axios.delete(STOCKADJUSTMENT_API_BASE_URL + '/delete?stockAdjustmentId=' + stockAdjustmentId);
    }
}

export default new StockAdjustmentService()