import axios from 'axios';

const MATTER_API_BASE_URL = "/Matter";

class MatterService {

    getMatters(){
        return axios.get(MATTER_API_BASE_URL + '/' );
    }

    createMatter(matter){
        return axios.post(MATTER_API_BASE_URL  + '/create', matter);
    }

    getMatterById(matterId){
        return axios.get(MATTER_API_BASE_URL + '/load?matterId=' + matterId);
    }

    updateMatter(matter){
        return axios.put(MATTER_API_BASE_URL + '/update', matter);
    }

    deleteMatter(matterId){
        return axios.delete(MATTER_API_BASE_URL + '/delete?matterId=' + matterId);
    }
}

export default new MatterService()