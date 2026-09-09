import axios from 'axios';

const RETENTIONSCHEDULE_API_BASE_URL = "/RetentionSchedule";

class RetentionScheduleService {

    getRetentionSchedules(){
        return axios.get(RETENTIONSCHEDULE_API_BASE_URL + '/' );
    }

    createRetentionSchedule(retentionSchedule){
        return axios.post(RETENTIONSCHEDULE_API_BASE_URL  + '/create', retentionSchedule);
    }

    getRetentionScheduleById(retentionScheduleId){
        return axios.get(RETENTIONSCHEDULE_API_BASE_URL + '/load?retentionScheduleId=' + retentionScheduleId);
    }

    updateRetentionSchedule(retentionSchedule){
        return axios.put(RETENTIONSCHEDULE_API_BASE_URL + '/update', retentionSchedule);
    }

    deleteRetentionSchedule(retentionScheduleId){
        return axios.delete(RETENTIONSCHEDULE_API_BASE_URL + '/delete?retentionScheduleId=' + retentionScheduleId);
    }
}

export default new RetentionScheduleService()