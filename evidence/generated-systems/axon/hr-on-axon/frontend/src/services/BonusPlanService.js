import axios from 'axios';

const BONUSPLAN_API_BASE_URL = "/BonusPlan";

class BonusPlanService {

    getBonusPlans(){
        return axios.get(BONUSPLAN_API_BASE_URL + '/' );
    }

    createBonusPlan(bonusPlan){
        return axios.post(BONUSPLAN_API_BASE_URL  + '/create', bonusPlan);
    }

    getBonusPlanById(bonusPlanId){
        return axios.get(BONUSPLAN_API_BASE_URL + '/load?bonusPlanId=' + bonusPlanId);
    }

    updateBonusPlan(bonusPlan){
        return axios.put(BONUSPLAN_API_BASE_URL + '/update', bonusPlan);
    }

    deleteBonusPlan(bonusPlanId){
        return axios.delete(BONUSPLAN_API_BASE_URL + '/delete?bonusPlanId=' + bonusPlanId);
    }
}

export default new BonusPlanService()