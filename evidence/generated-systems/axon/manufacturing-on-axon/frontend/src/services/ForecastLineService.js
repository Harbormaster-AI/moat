import axios from 'axios';

const FORECASTLINE_API_BASE_URL = "/ForecastLine";

class ForecastLineService {

    getForecastLines(){
        return axios.get(FORECASTLINE_API_BASE_URL + '/' );
    }

    createForecastLine(forecastLine){
        return axios.post(FORECASTLINE_API_BASE_URL  + '/create', forecastLine);
    }

    getForecastLineById(forecastLineId){
        return axios.get(FORECASTLINE_API_BASE_URL + '/load?forecastLineId=' + forecastLineId);
    }

    updateForecastLine(forecastLine){
        return axios.put(FORECASTLINE_API_BASE_URL + '/update', forecastLine);
    }

    deleteForecastLine(forecastLineId){
        return axios.delete(FORECASTLINE_API_BASE_URL + '/delete?forecastLineId=' + forecastLineId);
    }
}

export default new ForecastLineService()