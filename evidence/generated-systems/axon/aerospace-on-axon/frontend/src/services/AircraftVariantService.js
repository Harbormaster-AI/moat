import axios from 'axios';

const AIRCRAFTVARIANT_API_BASE_URL = "/AircraftVariant";

class AircraftVariantService {

    getAircraftVariants(){
        return axios.get(AIRCRAFTVARIANT_API_BASE_URL + '/' );
    }

    createAircraftVariant(aircraftVariant){
        return axios.post(AIRCRAFTVARIANT_API_BASE_URL  + '/create', aircraftVariant);
    }

    getAircraftVariantById(aircraftVariantId){
        return axios.get(AIRCRAFTVARIANT_API_BASE_URL + '/load?aircraftVariantId=' + aircraftVariantId);
    }

    updateAircraftVariant(aircraftVariant){
        return axios.put(AIRCRAFTVARIANT_API_BASE_URL + '/update', aircraftVariant);
    }

    deleteAircraftVariant(aircraftVariantId){
        return axios.delete(AIRCRAFTVARIANT_API_BASE_URL + '/delete?aircraftVariantId=' + aircraftVariantId);
    }
}

export default new AircraftVariantService()