import axios from 'axios';

const WORKSHIFT_API_BASE_URL = "/WorkShift";

class WorkShiftService {

    getWorkShifts(){
        return axios.get(WORKSHIFT_API_BASE_URL + '/' );
    }

    createWorkShift(workShift){
        return axios.post(WORKSHIFT_API_BASE_URL  + '/create', workShift);
    }

    getWorkShiftById(workShiftId){
        return axios.get(WORKSHIFT_API_BASE_URL + '/load?workShiftId=' + workShiftId);
    }

    updateWorkShift(workShift){
        return axios.put(WORKSHIFT_API_BASE_URL + '/update', workShift);
    }

    deleteWorkShift(workShiftId){
        return axios.delete(WORKSHIFT_API_BASE_URL + '/delete?workShiftId=' + workShiftId);
    }
}

export default new WorkShiftService()