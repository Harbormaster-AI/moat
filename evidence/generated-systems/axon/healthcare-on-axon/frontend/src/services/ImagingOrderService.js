import axios from 'axios';

const IMAGINGORDER_API_BASE_URL = "/ImagingOrder";

class ImagingOrderService {

    getImagingOrders(){
        return axios.get(IMAGINGORDER_API_BASE_URL + '/' );
    }

    createImagingOrder(imagingOrder){
        return axios.post(IMAGINGORDER_API_BASE_URL  + '/create', imagingOrder);
    }

    getImagingOrderById(imagingOrderId){
        return axios.get(IMAGINGORDER_API_BASE_URL + '/load?imagingOrderId=' + imagingOrderId);
    }

    updateImagingOrder(imagingOrder){
        return axios.put(IMAGINGORDER_API_BASE_URL + '/update', imagingOrder);
    }

    deleteImagingOrder(imagingOrderId){
        return axios.delete(IMAGINGORDER_API_BASE_URL + '/delete?imagingOrderId=' + imagingOrderId);
    }
}

export default new ImagingOrderService()