import axios from 'axios';

const RECOMMENDATIONSCENARIO_API_BASE_URL = "/RecommendationScenario";

class RecommendationScenarioService {

    getRecommendationScenarios(){
        return axios.get(RECOMMENDATIONSCENARIO_API_BASE_URL + '/' );
    }

    createRecommendationScenario(recommendationScenario){
        return axios.post(RECOMMENDATIONSCENARIO_API_BASE_URL  + '/create', recommendationScenario);
    }

    getRecommendationScenarioById(recommendationScenarioId){
        return axios.get(RECOMMENDATIONSCENARIO_API_BASE_URL + '/load?recommendationScenarioId=' + recommendationScenarioId);
    }

    updateRecommendationScenario(recommendationScenario){
        return axios.put(RECOMMENDATIONSCENARIO_API_BASE_URL + '/update', recommendationScenario);
    }

    deleteRecommendationScenario(recommendationScenarioId){
        return axios.delete(RECOMMENDATIONSCENARIO_API_BASE_URL + '/delete?recommendationScenarioId=' + recommendationScenarioId);
    }
}

export default new RecommendationScenarioService()