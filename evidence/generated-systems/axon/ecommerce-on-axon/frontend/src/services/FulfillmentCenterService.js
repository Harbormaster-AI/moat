import axios from 'axios';

const FULFILLMENTCENTER_API_BASE_URL = "/FulfillmentCenter";

class FulfillmentCenterService {

    getFulfillmentCenters(){
        return axios.get(FULFILLMENTCENTER_API_BASE_URL + '/' );
    }

    createFulfillmentCenter(fulfillmentCenter){
        return axios.post(FULFILLMENTCENTER_API_BASE_URL  + '/create', fulfillmentCenter);
    }

    getFulfillmentCenterById(fulfillmentCenterId){
        return axios.get(FULFILLMENTCENTER_API_BASE_URL + '/load?fulfillmentCenterId=' + fulfillmentCenterId);
    }

    updateFulfillmentCenter(fulfillmentCenter){
        return axios.put(FULFILLMENTCENTER_API_BASE_URL + '/update', fulfillmentCenter);
    }

    deleteFulfillmentCenter(fulfillmentCenterId){
        return axios.delete(FULFILLMENTCENTER_API_BASE_URL + '/delete?fulfillmentCenterId=' + fulfillmentCenterId);
    }
}

export default new FulfillmentCenterService()