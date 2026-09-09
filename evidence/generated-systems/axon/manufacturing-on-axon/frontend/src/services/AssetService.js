import axios from 'axios';

const ASSET_API_BASE_URL = "/Asset";

class AssetService {

    getAssets(){
        return axios.get(ASSET_API_BASE_URL + '/' );
    }

    createAsset(asset){
        return axios.post(ASSET_API_BASE_URL  + '/create', asset);
    }

    getAssetById(assetId){
        return axios.get(ASSET_API_BASE_URL + '/load?assetId=' + assetId);
    }

    updateAsset(asset){
        return axios.put(ASSET_API_BASE_URL + '/update', asset);
    }

    deleteAsset(assetId){
        return axios.delete(ASSET_API_BASE_URL + '/delete?assetId=' + assetId);
    }
}

export default new AssetService()