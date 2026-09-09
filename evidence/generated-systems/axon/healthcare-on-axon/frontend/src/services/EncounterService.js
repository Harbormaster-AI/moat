import axios from 'axios';

const ENCOUNTER_API_BASE_URL = "/Encounter";

class EncounterService {

    getEncounters(){
        return axios.get(ENCOUNTER_API_BASE_URL + '/' );
    }

    createEncounter(encounter){
        return axios.post(ENCOUNTER_API_BASE_URL  + '/create', encounter);
    }

    getEncounterById(encounterId){
        return axios.get(ENCOUNTER_API_BASE_URL + '/load?encounterId=' + encounterId);
    }

    updateEncounter(encounter){
        return axios.put(ENCOUNTER_API_BASE_URL + '/update', encounter);
    }

    deleteEncounter(encounterId){
        return axios.delete(ENCOUNTER_API_BASE_URL + '/delete?encounterId=' + encounterId);
    }
}

export default new EncounterService()