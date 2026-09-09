import axios from 'axios';

const AIRCRAFTFAMILY_API_BASE_URL = "/AircraftFamily";

class AircraftFamilyService {

    getAircraftFamilys(){
        return axios.get(AIRCRAFTFAMILY_API_BASE_URL + '/' );
    }

    createAircraftFamily(aircraftFamily){
        return axios.post(AIRCRAFTFAMILY_API_BASE_URL  + '/create', aircraftFamily);
    }

    getAircraftFamilyById(aircraftFamilyId){
        return axios.get(AIRCRAFTFAMILY_API_BASE_URL + '/load?aircraftFamilyId=' + aircraftFamilyId);
    }

    updateAircraftFamily(aircraftFamily){
        return axios.put(AIRCRAFTFAMILY_API_BASE_URL + '/update', aircraftFamily);
    }

    deleteAircraftFamily(aircraftFamilyId){
        return axios.delete(AIRCRAFTFAMILY_API_BASE_URL + '/delete?aircraftFamilyId=' + aircraftFamilyId);
    }
}

export default new AircraftFamilyService()