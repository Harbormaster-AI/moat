import axios from 'axios';

const WORKSCHEDULE_API_BASE_URL = "/WorkSchedule";

class WorkScheduleService {

    getWorkSchedules(){
        return axios.get(WORKSCHEDULE_API_BASE_URL + '/' );
    }

    createWorkSchedule(workSchedule){
        return axios.post(WORKSCHEDULE_API_BASE_URL  + '/create', workSchedule);
    }

    getWorkScheduleById(workScheduleId){
        return axios.get(WORKSCHEDULE_API_BASE_URL + '/load?workScheduleId=' + workScheduleId);
    }

    updateWorkSchedule(workSchedule){
        return axios.put(WORKSCHEDULE_API_BASE_URL + '/update', workSchedule);
    }

    deleteWorkSchedule(workScheduleId){
        return axios.delete(WORKSCHEDULE_API_BASE_URL + '/delete?workScheduleId=' + workScheduleId);
    }
}

export default new WorkScheduleService()