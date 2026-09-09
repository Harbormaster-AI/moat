import axios from 'axios';

const INVENTORYTHRESHOLDALERT_API_BASE_URL = "/InventoryThresholdAlert";

class InventoryThresholdAlertService {

    getInventoryThresholdAlerts(){
        return axios.get(INVENTORYTHRESHOLDALERT_API_BASE_URL + '/' );
    }

    createInventoryThresholdAlert(inventoryThresholdAlert){
        return axios.post(INVENTORYTHRESHOLDALERT_API_BASE_URL  + '/create', inventoryThresholdAlert);
    }

    getInventoryThresholdAlertById(inventoryThresholdAlertId){
        return axios.get(INVENTORYTHRESHOLDALERT_API_BASE_URL + '/load?inventoryThresholdAlertId=' + inventoryThresholdAlertId);
    }

    updateInventoryThresholdAlert(inventoryThresholdAlert){
        return axios.put(INVENTORYTHRESHOLDALERT_API_BASE_URL + '/update', inventoryThresholdAlert);
    }

    deleteInventoryThresholdAlert(inventoryThresholdAlertId){
        return axios.delete(INVENTORYTHRESHOLDALERT_API_BASE_URL + '/delete?inventoryThresholdAlertId=' + inventoryThresholdAlertId);
    }
}

export default new InventoryThresholdAlertService()