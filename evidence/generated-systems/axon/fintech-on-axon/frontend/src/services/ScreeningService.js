import axios from 'axios';

const SCREENING_API_BASE_URL = "/Screening";

class ScreeningService {

    getScreenings(){
        return axios.get(SCREENING_API_BASE_URL + '/' );
    }

    createScreening(screening){
        return axios.post(SCREENING_API_BASE_URL  + '/create', screening);
    }

    getScreeningById(screeningId){
        return axios.get(SCREENING_API_BASE_URL + '/load?screeningId=' + screeningId);
    }

    updateScreening(screening){
        return axios.put(SCREENING_API_BASE_URL + '/update', screening);
    }

    deleteScreening(screeningId){
        return axios.delete(SCREENING_API_BASE_URL + '/delete?screeningId=' + screeningId);
    }
}

export default new ScreeningService()