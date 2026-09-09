import axios from 'axios';

const FLIGHTHEALTHEVENT_API_BASE_URL = "/FlightHealthEvent";

class FlightHealthEventService {

    getFlightHealthEvents(){
        return axios.get(FLIGHTHEALTHEVENT_API_BASE_URL + '/' );
    }

    createFlightHealthEvent(flightHealthEvent){
        return axios.post(FLIGHTHEALTHEVENT_API_BASE_URL  + '/create', flightHealthEvent);
    }

    getFlightHealthEventById(flightHealthEventId){
        return axios.get(FLIGHTHEALTHEVENT_API_BASE_URL + '/load?flightHealthEventId=' + flightHealthEventId);
    }

    updateFlightHealthEvent(flightHealthEvent){
        return axios.put(FLIGHTHEALTHEVENT_API_BASE_URL + '/update', flightHealthEvent);
    }

    deleteFlightHealthEvent(flightHealthEventId){
        return axios.delete(FLIGHTHEALTHEVENT_API_BASE_URL + '/delete?flightHealthEventId=' + flightHealthEventId);
    }
}

export default new FlightHealthEventService()