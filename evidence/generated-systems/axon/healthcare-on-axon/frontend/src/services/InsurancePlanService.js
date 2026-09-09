import axios from 'axios';

const INSURANCEPLAN_API_BASE_URL = "/InsurancePlan";

class InsurancePlanService {

    getInsurancePlans(){
        return axios.get(INSURANCEPLAN_API_BASE_URL + '/' );
    }

    createInsurancePlan(insurancePlan){
        return axios.post(INSURANCEPLAN_API_BASE_URL  + '/create', insurancePlan);
    }

    getInsurancePlanById(insurancePlanId){
        return axios.get(INSURANCEPLAN_API_BASE_URL + '/load?insurancePlanId=' + insurancePlanId);
    }

    updateInsurancePlan(insurancePlan){
        return axios.put(INSURANCEPLAN_API_BASE_URL + '/update', insurancePlan);
    }

    deleteInsurancePlan(insurancePlanId){
        return axios.delete(INSURANCEPLAN_API_BASE_URL + '/delete?insurancePlanId=' + insurancePlanId);
    }
}

export default new InsurancePlanService()