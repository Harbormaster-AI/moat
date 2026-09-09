import axios from 'axios';

const STOCKADJUSTMENTLINE_API_BASE_URL = "/StockAdjustmentLine";

class StockAdjustmentLineService {

    getStockAdjustmentLines(){
        return axios.get(STOCKADJUSTMENTLINE_API_BASE_URL + '/' );
    }

    createStockAdjustmentLine(stockAdjustmentLine){
        return axios.post(STOCKADJUSTMENTLINE_API_BASE_URL  + '/create', stockAdjustmentLine);
    }

    getStockAdjustmentLineById(stockAdjustmentLineId){
        return axios.get(STOCKADJUSTMENTLINE_API_BASE_URL + '/load?stockAdjustmentLineId=' + stockAdjustmentLineId);
    }

    updateStockAdjustmentLine(stockAdjustmentLine){
        return axios.put(STOCKADJUSTMENTLINE_API_BASE_URL + '/update', stockAdjustmentLine);
    }

    deleteStockAdjustmentLine(stockAdjustmentLineId){
        return axios.delete(STOCKADJUSTMENTLINE_API_BASE_URL + '/delete?stockAdjustmentLineId=' + stockAdjustmentLineId);
    }
}

export default new StockAdjustmentLineService()