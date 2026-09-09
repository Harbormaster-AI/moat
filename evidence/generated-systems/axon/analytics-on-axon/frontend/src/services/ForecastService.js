import axios from 'axios';

const FORECAST_API_BASE_URL = "/Forecast";

class ForecastService {

    getForecasts(){
        return axios.get(FORECAST_API_BASE_URL + '/' );
    }

    createForecast(forecast){
        return axios.post(FORECAST_API_BASE_URL  + '/create', forecast);
    }

    getForecastById(forecastId){
        return axios.get(FORECAST_API_BASE_URL + '/load?forecastId=' + forecastId);
    }

    updateForecast(forecast){
        return axios.put(FORECAST_API_BASE_URL + '/update', forecast);
    }

    deleteForecast(forecastId){
        return axios.delete(FORECAST_API_BASE_URL + '/delete?forecastId=' + forecastId);
    }
}

export default new ForecastService()