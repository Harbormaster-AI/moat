import axios from 'axios';

const ADVERTISER_API_BASE_URL = "/Advertiser";

class AdvertiserService {

    getAdvertisers(){
        return axios.get(ADVERTISER_API_BASE_URL + '/' );
    }

    createAdvertiser(advertiser){
        return axios.post(ADVERTISER_API_BASE_URL  + '/create', advertiser);
    }

    getAdvertiserById(advertiserId){
        return axios.get(ADVERTISER_API_BASE_URL + '/load?advertiserId=' + advertiserId);
    }

    updateAdvertiser(advertiser){
        return axios.put(ADVERTISER_API_BASE_URL + '/update', advertiser);
    }

    deleteAdvertiser(advertiserId){
        return axios.delete(ADVERTISER_API_BASE_URL + '/delete?advertiserId=' + advertiserId);
    }
}

export default new AdvertiserService()