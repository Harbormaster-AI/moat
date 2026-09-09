import axios from 'axios';

const AIRCRAFTPROGRAM_API_BASE_URL = "/AircraftProgram";

class AircraftProgramService {

    getAircraftPrograms(){
        return axios.get(AIRCRAFTPROGRAM_API_BASE_URL + '/' );
    }

    createAircraftProgram(aircraftProgram){
        return axios.post(AIRCRAFTPROGRAM_API_BASE_URL  + '/create', aircraftProgram);
    }

    getAircraftProgramById(aircraftProgramId){
        return axios.get(AIRCRAFTPROGRAM_API_BASE_URL + '/load?aircraftProgramId=' + aircraftProgramId);
    }

    updateAircraftProgram(aircraftProgram){
        return axios.put(AIRCRAFTPROGRAM_API_BASE_URL + '/update', aircraftProgram);
    }

    deleteAircraftProgram(aircraftProgramId){
        return axios.delete(AIRCRAFTPROGRAM_API_BASE_URL + '/delete?aircraftProgramId=' + aircraftProgramId);
    }
}

export default new AircraftProgramService()