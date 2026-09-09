import axios from 'axios';

const PRODUCTIONSCHEDULE_API_BASE_URL = "/ProductionSchedule";

class ProductionScheduleService {

    getProductionSchedules(){
        return axios.get(PRODUCTIONSCHEDULE_API_BASE_URL + '/' );
    }

    createProductionSchedule(productionSchedule){
        return axios.post(PRODUCTIONSCHEDULE_API_BASE_URL  + '/create', productionSchedule);
    }

    getProductionScheduleById(productionScheduleId){
        return axios.get(PRODUCTIONSCHEDULE_API_BASE_URL + '/load?productionScheduleId=' + productionScheduleId);
    }

    updateProductionSchedule(productionSchedule){
        return axios.put(PRODUCTIONSCHEDULE_API_BASE_URL + '/update', productionSchedule);
    }

    deleteProductionSchedule(productionScheduleId){
        return axios.delete(PRODUCTIONSCHEDULE_API_BASE_URL + '/delete?productionScheduleId=' + productionScheduleId);
    }
}

export default new ProductionScheduleService()