import axios from 'axios';

const CREATIVEVARIATION_API_BASE_URL = "/CreativeVariation";

class CreativeVariationService {

    getCreativeVariations(){
        return axios.get(CREATIVEVARIATION_API_BASE_URL + '/' );
    }

    createCreativeVariation(creativeVariation){
        return axios.post(CREATIVEVARIATION_API_BASE_URL  + '/create', creativeVariation);
    }

    getCreativeVariationById(creativeVariationId){
        return axios.get(CREATIVEVARIATION_API_BASE_URL + '/load?creativeVariationId=' + creativeVariationId);
    }

    updateCreativeVariation(creativeVariation){
        return axios.put(CREATIVEVARIATION_API_BASE_URL + '/update', creativeVariation);
    }

    deleteCreativeVariation(creativeVariationId){
        return axios.delete(CREATIVEVARIATION_API_BASE_URL + '/delete?creativeVariationId=' + creativeVariationId);
    }
}

export default new CreativeVariationService()