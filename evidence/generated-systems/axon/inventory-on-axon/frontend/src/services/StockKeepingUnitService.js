import axios from 'axios';

const STOCKKEEPINGUNIT_API_BASE_URL = "/StockKeepingUnit";

class StockKeepingUnitService {

    getStockKeepingUnits(){
        return axios.get(STOCKKEEPINGUNIT_API_BASE_URL + '/' );
    }

    createStockKeepingUnit(stockKeepingUnit){
        return axios.post(STOCKKEEPINGUNIT_API_BASE_URL  + '/create', stockKeepingUnit);
    }

    getStockKeepingUnitById(stockKeepingUnitId){
        return axios.get(STOCKKEEPINGUNIT_API_BASE_URL + '/load?stockKeepingUnitId=' + stockKeepingUnitId);
    }

    updateStockKeepingUnit(stockKeepingUnit){
        return axios.put(STOCKKEEPINGUNIT_API_BASE_URL + '/update', stockKeepingUnit);
    }

    deleteStockKeepingUnit(stockKeepingUnitId){
        return axios.delete(STOCKKEEPINGUNIT_API_BASE_URL + '/delete?stockKeepingUnitId=' + stockKeepingUnitId);
    }
}

export default new StockKeepingUnitService()