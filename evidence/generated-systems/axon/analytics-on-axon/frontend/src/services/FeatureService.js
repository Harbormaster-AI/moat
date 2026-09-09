import axios from 'axios';

const FEATURE_API_BASE_URL = "/Feature";

class FeatureService {

    getFeatures(){
        return axios.get(FEATURE_API_BASE_URL + '/' );
    }

    createFeature(feature){
        return axios.post(FEATURE_API_BASE_URL  + '/create', feature);
    }

    getFeatureById(featureId){
        return axios.get(FEATURE_API_BASE_URL + '/load?featureId=' + featureId);
    }

    updateFeature(feature){
        return axios.put(FEATURE_API_BASE_URL + '/update', feature);
    }

    deleteFeature(featureId){
        return axios.delete(FEATURE_API_BASE_URL + '/delete?featureId=' + featureId);
    }
}

export default new FeatureService()