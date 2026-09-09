import axios from 'axios';

const APU_API_BASE_URL = "/APU";

class APUService {

    getAPUs(){
        return axios.get(APU_API_BASE_URL + '/' );
    }

    createAPU(aPU){
        return axios.post(APU_API_BASE_URL  + '/create', aPU);
    }

    getAPUById(aPUId){
        return axios.get(APU_API_BASE_URL + '/load?aPUId=' + aPUId);
    }

    updateAPU(aPU){
        return axios.put(APU_API_BASE_URL + '/update', aPU);
    }

    deleteAPU(aPUId){
        return axios.delete(APU_API_BASE_URL + '/delete?aPUId=' + aPUId);
    }
}

export default new APUService()