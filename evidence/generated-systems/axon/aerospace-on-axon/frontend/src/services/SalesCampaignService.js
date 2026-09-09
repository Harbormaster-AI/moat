import axios from 'axios';

const SALESCAMPAIGN_API_BASE_URL = "/SalesCampaign";

class SalesCampaignService {

    getSalesCampaigns(){
        return axios.get(SALESCAMPAIGN_API_BASE_URL + '/' );
    }

    createSalesCampaign(salesCampaign){
        return axios.post(SALESCAMPAIGN_API_BASE_URL  + '/create', salesCampaign);
    }

    getSalesCampaignById(salesCampaignId){
        return axios.get(SALESCAMPAIGN_API_BASE_URL + '/load?salesCampaignId=' + salesCampaignId);
    }

    updateSalesCampaign(salesCampaign){
        return axios.put(SALESCAMPAIGN_API_BASE_URL + '/update', salesCampaign);
    }

    deleteSalesCampaign(salesCampaignId){
        return axios.delete(SALESCAMPAIGN_API_BASE_URL + '/delete?salesCampaignId=' + salesCampaignId);
    }
}

export default new SalesCampaignService()