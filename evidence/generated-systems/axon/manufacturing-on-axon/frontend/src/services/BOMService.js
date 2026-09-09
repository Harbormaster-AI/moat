import axios from 'axios';

const BOM_API_BASE_URL = "/BOM";

class BOMService {

    getBOMs(){
        return axios.get(BOM_API_BASE_URL + '/' );
    }

    createBOM(bOM){
        return axios.post(BOM_API_BASE_URL  + '/create', bOM);
    }

    getBOMById(bOMId){
        return axios.get(BOM_API_BASE_URL + '/load?bOMId=' + bOMId);
    }

    updateBOM(bOM){
        return axios.put(BOM_API_BASE_URL + '/update', bOM);
    }

    deleteBOM(bOMId){
        return axios.delete(BOM_API_BASE_URL + '/delete?bOMId=' + bOMId);
    }
}

export default new BOMService()