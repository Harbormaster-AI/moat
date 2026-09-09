import axios from 'axios';

const COMPETENCYRATING_API_BASE_URL = "/CompetencyRating";

class CompetencyRatingService {

    getCompetencyRatings(){
        return axios.get(COMPETENCYRATING_API_BASE_URL + '/' );
    }

    createCompetencyRating(competencyRating){
        return axios.post(COMPETENCYRATING_API_BASE_URL  + '/create', competencyRating);
    }

    getCompetencyRatingById(competencyRatingId){
        return axios.get(COMPETENCYRATING_API_BASE_URL + '/load?competencyRatingId=' + competencyRatingId);
    }

    updateCompetencyRating(competencyRating){
        return axios.put(COMPETENCYRATING_API_BASE_URL + '/update', competencyRating);
    }

    deleteCompetencyRating(competencyRatingId){
        return axios.delete(COMPETENCYRATING_API_BASE_URL + '/delete?competencyRatingId=' + competencyRatingId);
    }
}

export default new CompetencyRatingService()