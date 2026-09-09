import axios from 'axios';

const COVERAGEDEFINITION_API_BASE_URL = "/CoverageDefinition";

class CoverageDefinitionService {

    getCoverageDefinitions(){
        return axios.get(COVERAGEDEFINITION_API_BASE_URL + '/' );
    }

    createCoverageDefinition(coverageDefinition){
        return axios.post(COVERAGEDEFINITION_API_BASE_URL  + '/create', coverageDefinition);
    }

    getCoverageDefinitionById(coverageDefinitionId){
        return axios.get(COVERAGEDEFINITION_API_BASE_URL + '/load?coverageDefinitionId=' + coverageDefinitionId);
    }

    updateCoverageDefinition(coverageDefinition){
        return axios.put(COVERAGEDEFINITION_API_BASE_URL + '/update', coverageDefinition);
    }

    deleteCoverageDefinition(coverageDefinitionId){
        return axios.delete(COVERAGEDEFINITION_API_BASE_URL + '/delete?coverageDefinitionId=' + coverageDefinitionId);
    }
}

export default new CoverageDefinitionService()