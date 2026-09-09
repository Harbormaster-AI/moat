import axios from 'axios';

const TRAININGRUN_API_BASE_URL = "/TrainingRun";

class TrainingRunService {

    getTrainingRuns(){
        return axios.get(TRAININGRUN_API_BASE_URL + '/' );
    }

    createTrainingRun(trainingRun){
        return axios.post(TRAININGRUN_API_BASE_URL  + '/create', trainingRun);
    }

    getTrainingRunById(trainingRunId){
        return axios.get(TRAININGRUN_API_BASE_URL + '/load?trainingRunId=' + trainingRunId);
    }

    updateTrainingRun(trainingRun){
        return axios.put(TRAININGRUN_API_BASE_URL + '/update', trainingRun);
    }

    deleteTrainingRun(trainingRunId){
        return axios.delete(TRAININGRUN_API_BASE_URL + '/delete?trainingRunId=' + trainingRunId);
    }
}

export default new TrainingRunService()