import axios from 'axios';

const INBOUNDSHIPMENT_API_BASE_URL = "/InboundShipment";

class InboundShipmentService {

    getInboundShipments(){
        return axios.get(INBOUNDSHIPMENT_API_BASE_URL + '/' );
    }

    createInboundShipment(inboundShipment){
        return axios.post(INBOUNDSHIPMENT_API_BASE_URL  + '/create', inboundShipment);
    }

    getInboundShipmentById(inboundShipmentId){
        return axios.get(INBOUNDSHIPMENT_API_BASE_URL + '/load?inboundShipmentId=' + inboundShipmentId);
    }

    updateInboundShipment(inboundShipment){
        return axios.put(INBOUNDSHIPMENT_API_BASE_URL + '/update', inboundShipment);
    }

    deleteInboundShipment(inboundShipmentId){
        return axios.delete(INBOUNDSHIPMENT_API_BASE_URL + '/delete?inboundShipmentId=' + inboundShipmentId);
    }
}

export default new InboundShipmentService()