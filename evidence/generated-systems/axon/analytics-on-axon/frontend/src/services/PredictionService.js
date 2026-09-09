import axios from 'axios';

const PREDICTION_API_BASE_URL = "/Prediction";

class PredictionService {

    getPredictions(){
        return axios.get(PREDICTION_API_BASE_URL + '/' );
    }

    createPrediction(prediction){
        return axios.post(PREDICTION_API_BASE_URL  + '/create', prediction);
    }

    getPredictionById(predictionId){
        return axios.get(PREDICTION_API_BASE_URL + '/load?predictionId=' + predictionId);
    }

    updatePrediction(prediction){
        return axios.put(PREDICTION_API_BASE_URL + '/update', prediction);
    }

    deletePrediction(predictionId){
        return axios.delete(PREDICTION_API_BASE_URL + '/delete?predictionId=' + predictionId);
    }
}

export default new PredictionService()