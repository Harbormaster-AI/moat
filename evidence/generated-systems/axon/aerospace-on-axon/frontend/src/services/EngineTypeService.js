import axios from 'axios';

const ENGINETYPE_API_BASE_URL = "/EngineType";

class EngineTypeService {

    getEngineTypes(){
        return axios.get(ENGINETYPE_API_BASE_URL + '/' );
    }

    createEngineType(engineType){
        return axios.post(ENGINETYPE_API_BASE_URL  + '/create', engineType);
    }

    getEngineTypeById(engineTypeId){
        return axios.get(ENGINETYPE_API_BASE_URL + '/load?engineTypeId=' + engineTypeId);
    }

    updateEngineType(engineType){
        return axios.put(ENGINETYPE_API_BASE_URL + '/update', engineType);
    }

    deleteEngineType(engineTypeId){
        return axios.delete(ENGINETYPE_API_BASE_URL + '/delete?engineTypeId=' + engineTypeId);
    }
}

export default new EngineTypeService()