import axios from 'axios';

const SHIFT_API_BASE_URL = "/Shift";

class ShiftService {

    getShifts(){
        return axios.get(SHIFT_API_BASE_URL + '/' );
    }

    createShift(shift){
        return axios.post(SHIFT_API_BASE_URL  + '/create', shift);
    }

    getShiftById(shiftId){
        return axios.get(SHIFT_API_BASE_URL + '/load?shiftId=' + shiftId);
    }

    updateShift(shift){
        return axios.put(SHIFT_API_BASE_URL + '/update', shift);
    }

    deleteShift(shiftId){
        return axios.delete(SHIFT_API_BASE_URL + '/delete?shiftId=' + shiftId);
    }
}

export default new ShiftService()