import axios from 'axios';

const INBOUNDSHIPMENTLINE_API_BASE_URL = "/InboundShipmentLine";

class InboundShipmentLineService {

    getInboundShipmentLines(){
        return axios.get(INBOUNDSHIPMENTLINE_API_BASE_URL + '/' );
    }

    createInboundShipmentLine(inboundShipmentLine){
        return axios.post(INBOUNDSHIPMENTLINE_API_BASE_URL  + '/create', inboundShipmentLine);
    }

    getInboundShipmentLineById(inboundShipmentLineId){
        return axios.get(INBOUNDSHIPMENTLINE_API_BASE_URL + '/load?inboundShipmentLineId=' + inboundShipmentLineId);
    }

    updateInboundShipmentLine(inboundShipmentLine){
        return axios.put(INBOUNDSHIPMENTLINE_API_BASE_URL + '/update', inboundShipmentLine);
    }

    deleteInboundShipmentLine(inboundShipmentLineId){
        return axios.delete(INBOUNDSHIPMENTLINE_API_BASE_URL + '/delete?inboundShipmentLineId=' + inboundShipmentLineId);
    }
}

export default new InboundShipmentLineService()