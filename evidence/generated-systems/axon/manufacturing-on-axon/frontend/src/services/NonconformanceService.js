import axios from 'axios';

const NONCONFORMANCE_API_BASE_URL = "/Nonconformance";

class NonconformanceService {

    getNonconformances(){
        return axios.get(NONCONFORMANCE_API_BASE_URL + '/' );
    }

    createNonconformance(nonconformance){
        return axios.post(NONCONFORMANCE_API_BASE_URL  + '/create', nonconformance);
    }

    getNonconformanceById(nonconformanceId){
        return axios.get(NONCONFORMANCE_API_BASE_URL + '/load?nonconformanceId=' + nonconformanceId);
    }

    updateNonconformance(nonconformance){
        return axios.put(NONCONFORMANCE_API_BASE_URL + '/update', nonconformance);
    }

    deleteNonconformance(nonconformanceId){
        return axios.delete(NONCONFORMANCE_API_BASE_URL + '/delete?nonconformanceId=' + nonconformanceId);
    }
}

export default new NonconformanceService()