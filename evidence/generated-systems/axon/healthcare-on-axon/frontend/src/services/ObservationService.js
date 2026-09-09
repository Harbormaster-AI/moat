import axios from 'axios';

const OBSERVATION_API_BASE_URL = "/Observation";

class ObservationService {

    getObservations(){
        return axios.get(OBSERVATION_API_BASE_URL + '/' );
    }

    createObservation(observation){
        return axios.post(OBSERVATION_API_BASE_URL  + '/create', observation);
    }

    getObservationById(observationId){
        return axios.get(OBSERVATION_API_BASE_URL + '/load?observationId=' + observationId);
    }

    updateObservation(observation){
        return axios.put(OBSERVATION_API_BASE_URL + '/update', observation);
    }

    deleteObservation(observationId){
        return axios.delete(OBSERVATION_API_BASE_URL + '/delete?observationId=' + observationId);
    }
}

export default new ObservationService()