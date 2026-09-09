import axios from 'axios';

const DSP_API_BASE_URL = "/DSP";

class DSPService {

    getDSPs(){
        return axios.get(DSP_API_BASE_URL + '/' );
    }

    createDSP(dSP){
        return axios.post(DSP_API_BASE_URL  + '/create', dSP);
    }

    getDSPById(dSPId){
        return axios.get(DSP_API_BASE_URL + '/load?dSPId=' + dSPId);
    }

    updateDSP(dSP){
        return axios.put(DSP_API_BASE_URL + '/update', dSP);
    }

    deleteDSP(dSPId){
        return axios.delete(DSP_API_BASE_URL + '/delete?dSPId=' + dSPId);
    }
}

export default new DSPService()