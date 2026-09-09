import axios from 'axios';

const COUPONREDEMPTION_API_BASE_URL = "/CouponRedemption";

class CouponRedemptionService {

    getCouponRedemptions(){
        return axios.get(COUPONREDEMPTION_API_BASE_URL + '/' );
    }

    createCouponRedemption(couponRedemption){
        return axios.post(COUPONREDEMPTION_API_BASE_URL  + '/create', couponRedemption);
    }

    getCouponRedemptionById(couponRedemptionId){
        return axios.get(COUPONREDEMPTION_API_BASE_URL + '/load?couponRedemptionId=' + couponRedemptionId);
    }

    updateCouponRedemption(couponRedemption){
        return axios.put(COUPONREDEMPTION_API_BASE_URL + '/update', couponRedemption);
    }

    deleteCouponRedemption(couponRedemptionId){
        return axios.delete(COUPONREDEMPTION_API_BASE_URL + '/delete?couponRedemptionId=' + couponRedemptionId);
    }
}

export default new CouponRedemptionService()