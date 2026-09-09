import axios from 'axios';

const FXDEAL_API_BASE_URL = "/FXDeal";

class FXDealService {

    getFXDeals(){
        return axios.get(FXDEAL_API_BASE_URL + '/' );
    }

    createFXDeal(fXDeal){
        return axios.post(FXDEAL_API_BASE_URL  + '/create', fXDeal);
    }

    getFXDealById(fXDealId){
        return axios.get(FXDEAL_API_BASE_URL + '/load?fXDealId=' + fXDealId);
    }

    updateFXDeal(fXDeal){
        return axios.put(FXDEAL_API_BASE_URL + '/update', fXDeal);
    }

    deleteFXDeal(fXDealId){
        return axios.delete(FXDEAL_API_BASE_URL + '/delete?fXDealId=' + fXDealId);
    }
}

export default new FXDealService()