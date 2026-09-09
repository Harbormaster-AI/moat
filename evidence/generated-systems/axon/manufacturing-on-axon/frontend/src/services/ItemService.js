import axios from 'axios';

const ITEM_API_BASE_URL = "/Item";

class ItemService {

    getItems(){
        return axios.get(ITEM_API_BASE_URL + '/' );
    }

    createItem(item){
        return axios.post(ITEM_API_BASE_URL  + '/create', item);
    }

    getItemById(itemId){
        return axios.get(ITEM_API_BASE_URL + '/load?itemId=' + itemId);
    }

    updateItem(item){
        return axios.put(ITEM_API_BASE_URL + '/update', item);
    }

    deleteItem(itemId){
        return axios.delete(ITEM_API_BASE_URL + '/delete?itemId=' + itemId);
    }
}

export default new ItemService()