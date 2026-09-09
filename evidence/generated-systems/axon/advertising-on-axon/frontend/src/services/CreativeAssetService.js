import axios from 'axios';

const CREATIVEASSET_API_BASE_URL = "/CreativeAsset";

class CreativeAssetService {

    getCreativeAssets(){
        return axios.get(CREATIVEASSET_API_BASE_URL + '/' );
    }

    createCreativeAsset(creativeAsset){
        return axios.post(CREATIVEASSET_API_BASE_URL  + '/create', creativeAsset);
    }

    getCreativeAssetById(creativeAssetId){
        return axios.get(CREATIVEASSET_API_BASE_URL + '/load?creativeAssetId=' + creativeAssetId);
    }

    updateCreativeAsset(creativeAsset){
        return axios.put(CREATIVEASSET_API_BASE_URL + '/update', creativeAsset);
    }

    deleteCreativeAsset(creativeAssetId){
        return axios.delete(CREATIVEASSET_API_BASE_URL + '/delete?creativeAssetId=' + creativeAssetId);
    }
}

export default new CreativeAssetService()