import axios from 'axios';

const CARRIERSERVICE_API_BASE_URL = "/CarrierService";

class CarrierServiceService {

    getCarrierServices(){
        return axios.get(CARRIERSERVICE_API_BASE_URL + '/' );
    }

    createCarrierService(carrierService){
        return axios.post(CARRIERSERVICE_API_BASE_URL  + '/create', carrierService);
    }

    getCarrierServiceById(carrierServiceId){
        return axios.get(CARRIERSERVICE_API_BASE_URL + '/load?carrierServiceId=' + carrierServiceId);
    }

    updateCarrierService(carrierService){
        return axios.put(CARRIERSERVICE_API_BASE_URL + '/update', carrierService);
    }

    deleteCarrierService(carrierServiceId){
        return axios.delete(CARRIERSERVICE_API_BASE_URL + '/delete?carrierServiceId=' + carrierServiceId);
    }
}

export default new CarrierServiceService()