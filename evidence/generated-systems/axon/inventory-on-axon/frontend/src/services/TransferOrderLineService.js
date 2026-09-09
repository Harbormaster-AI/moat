import axios from 'axios';

const TRANSFERORDERLINE_API_BASE_URL = "/TransferOrderLine";

class TransferOrderLineService {

    getTransferOrderLines(){
        return axios.get(TRANSFERORDERLINE_API_BASE_URL + '/' );
    }

    createTransferOrderLine(transferOrderLine){
        return axios.post(TRANSFERORDERLINE_API_BASE_URL  + '/create', transferOrderLine);
    }

    getTransferOrderLineById(transferOrderLineId){
        return axios.get(TRANSFERORDERLINE_API_BASE_URL + '/load?transferOrderLineId=' + transferOrderLineId);
    }

    updateTransferOrderLine(transferOrderLine){
        return axios.put(TRANSFERORDERLINE_API_BASE_URL + '/update', transferOrderLine);
    }

    deleteTransferOrderLine(transferOrderLineId){
        return axios.delete(TRANSFERORDERLINE_API_BASE_URL + '/delete?transferOrderLineId=' + transferOrderLineId);
    }
}

export default new TransferOrderLineService()