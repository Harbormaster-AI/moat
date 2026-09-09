import axios from 'axios';

const CYCLECOUNTENTRY_API_BASE_URL = "/CycleCountEntry";

class CycleCountEntryService {

    getCycleCountEntrys(){
        return axios.get(CYCLECOUNTENTRY_API_BASE_URL + '/' );
    }

    createCycleCountEntry(cycleCountEntry){
        return axios.post(CYCLECOUNTENTRY_API_BASE_URL  + '/create', cycleCountEntry);
    }

    getCycleCountEntryById(cycleCountEntryId){
        return axios.get(CYCLECOUNTENTRY_API_BASE_URL + '/load?cycleCountEntryId=' + cycleCountEntryId);
    }

    updateCycleCountEntry(cycleCountEntry){
        return axios.put(CYCLECOUNTENTRY_API_BASE_URL + '/update', cycleCountEntry);
    }

    deleteCycleCountEntry(cycleCountEntryId){
        return axios.delete(CYCLECOUNTENTRY_API_BASE_URL + '/delete?cycleCountEntryId=' + cycleCountEntryId);
    }
}

export default new CycleCountEntryService()