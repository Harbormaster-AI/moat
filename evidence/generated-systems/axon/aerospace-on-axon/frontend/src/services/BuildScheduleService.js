import axios from 'axios';

const BUILDSCHEDULE_API_BASE_URL = "/BuildSchedule";

class BuildScheduleService {

    getBuildSchedules(){
        return axios.get(BUILDSCHEDULE_API_BASE_URL + '/' );
    }

    createBuildSchedule(buildSchedule){
        return axios.post(BUILDSCHEDULE_API_BASE_URL  + '/create', buildSchedule);
    }

    getBuildScheduleById(buildScheduleId){
        return axios.get(BUILDSCHEDULE_API_BASE_URL + '/load?buildScheduleId=' + buildScheduleId);
    }

    updateBuildSchedule(buildSchedule){
        return axios.put(BUILDSCHEDULE_API_BASE_URL + '/update', buildSchedule);
    }

    deleteBuildSchedule(buildScheduleId){
        return axios.delete(BUILDSCHEDULE_API_BASE_URL + '/delete?buildScheduleId=' + buildScheduleId);
    }
}

export default new BuildScheduleService()