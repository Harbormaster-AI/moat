import axios from 'axios';

const ADJUSTER_API_BASE_URL = "/Adjuster";

class AdjusterService {

    getAdjusters(){
        return axios.get(ADJUSTER_API_BASE_URL + '/' );
    }

    createAdjuster(adjuster){
        return axios.post(ADJUSTER_API_BASE_URL  + '/create', adjuster);
    }

    getAdjusterById(adjusterId){
        return axios.get(ADJUSTER_API_BASE_URL + '/load?adjusterId=' + adjusterId);
    }

    updateAdjuster(adjuster){
        return axios.put(ADJUSTER_API_BASE_URL + '/update', adjuster);
    }

    deleteAdjuster(adjusterId){
        return axios.delete(ADJUSTER_API_BASE_URL + '/delete?adjusterId=' + adjusterId);
    }
}

export default new AdjusterService()