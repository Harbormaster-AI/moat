import axios from 'axios';

const INVENTORYITEM_API_BASE_URL = "/InventoryItem";

class InventoryItemService {

    getInventoryItems(){
        return axios.get(INVENTORYITEM_API_BASE_URL + '/' );
    }

    createInventoryItem(inventoryItem){
        return axios.post(INVENTORYITEM_API_BASE_URL  + '/create', inventoryItem);
    }

    getInventoryItemById(inventoryItemId){
        return axios.get(INVENTORYITEM_API_BASE_URL + '/load?inventoryItemId=' + inventoryItemId);
    }

    updateInventoryItem(inventoryItem){
        return axios.put(INVENTORYITEM_API_BASE_URL + '/update', inventoryItem);
    }

    deleteInventoryItem(inventoryItemId){
        return axios.delete(INVENTORYITEM_API_BASE_URL + '/delete?inventoryItemId=' + inventoryItemId);
    }
}

export default new InventoryItemService()