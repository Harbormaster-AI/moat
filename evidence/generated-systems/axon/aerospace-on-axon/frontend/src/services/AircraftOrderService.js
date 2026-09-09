import axios from 'axios';

const AIRCRAFTORDER_API_BASE_URL = "/AircraftOrder";

class AircraftOrderService {

    getAircraftOrders(){
        return axios.get(AIRCRAFTORDER_API_BASE_URL + '/' );
    }

    createAircraftOrder(aircraftOrder){
        return axios.post(AIRCRAFTORDER_API_BASE_URL  + '/create', aircraftOrder);
    }

    getAircraftOrderById(aircraftOrderId){
        return axios.get(AIRCRAFTORDER_API_BASE_URL + '/load?aircraftOrderId=' + aircraftOrderId);
    }

    updateAircraftOrder(aircraftOrder){
        return axios.put(AIRCRAFTORDER_API_BASE_URL + '/update', aircraftOrder);
    }

    deleteAircraftOrder(aircraftOrderId){
        return axios.delete(AIRCRAFTORDER_API_BASE_URL + '/delete?aircraftOrderId=' + aircraftOrderId);
    }
}

export default new AircraftOrderService()