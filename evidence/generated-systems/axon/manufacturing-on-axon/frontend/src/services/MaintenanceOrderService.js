import axios from 'axios';

const MAINTENANCEORDER_API_BASE_URL = "/MaintenanceOrder";

class MaintenanceOrderService {

    getMaintenanceOrders(){
        return axios.get(MAINTENANCEORDER_API_BASE_URL + '/' );
    }

    createMaintenanceOrder(maintenanceOrder){
        return axios.post(MAINTENANCEORDER_API_BASE_URL  + '/create', maintenanceOrder);
    }

    getMaintenanceOrderById(maintenanceOrderId){
        return axios.get(MAINTENANCEORDER_API_BASE_URL + '/load?maintenanceOrderId=' + maintenanceOrderId);
    }

    updateMaintenanceOrder(maintenanceOrder){
        return axios.put(MAINTENANCEORDER_API_BASE_URL + '/update', maintenanceOrder);
    }

    deleteMaintenanceOrder(maintenanceOrderId){
        return axios.delete(MAINTENANCEORDER_API_BASE_URL + '/delete?maintenanceOrderId=' + maintenanceOrderId);
    }
}

export default new MaintenanceOrderService()