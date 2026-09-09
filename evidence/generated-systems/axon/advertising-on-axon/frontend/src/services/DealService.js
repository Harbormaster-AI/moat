import axios from 'axios';

const DEAL_API_BASE_URL = "/Deal";

class DealService {

    getDeals(){
        return axios.get(DEAL_API_BASE_URL + '/' );
    }

    createDeal(deal){
        return axios.post(DEAL_API_BASE_URL  + '/create', deal);
    }

    getDealById(dealId){
        return axios.get(DEAL_API_BASE_URL + '/load?dealId=' + dealId);
    }

    updateDeal(deal){
        return axios.put(DEAL_API_BASE_URL + '/update', deal);
    }

    deleteDeal(dealId){
        return axios.delete(DEAL_API_BASE_URL + '/delete?dealId=' + dealId);
    }
}

export default new DealService()