import axios from 'axios';

const BENEFITPLAN_API_BASE_URL = "/BenefitPlan";

class BenefitPlanService {

    getBenefitPlans(){
        return axios.get(BENEFITPLAN_API_BASE_URL + '/' );
    }

    createBenefitPlan(benefitPlan){
        return axios.post(BENEFITPLAN_API_BASE_URL  + '/create', benefitPlan);
    }

    getBenefitPlanById(benefitPlanId){
        return axios.get(BENEFITPLAN_API_BASE_URL + '/load?benefitPlanId=' + benefitPlanId);
    }

    updateBenefitPlan(benefitPlan){
        return axios.put(BENEFITPLAN_API_BASE_URL + '/update', benefitPlan);
    }

    deleteBenefitPlan(benefitPlanId){
        return axios.delete(BENEFITPLAN_API_BASE_URL + '/delete?benefitPlanId=' + benefitPlanId);
    }
}

export default new BenefitPlanService()