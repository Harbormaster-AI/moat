import axios from 'axios';

const CYCLECOUNT_API_BASE_URL = "/CycleCount";

class CycleCountService {

    getCycleCounts(){
        return axios.get(CYCLECOUNT_API_BASE_URL + '/' );
    }

    createCycleCount(cycleCount){
        return axios.post(CYCLECOUNT_API_BASE_URL  + '/create', cycleCount);
    }

    getCycleCountById(cycleCountId){
        return axios.get(CYCLECOUNT_API_BASE_URL + '/load?cycleCountId=' + cycleCountId);
    }

    updateCycleCount(cycleCount){
        return axios.put(CYCLECOUNT_API_BASE_URL + '/update', cycleCount);
    }

    deleteCycleCount(cycleCountId){
        return axios.delete(CYCLECOUNT_API_BASE_URL + '/delete?cycleCountId=' + cycleCountId);
    }
}

export default new CycleCountService()