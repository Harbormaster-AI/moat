import axios from 'axios';

const MODELVERSION_API_BASE_URL = "/ModelVersion";

class ModelVersionService {

    getModelVersions(){
        return axios.get(MODELVERSION_API_BASE_URL + '/' );
    }

    createModelVersion(modelVersion){
        return axios.post(MODELVERSION_API_BASE_URL  + '/create', modelVersion);
    }

    getModelVersionById(modelVersionId){
        return axios.get(MODELVERSION_API_BASE_URL + '/load?modelVersionId=' + modelVersionId);
    }

    updateModelVersion(modelVersion){
        return axios.put(MODELVERSION_API_BASE_URL + '/update', modelVersion);
    }

    deleteModelVersion(modelVersionId){
        return axios.delete(MODELVERSION_API_BASE_URL + '/delete?modelVersionId=' + modelVersionId);
    }
}

export default new ModelVersionService()