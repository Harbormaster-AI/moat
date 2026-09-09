import axios from 'axios';

const SHIPMENT_API_BASE_URL = "/Shipment";

class ShipmentService {

    getShipments(){
        return axios.get(SHIPMENT_API_BASE_URL + '/' );
    }

    createShipment(shipment){
        return axios.post(SHIPMENT_API_BASE_URL  + '/create', shipment);
    }

    getShipmentById(shipmentId){
        return axios.get(SHIPMENT_API_BASE_URL + '/load?shipmentId=' + shipmentId);
    }

    updateShipment(shipment){
        return axios.put(SHIPMENT_API_BASE_URL + '/update', shipment);
    }

    deleteShipment(shipmentId){
        return axios.delete(SHIPMENT_API_BASE_URL + '/delete?shipmentId=' + shipmentId);
    }
}

export default new ShipmentService()