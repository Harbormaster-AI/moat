import axios from 'axios';

const INVENTORYSOURCE_API_BASE_URL = "/InventorySource";

class InventorySourceService {

    getInventorySources(){
        return axios.get(INVENTORYSOURCE_API_BASE_URL + '/' );
    }

    createInventorySource(inventorySource){
        return axios.post(INVENTORYSOURCE_API_BASE_URL  + '/create', inventorySource);
    }

    getInventorySourceById(inventorySourceId){
        return axios.get(INVENTORYSOURCE_API_BASE_URL + '/load?inventorySourceId=' + inventorySourceId);
    }

    updateInventorySource(inventorySource){
        return axios.put(INVENTORYSOURCE_API_BASE_URL + '/update', inventorySource);
    }

    deleteInventorySource(inventorySourceId){
        return axios.delete(INVENTORYSOURCE_API_BASE_URL + '/delete?inventorySourceId=' + inventorySourceId);
    }
}

export default new InventorySourceService()