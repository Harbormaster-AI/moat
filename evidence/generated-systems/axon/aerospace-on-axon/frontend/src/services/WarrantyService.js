import axios from 'axios';

const WARRANTY_API_BASE_URL = "/Warranty";

class WarrantyService {

    getWarrantys(){
        return axios.get(WARRANTY_API_BASE_URL + '/' );
    }

    createWarranty(warranty){
        return axios.post(WARRANTY_API_BASE_URL  + '/create', warranty);
    }

    getWarrantyById(warrantyId){
        return axios.get(WARRANTY_API_BASE_URL + '/load?warrantyId=' + warrantyId);
    }

    updateWarranty(warranty){
        return axios.put(WARRANTY_API_BASE_URL + '/update', warranty);
    }

    deleteWarranty(warrantyId){
        return axios.delete(WARRANTY_API_BASE_URL + '/delete?warrantyId=' + warrantyId);
    }
}

export default new WarrantyService()