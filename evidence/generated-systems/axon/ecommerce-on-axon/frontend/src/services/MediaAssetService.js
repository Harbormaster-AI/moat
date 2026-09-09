import axios from 'axios';

const MEDIAASSET_API_BASE_URL = "/MediaAsset";

class MediaAssetService {

    getMediaAssets(){
        return axios.get(MEDIAASSET_API_BASE_URL + '/' );
    }

    createMediaAsset(mediaAsset){
        return axios.post(MEDIAASSET_API_BASE_URL  + '/create', mediaAsset);
    }

    getMediaAssetById(mediaAssetId){
        return axios.get(MEDIAASSET_API_BASE_URL + '/load?mediaAssetId=' + mediaAssetId);
    }

    updateMediaAsset(mediaAsset){
        return axios.put(MEDIAASSET_API_BASE_URL + '/update', mediaAsset);
    }

    deleteMediaAsset(mediaAssetId){
        return axios.delete(MEDIAASSET_API_BASE_URL + '/delete?mediaAssetId=' + mediaAssetId);
    }
}

export default new MediaAssetService()