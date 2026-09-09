import axios from 'axios';

const EXPERIMENT_API_BASE_URL = "/Experiment";

class ExperimentService {

    getExperiments(){
        return axios.get(EXPERIMENT_API_BASE_URL + '/' );
    }

    createExperiment(experiment){
        return axios.post(EXPERIMENT_API_BASE_URL  + '/create', experiment);
    }

    getExperimentById(experimentId){
        return axios.get(EXPERIMENT_API_BASE_URL + '/load?experimentId=' + experimentId);
    }

    updateExperiment(experiment){
        return axios.put(EXPERIMENT_API_BASE_URL + '/update', experiment);
    }

    deleteExperiment(experimentId){
        return axios.delete(EXPERIMENT_API_BASE_URL + '/delete?experimentId=' + experimentId);
    }
}

export default new ExperimentService()