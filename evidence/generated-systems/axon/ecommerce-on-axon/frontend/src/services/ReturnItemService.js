import axios from 'axios';

const RETURNITEM_API_BASE_URL = "/ReturnItem";

class ReturnItemService {

    getReturnItems(){
        return axios.get(RETURNITEM_API_BASE_URL + '/' );
    }

    createReturnItem(returnItem){
        return axios.post(RETURNITEM_API_BASE_URL  + '/create', returnItem);
    }

    getReturnItemById(returnItemId){
        return axios.get(RETURNITEM_API_BASE_URL + '/load?returnItemId=' + returnItemId);
    }

    updateReturnItem(returnItem){
        return axios.put(RETURNITEM_API_BASE_URL + '/update', returnItem);
    }

    deleteReturnItem(returnItemId){
        return axios.delete(RETURNITEM_API_BASE_URL + '/delete?returnItemId=' + returnItemId);
    }
}

export default new ReturnItemService()