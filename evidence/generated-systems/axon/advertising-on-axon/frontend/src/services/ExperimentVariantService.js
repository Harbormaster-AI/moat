import axios from 'axios';

const EXPERIMENTVARIANT_API_BASE_URL = "/ExperimentVariant";

class ExperimentVariantService {

    getExperimentVariants(){
        return axios.get(EXPERIMENTVARIANT_API_BASE_URL + '/' );
    }

    createExperimentVariant(experimentVariant){
        return axios.post(EXPERIMENTVARIANT_API_BASE_URL  + '/create', experimentVariant);
    }

    getExperimentVariantById(experimentVariantId){
        return axios.get(EXPERIMENTVARIANT_API_BASE_URL + '/load?experimentVariantId=' + experimentVariantId);
    }

    updateExperimentVariant(experimentVariant){
        return axios.put(EXPERIMENTVARIANT_API_BASE_URL + '/update', experimentVariant);
    }

    deleteExperimentVariant(experimentVariantId){
        return axios.delete(EXPERIMENTVARIANT_API_BASE_URL + '/delete?experimentVariantId=' + experimentVariantId);
    }
}

export default new ExperimentVariantService()