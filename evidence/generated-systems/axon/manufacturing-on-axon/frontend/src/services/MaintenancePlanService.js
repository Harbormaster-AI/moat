import axios from 'axios';

const MAINTENANCEPLAN_API_BASE_URL = "/MaintenancePlan";

class MaintenancePlanService {

    getMaintenancePlans(){
        return axios.get(MAINTENANCEPLAN_API_BASE_URL + '/' );
    }

    createMaintenancePlan(maintenancePlan){
        return axios.post(MAINTENANCEPLAN_API_BASE_URL  + '/create', maintenancePlan);
    }

    getMaintenancePlanById(maintenancePlanId){
        return axios.get(MAINTENANCEPLAN_API_BASE_URL + '/load?maintenancePlanId=' + maintenancePlanId);
    }

    updateMaintenancePlan(maintenancePlan){
        return axios.put(MAINTENANCEPLAN_API_BASE_URL + '/update', maintenancePlan);
    }

    deleteMaintenancePlan(maintenancePlanId){
        return axios.delete(MAINTENANCEPLAN_API_BASE_URL + '/delete?maintenancePlanId=' + maintenancePlanId);
    }
}

export default new MaintenancePlanService()