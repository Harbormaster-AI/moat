import axios from 'axios';

const FEATURESET_API_BASE_URL = "/FeatureSet";

class FeatureSetService {

    getFeatureSets(){
        return axios.get(FEATURESET_API_BASE_URL + '/' );
    }

    createFeatureSet(featureSet){
        return axios.post(FEATURESET_API_BASE_URL  + '/create', featureSet);
    }

    getFeatureSetById(featureSetId){
        return axios.get(FEATURESET_API_BASE_URL + '/load?featureSetId=' + featureSetId);
    }

    updateFeatureSet(featureSet){
        return axios.put(FEATURESET_API_BASE_URL + '/update', featureSet);
    }

    deleteFeatureSet(featureSetId){
        return axios.delete(FEATURESET_API_BASE_URL + '/delete?featureSetId=' + featureSetId);
    }
}

export default new FeatureSetService()