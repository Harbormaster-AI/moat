import axios from 'axios';

const PROMOTION_API_BASE_URL = "/Promotion";

class PromotionService {

    getPromotions(){
        return axios.get(PROMOTION_API_BASE_URL + '/' );
    }

    createPromotion(promotion){
        return axios.post(PROMOTION_API_BASE_URL  + '/create', promotion);
    }

    getPromotionById(promotionId){
        return axios.get(PROMOTION_API_BASE_URL + '/load?promotionId=' + promotionId);
    }

    updatePromotion(promotion){
        return axios.put(PROMOTION_API_BASE_URL + '/update', promotion);
    }

    deletePromotion(promotionId){
        return axios.delete(PROMOTION_API_BASE_URL + '/delete?promotionId=' + promotionId);
    }
}

export default new PromotionService()