import axios from 'axios';

const SCHEDULEEXCEPTION_API_BASE_URL = "/ScheduleException";

class ScheduleExceptionService {

    getScheduleExceptions(){
        return axios.get(SCHEDULEEXCEPTION_API_BASE_URL + '/' );
    }

    createScheduleException(scheduleException){
        return axios.post(SCHEDULEEXCEPTION_API_BASE_URL  + '/create', scheduleException);
    }

    getScheduleExceptionById(scheduleExceptionId){
        return axios.get(SCHEDULEEXCEPTION_API_BASE_URL + '/load?scheduleExceptionId=' + scheduleExceptionId);
    }

    updateScheduleException(scheduleException){
        return axios.put(SCHEDULEEXCEPTION_API_BASE_URL + '/update', scheduleException);
    }

    deleteScheduleException(scheduleExceptionId){
        return axios.delete(SCHEDULEEXCEPTION_API_BASE_URL + '/delete?scheduleExceptionId=' + scheduleExceptionId);
    }
}

export default new ScheduleExceptionService()