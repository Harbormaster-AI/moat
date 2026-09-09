import axios from 'axios';

const AIRWORTHINESSDIRECTIVE_API_BASE_URL = "/AirworthinessDirective";

class AirworthinessDirectiveService {

    getAirworthinessDirectives(){
        return axios.get(AIRWORTHINESSDIRECTIVE_API_BASE_URL + '/' );
    }

    createAirworthinessDirective(airworthinessDirective){
        return axios.post(AIRWORTHINESSDIRECTIVE_API_BASE_URL  + '/create', airworthinessDirective);
    }

    getAirworthinessDirectiveById(airworthinessDirectiveId){
        return axios.get(AIRWORTHINESSDIRECTIVE_API_BASE_URL + '/load?airworthinessDirectiveId=' + airworthinessDirectiveId);
    }

    updateAirworthinessDirective(airworthinessDirective){
        return axios.put(AIRWORTHINESSDIRECTIVE_API_BASE_URL + '/update', airworthinessDirective);
    }

    deleteAirworthinessDirective(airworthinessDirectiveId){
        return axios.delete(AIRWORTHINESSDIRECTIVE_API_BASE_URL + '/delete?airworthinessDirectiveId=' + airworthinessDirectiveId);
    }
}

export default new AirworthinessDirectiveService()