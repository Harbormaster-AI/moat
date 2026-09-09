import axios from 'axios';

const PRODUCTIONLINE_API_BASE_URL = "/ProductionLine";

class ProductionLineService {

    getProductionLines(){
        return axios.get(PRODUCTIONLINE_API_BASE_URL + '/' );
    }

    createProductionLine(productionLine){
        return axios.post(PRODUCTIONLINE_API_BASE_URL  + '/create', productionLine);
    }

    getProductionLineById(productionLineId){
        return axios.get(PRODUCTIONLINE_API_BASE_URL + '/load?productionLineId=' + productionLineId);
    }

    updateProductionLine(productionLine){
        return axios.put(PRODUCTIONLINE_API_BASE_URL + '/update', productionLine);
    }

    deleteProductionLine(productionLineId){
        return axios.delete(PRODUCTIONLINE_API_BASE_URL + '/delete?productionLineId=' + productionLineId);
    }
}

export default new ProductionLineService()