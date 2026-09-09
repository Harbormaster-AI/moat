import axios from 'axios';

const CAMPAIGN_API_BASE_URL = "/Campaign";

class CampaignService {

    getCampaigns(){
        return axios.get(CAMPAIGN_API_BASE_URL + '/' );
    }

    createCampaign(campaign){
        return axios.post(CAMPAIGN_API_BASE_URL  + '/create', campaign);
    }

    getCampaignById(campaignId){
        return axios.get(CAMPAIGN_API_BASE_URL + '/load?campaignId=' + campaignId);
    }

    updateCampaign(campaign){
        return axios.put(CAMPAIGN_API_BASE_URL + '/update', campaign);
    }

    deleteCampaign(campaignId){
        return axios.delete(CAMPAIGN_API_BASE_URL + '/delete?campaignId=' + campaignId);
    }
}

export default new CampaignService()