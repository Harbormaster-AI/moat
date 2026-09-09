import axios from 'axios';

const REFUND_API_BASE_URL = "/Refund";

class RefundService {

    getRefunds(){
        return axios.get(REFUND_API_BASE_URL + '/' );
    }

    createRefund(refund){
        return axios.post(REFUND_API_BASE_URL  + '/create', refund);
    }

    getRefundById(refundId){
        return axios.get(REFUND_API_BASE_URL + '/load?refundId=' + refundId);
    }

    updateRefund(refund){
        return axios.put(REFUND_API_BASE_URL + '/update', refund);
    }

    deleteRefund(refundId){
        return axios.delete(REFUND_API_BASE_URL + '/delete?refundId=' + refundId);
    }
}

export default new RefundService()