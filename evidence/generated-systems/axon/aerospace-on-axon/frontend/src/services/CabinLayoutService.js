import axios from 'axios';

const CABINLAYOUT_API_BASE_URL = "/CabinLayout";

class CabinLayoutService {

    getCabinLayouts(){
        return axios.get(CABINLAYOUT_API_BASE_URL + '/' );
    }

    createCabinLayout(cabinLayout){
        return axios.post(CABINLAYOUT_API_BASE_URL  + '/create', cabinLayout);
    }

    getCabinLayoutById(cabinLayoutId){
        return axios.get(CABINLAYOUT_API_BASE_URL + '/load?cabinLayoutId=' + cabinLayoutId);
    }

    updateCabinLayout(cabinLayout){
        return axios.put(CABINLAYOUT_API_BASE_URL + '/update', cabinLayout);
    }

    deleteCabinLayout(cabinLayoutId){
        return axios.delete(CABINLAYOUT_API_BASE_URL + '/delete?cabinLayoutId=' + cabinLayoutId);
    }
}

export default new CabinLayoutService()