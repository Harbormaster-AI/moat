import axios from 'axios';

const DEMANDSIGNAL_API_BASE_URL = "/DemandSignal";

class DemandSignalService {

    getDemandSignals(){
        return axios.get(DEMANDSIGNAL_API_BASE_URL + '/' );
    }

    createDemandSignal(demandSignal){
        return axios.post(DEMANDSIGNAL_API_BASE_URL  + '/create', demandSignal);
    }

    getDemandSignalById(demandSignalId){
        return axios.get(DEMANDSIGNAL_API_BASE_URL + '/load?demandSignalId=' + demandSignalId);
    }

    updateDemandSignal(demandSignal){
        return axios.put(DEMANDSIGNAL_API_BASE_URL + '/update', demandSignal);
    }

    deleteDemandSignal(demandSignalId){
        return axios.delete(DEMANDSIGNAL_API_BASE_URL + '/delete?demandSignalId=' + demandSignalId);
    }
}

export default new DemandSignalService()