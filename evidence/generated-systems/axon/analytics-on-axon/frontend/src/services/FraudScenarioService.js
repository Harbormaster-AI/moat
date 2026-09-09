import axios from 'axios';

const FRAUDSCENARIO_API_BASE_URL = "/FraudScenario";

class FraudScenarioService {

    getFraudScenarios(){
        return axios.get(FRAUDSCENARIO_API_BASE_URL + '/' );
    }

    createFraudScenario(fraudScenario){
        return axios.post(FRAUDSCENARIO_API_BASE_URL  + '/create', fraudScenario);
    }

    getFraudScenarioById(fraudScenarioId){
        return axios.get(FRAUDSCENARIO_API_BASE_URL + '/load?fraudScenarioId=' + fraudScenarioId);
    }

    updateFraudScenario(fraudScenario){
        return axios.put(FRAUDSCENARIO_API_BASE_URL + '/update', fraudScenario);
    }

    deleteFraudScenario(fraudScenarioId){
        return axios.delete(FRAUDSCENARIO_API_BASE_URL + '/delete?fraudScenarioId=' + fraudScenarioId);
    }
}

export default new FraudScenarioService()