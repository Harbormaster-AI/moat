import axios from 'axios';

const LANDINGGEAR_API_BASE_URL = "/LandingGear";

class LandingGearService {

    getLandingGears(){
        return axios.get(LANDINGGEAR_API_BASE_URL + '/' );
    }

    createLandingGear(landingGear){
        return axios.post(LANDINGGEAR_API_BASE_URL  + '/create', landingGear);
    }

    getLandingGearById(landingGearId){
        return axios.get(LANDINGGEAR_API_BASE_URL + '/load?landingGearId=' + landingGearId);
    }

    updateLandingGear(landingGear){
        return axios.put(LANDINGGEAR_API_BASE_URL + '/update', landingGear);
    }

    deleteLandingGear(landingGearId){
        return axios.delete(LANDINGGEAR_API_BASE_URL + '/delete?landingGearId=' + landingGearId);
    }
}

export default new LandingGearService()