import axios from 'axios';

const CAREPLAN_API_BASE_URL = "/CarePlan";

class CarePlanService {

    getCarePlans(){
        return axios.get(CAREPLAN_API_BASE_URL + '/' );
    }

    createCarePlan(carePlan){
        return axios.post(CAREPLAN_API_BASE_URL  + '/create', carePlan);
    }

    getCarePlanById(carePlanId){
        return axios.get(CAREPLAN_API_BASE_URL + '/load?carePlanId=' + carePlanId);
    }

    updateCarePlan(carePlan){
        return axios.put(CAREPLAN_API_BASE_URL + '/update', carePlan);
    }

    deleteCarePlan(carePlanId){
        return axios.delete(CAREPLAN_API_BASE_URL + '/delete?carePlanId=' + carePlanId);
    }
}

export default new CarePlanService()