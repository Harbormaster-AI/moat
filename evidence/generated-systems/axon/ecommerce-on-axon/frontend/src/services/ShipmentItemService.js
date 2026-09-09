import axios from 'axios';

const SHIPMENTITEM_API_BASE_URL = "/ShipmentItem";

class ShipmentItemService {

    getShipmentItems(){
        return axios.get(SHIPMENTITEM_API_BASE_URL + '/' );
    }

    createShipmentItem(shipmentItem){
        return axios.post(SHIPMENTITEM_API_BASE_URL  + '/create', shipmentItem);
    }

    getShipmentItemById(shipmentItemId){
        return axios.get(SHIPMENTITEM_API_BASE_URL + '/load?shipmentItemId=' + shipmentItemId);
    }

    updateShipmentItem(shipmentItem){
        return axios.put(SHIPMENTITEM_API_BASE_URL + '/update', shipmentItem);
    }

    deleteShipmentItem(shipmentItemId){
        return axios.delete(SHIPMENTITEM_API_BASE_URL + '/delete?shipmentItemId=' + shipmentItemId);
    }
}

export default new ShipmentItemService()