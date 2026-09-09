import axios from 'axios';

const AIRCRAFTOPTION_API_BASE_URL = "/AircraftOption";

class AircraftOptionService {

    getAircraftOptions(){
        return axios.get(AIRCRAFTOPTION_API_BASE_URL + '/' );
    }

    createAircraftOption(aircraftOption){
        return axios.post(AIRCRAFTOPTION_API_BASE_URL  + '/create', aircraftOption);
    }

    getAircraftOptionById(aircraftOptionId){
        return axios.get(AIRCRAFTOPTION_API_BASE_URL + '/load?aircraftOptionId=' + aircraftOptionId);
    }

    updateAircraftOption(aircraftOption){
        return axios.put(AIRCRAFTOPTION_API_BASE_URL + '/update', aircraftOption);
    }

    deleteAircraftOption(aircraftOptionId){
        return axios.delete(AIRCRAFTOPTION_API_BASE_URL + '/delete?aircraftOptionId=' + aircraftOptionId);
    }
}

export default new AircraftOptionService()