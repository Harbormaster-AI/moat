import axios from 'axios';

const REGULATION_API_BASE_URL = "/Regulation";

class RegulationService {

    getRegulations(){
        return axios.get(REGULATION_API_BASE_URL + '/' );
    }

    createRegulation(regulation){
        return axios.post(REGULATION_API_BASE_URL  + '/create', regulation);
    }

    getRegulationById(regulationId){
        return axios.get(REGULATION_API_BASE_URL + '/load?regulationId=' + regulationId);
    }

    updateRegulation(regulation){
        return axios.put(REGULATION_API_BASE_URL + '/update', regulation);
    }

    deleteRegulation(regulationId){
        return axios.delete(REGULATION_API_BASE_URL + '/delete?regulationId=' + regulationId);
    }
}

export default new RegulationService()