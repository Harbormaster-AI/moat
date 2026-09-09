import axios from 'axios';

const SEMANTICMODEL_API_BASE_URL = "/SemanticModel";

class SemanticModelService {

    getSemanticModels(){
        return axios.get(SEMANTICMODEL_API_BASE_URL + '/' );
    }

    createSemanticModel(semanticModel){
        return axios.post(SEMANTICMODEL_API_BASE_URL  + '/create', semanticModel);
    }

    getSemanticModelById(semanticModelId){
        return axios.get(SEMANTICMODEL_API_BASE_URL + '/load?semanticModelId=' + semanticModelId);
    }

    updateSemanticModel(semanticModel){
        return axios.put(SEMANTICMODEL_API_BASE_URL + '/update', semanticModel);
    }

    deleteSemanticModel(semanticModelId){
        return axios.delete(SEMANTICMODEL_API_BASE_URL + '/delete?semanticModelId=' + semanticModelId);
    }
}

export default new SemanticModelService()