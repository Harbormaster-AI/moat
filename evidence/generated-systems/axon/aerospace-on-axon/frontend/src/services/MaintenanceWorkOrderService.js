import axios from 'axios';

const MAINTENANCEWORKORDER_API_BASE_URL = "/MaintenanceWorkOrder";

class MaintenanceWorkOrderService {

    getMaintenanceWorkOrders(){
        return axios.get(MAINTENANCEWORKORDER_API_BASE_URL + '/' );
    }

    createMaintenanceWorkOrder(maintenanceWorkOrder){
        return axios.post(MAINTENANCEWORKORDER_API_BASE_URL  + '/create', maintenanceWorkOrder);
    }

    getMaintenanceWorkOrderById(maintenanceWorkOrderId){
        return axios.get(MAINTENANCEWORKORDER_API_BASE_URL + '/load?maintenanceWorkOrderId=' + maintenanceWorkOrderId);
    }

    updateMaintenanceWorkOrder(maintenanceWorkOrder){
        return axios.put(MAINTENANCEWORKORDER_API_BASE_URL + '/update', maintenanceWorkOrder);
    }

    deleteMaintenanceWorkOrder(maintenanceWorkOrderId){
        return axios.delete(MAINTENANCEWORKORDER_API_BASE_URL + '/delete?maintenanceWorkOrderId=' + maintenanceWorkOrderId);
    }
}

export default new MaintenanceWorkOrderService()