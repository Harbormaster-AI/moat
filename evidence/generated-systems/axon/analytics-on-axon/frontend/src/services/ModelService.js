import axios from 'axios';

const MODEL_API_BASE_URL = "/Model";

class ModelService {

    getModels(){
        return axios.get(MODEL_API_BASE_URL + '/' );
    }

    createModel(model){
        return axios.post(MODEL_API_BASE_URL  + '/create', model);
    }

    getModelById(modelId){
        return axios.get(MODEL_API_BASE_URL + '/load?modelId=' + modelId);
    }

    updateModel(model){
        return axios.put(MODEL_API_BASE_URL + '/update', model);
    }

    deleteModel(modelId){
        return axios.delete(MODEL_API_BASE_URL + '/delete?modelId=' + modelId);
    }
}

export default new ModelService()