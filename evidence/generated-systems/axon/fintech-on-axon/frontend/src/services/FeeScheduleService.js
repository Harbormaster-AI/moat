import axios from 'axios';

const FEESCHEDULE_API_BASE_URL = "/FeeSchedule";

class FeeScheduleService {

    getFeeSchedules(){
        return axios.get(FEESCHEDULE_API_BASE_URL + '/' );
    }

    createFeeSchedule(feeSchedule){
        return axios.post(FEESCHEDULE_API_BASE_URL  + '/create', feeSchedule);
    }

    getFeeScheduleById(feeScheduleId){
        return axios.get(FEESCHEDULE_API_BASE_URL + '/load?feeScheduleId=' + feeScheduleId);
    }

    updateFeeSchedule(feeSchedule){
        return axios.put(FEESCHEDULE_API_BASE_URL + '/update', feeSchedule);
    }

    deleteFeeSchedule(feeScheduleId){
        return axios.delete(FEESCHEDULE_API_BASE_URL + '/delete?feeScheduleId=' + feeScheduleId);
    }
}

export default new FeeScheduleService()