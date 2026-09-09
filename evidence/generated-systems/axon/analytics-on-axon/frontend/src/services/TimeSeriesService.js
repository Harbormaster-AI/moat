import axios from 'axios';

const TIMESERIES_API_BASE_URL = "/TimeSeries";

class TimeSeriesService {

    getTimeSeriess(){
        return axios.get(TIMESERIES_API_BASE_URL + '/' );
    }

    createTimeSeries(timeSeries){
        return axios.post(TIMESERIES_API_BASE_URL  + '/create', timeSeries);
    }

    getTimeSeriesById(timeSeriesId){
        return axios.get(TIMESERIES_API_BASE_URL + '/load?timeSeriesId=' + timeSeriesId);
    }

    updateTimeSeries(timeSeries){
        return axios.put(TIMESERIES_API_BASE_URL + '/update', timeSeries);
    }

    deleteTimeSeries(timeSeriesId){
        return axios.delete(TIMESERIES_API_BASE_URL + '/delete?timeSeriesId=' + timeSeriesId);
    }
}

export default new TimeSeriesService()