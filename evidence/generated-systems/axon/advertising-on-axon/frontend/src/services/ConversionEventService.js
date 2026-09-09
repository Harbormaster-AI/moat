import axios from 'axios';

const CONVERSIONEVENT_API_BASE_URL = "/ConversionEvent";

class ConversionEventService {

    getConversionEvents(){
        return axios.get(CONVERSIONEVENT_API_BASE_URL + '/' );
    }

    createConversionEvent(conversionEvent){
        return axios.post(CONVERSIONEVENT_API_BASE_URL  + '/create', conversionEvent);
    }

    getConversionEventById(conversionEventId){
        return axios.get(CONVERSIONEVENT_API_BASE_URL + '/load?conversionEventId=' + conversionEventId);
    }

    updateConversionEvent(conversionEvent){
        return axios.put(CONVERSIONEVENT_API_BASE_URL + '/update', conversionEvent);
    }

    deleteConversionEvent(conversionEventId){
        return axios.delete(CONVERSIONEVENT_API_BASE_URL + '/delete?conversionEventId=' + conversionEventId);
    }
}

export default new ConversionEventService()