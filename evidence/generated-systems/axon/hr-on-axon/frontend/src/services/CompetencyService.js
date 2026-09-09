import axios from 'axios';

const COMPETENCY_API_BASE_URL = "/Competency";

class CompetencyService {

    getCompetencys(){
        return axios.get(COMPETENCY_API_BASE_URL + '/' );
    }

    createCompetency(competency){
        return axios.post(COMPETENCY_API_BASE_URL  + '/create', competency);
    }

    getCompetencyById(competencyId){
        return axios.get(COMPETENCY_API_BASE_URL + '/load?competencyId=' + competencyId);
    }

    updateCompetency(competency){
        return axios.put(COMPETENCY_API_BASE_URL + '/update', competency);
    }

    deleteCompetency(competencyId){
        return axios.delete(COMPETENCY_API_BASE_URL + '/delete?competencyId=' + competencyId);
    }
}

export default new CompetencyService()