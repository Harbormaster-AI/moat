import axios from 'axios';

const TEAM_API_BASE_URL = "/Team";

class TeamService {

    getTeams(){
        return axios.get(TEAM_API_BASE_URL + '/' );
    }

    createTeam(team){
        return axios.post(TEAM_API_BASE_URL  + '/create', team);
    }

    getTeamById(teamId){
        return axios.get(TEAM_API_BASE_URL + '/load?teamId=' + teamId);
    }

    updateTeam(team){
        return axios.put(TEAM_API_BASE_URL + '/update', team);
    }

    deleteTeam(teamId){
        return axios.delete(TEAM_API_BASE_URL + '/delete?teamId=' + teamId);
    }
}

export default new TeamService()