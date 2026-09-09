import axios from 'axios';

const COUPON_API_BASE_URL = "/Coupon";

class CouponService {

    getCoupons(){
        return axios.get(COUPON_API_BASE_URL + '/' );
    }

    createCoupon(coupon){
        return axios.post(COUPON_API_BASE_URL  + '/create', coupon);
    }

    getCouponById(couponId){
        return axios.get(COUPON_API_BASE_URL + '/load?couponId=' + couponId);
    }

    updateCoupon(coupon){
        return axios.put(COUPON_API_BASE_URL + '/update', coupon);
    }

    deleteCoupon(couponId){
        return axios.delete(COUPON_API_BASE_URL + '/delete?couponId=' + couponId);
    }
}

export default new CouponService()