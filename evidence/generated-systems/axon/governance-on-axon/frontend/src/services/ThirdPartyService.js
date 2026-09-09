import axios from 'axios';

const THIRDPARTY_API_BASE_URL = "/ThirdParty";

class ThirdPartyService {

    getThirdPartys(){
        return axios.get(THIRDPARTY_API_BASE_URL + '/' );
    }

    createThirdParty(thirdParty){
        return axios.post(THIRDPARTY_API_BASE_URL  + '/create', thirdParty);
    }

    getThirdPartyById(thirdPartyId){
        return axios.get(THIRDPARTY_API_BASE_URL + '/load?thirdPartyId=' + thirdPartyId);
    }

    updateThirdParty(thirdParty){
        return axios.put(THIRDPARTY_API_BASE_URL + '/update', thirdParty);
    }

    deleteThirdParty(thirdPartyId){
        return axios.delete(THIRDPARTY_API_BASE_URL + '/delete?thirdPartyId=' + thirdPartyId);
    }
}

export default new ThirdPartyService()