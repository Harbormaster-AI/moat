import axios from 'axios';

const CAMPAIGNMEMBER_API_BASE_URL = "/CampaignMember";

class CampaignMemberService {

    getCampaignMembers(){
        return axios.get(CAMPAIGNMEMBER_API_BASE_URL + '/' );
    }

    createCampaignMember(campaignMember){
        return axios.post(CAMPAIGNMEMBER_API_BASE_URL  + '/create', campaignMember);
    }

    getCampaignMemberById(campaignMemberId){
        return axios.get(CAMPAIGNMEMBER_API_BASE_URL + '/load?campaignMemberId=' + campaignMemberId);
    }

    updateCampaignMember(campaignMember){
        return axios.put(CAMPAIGNMEMBER_API_BASE_URL + '/update', campaignMember);
    }

    deleteCampaignMember(campaignMemberId){
        return axios.delete(CAMPAIGNMEMBER_API_BASE_URL + '/delete?campaignMemberId=' + campaignMemberId);
    }
}

export default new CampaignMemberService()