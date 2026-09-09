import axios from 'axios';

const CARETEAM_API_BASE_URL = "/CareTeam";

class CareTeamService {

    getCareTeams(){
        return axios.get(CARETEAM_API_BASE_URL + '/' );
    }

    createCareTeam(careTeam){
        return axios.post(CARETEAM_API_BASE_URL  + '/create', careTeam);
    }

    getCareTeamById(careTeamId){
        return axios.get(CARETEAM_API_BASE_URL + '/load?careTeamId=' + careTeamId);
    }

    updateCareTeam(careTeam){
        return axios.put(CARETEAM_API_BASE_URL + '/update', careTeam);
    }

    deleteCareTeam(careTeamId){
        return axios.delete(CARETEAM_API_BASE_URL + '/delete?careTeamId=' + careTeamId);
    }
}

export default new CareTeamService()