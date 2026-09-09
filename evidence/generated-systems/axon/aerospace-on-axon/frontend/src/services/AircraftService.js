import axios from 'axios';

const AIRCRAFT_API_BASE_URL = "/Aircraft";

class AircraftService {

    getAircrafts(){
        return axios.get(AIRCRAFT_API_BASE_URL + '/' );
    }

    createAircraft(aircraft){
        return axios.post(AIRCRAFT_API_BASE_URL  + '/create', aircraft);
    }

    getAircraftById(aircraftId){
        return axios.get(AIRCRAFT_API_BASE_URL + '/load?aircraftId=' + aircraftId);
    }

    updateAircraft(aircraft){
        return axios.put(AIRCRAFT_API_BASE_URL + '/update', aircraft);
    }

    deleteAircraft(aircraftId){
        return axios.delete(AIRCRAFT_API_BASE_URL + '/delete?aircraftId=' + aircraftId);
    }
}

export default new AircraftService()