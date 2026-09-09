import axios from 'axios';

const BOMITEM_API_BASE_URL = "/BOMItem";

class BOMItemService {

    getBOMItems(){
        return axios.get(BOMITEM_API_BASE_URL + '/' );
    }

    createBOMItem(bOMItem){
        return axios.post(BOMITEM_API_BASE_URL  + '/create', bOMItem);
    }

    getBOMItemById(bOMItemId){
        return axios.get(BOMITEM_API_BASE_URL + '/load?bOMItemId=' + bOMItemId);
    }

    updateBOMItem(bOMItem){
        return axios.put(BOMITEM_API_BASE_URL + '/update', bOMItem);
    }

    deleteBOMItem(bOMItemId){
        return axios.delete(BOMITEM_API_BASE_URL + '/delete?bOMItemId=' + bOMItemId);
    }
}

export default new BOMItemService()