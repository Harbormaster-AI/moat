import axios from 'axios';

const PRICINGPLAN_API_BASE_URL = "/PricingPlan";

class PricingPlanService {

    getPricingPlans(){
        return axios.get(PRICINGPLAN_API_BASE_URL + '/' );
    }

    createPricingPlan(pricingPlan){
        return axios.post(PRICINGPLAN_API_BASE_URL  + '/create', pricingPlan);
    }

    getPricingPlanById(pricingPlanId){
        return axios.get(PRICINGPLAN_API_BASE_URL + '/load?pricingPlanId=' + pricingPlanId);
    }

    updatePricingPlan(pricingPlan){
        return axios.put(PRICINGPLAN_API_BASE_URL + '/update', pricingPlan);
    }

    deletePricingPlan(pricingPlanId){
        return axios.delete(PRICINGPLAN_API_BASE_URL + '/delete?pricingPlanId=' + pricingPlanId);
    }
}

export default new PricingPlanService()