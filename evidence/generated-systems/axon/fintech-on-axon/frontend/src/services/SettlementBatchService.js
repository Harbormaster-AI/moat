import axios from 'axios';

const SETTLEMENTBATCH_API_BASE_URL = "/SettlementBatch";

class SettlementBatchService {

    getSettlementBatchs(){
        return axios.get(SETTLEMENTBATCH_API_BASE_URL + '/' );
    }

    createSettlementBatch(settlementBatch){
        return axios.post(SETTLEMENTBATCH_API_BASE_URL  + '/create', settlementBatch);
    }

    getSettlementBatchById(settlementBatchId){
        return axios.get(SETTLEMENTBATCH_API_BASE_URL + '/load?settlementBatchId=' + settlementBatchId);
    }

    updateSettlementBatch(settlementBatch){
        return axios.put(SETTLEMENTBATCH_API_BASE_URL + '/update', settlementBatch);
    }

    deleteSettlementBatch(settlementBatchId){
        return axios.delete(SETTLEMENTBATCH_API_BASE_URL + '/delete?settlementBatchId=' + settlementBatchId);
    }
}

export default new SettlementBatchService()