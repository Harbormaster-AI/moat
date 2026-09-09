import axios from 'axios';

const DISCHARGE_API_BASE_URL = "/Discharge";

class DischargeService {

    getDischarges(){
        return axios.get(DISCHARGE_API_BASE_URL + '/' );
    }

    createDischarge(discharge){
        return axios.post(DISCHARGE_API_BASE_URL  + '/create', discharge);
    }

    getDischargeById(dischargeId){
        return axios.get(DISCHARGE_API_BASE_URL + '/load?dischargeId=' + dischargeId);
    }

    updateDischarge(discharge){
        return axios.put(DISCHARGE_API_BASE_URL + '/update', discharge);
    }

    deleteDischarge(dischargeId){
        return axios.delete(DISCHARGE_API_BASE_URL + '/delete?dischargeId=' + dischargeId);
    }
}

export default new DischargeService()