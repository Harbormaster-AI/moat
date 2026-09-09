import axios from 'axios';

const SHIFTASSIGNMENT_API_BASE_URL = "/ShiftAssignment";

class ShiftAssignmentService {

    getShiftAssignments(){
        return axios.get(SHIFTASSIGNMENT_API_BASE_URL + '/' );
    }

    createShiftAssignment(shiftAssignment){
        return axios.post(SHIFTASSIGNMENT_API_BASE_URL  + '/create', shiftAssignment);
    }

    getShiftAssignmentById(shiftAssignmentId){
        return axios.get(SHIFTASSIGNMENT_API_BASE_URL + '/load?shiftAssignmentId=' + shiftAssignmentId);
    }

    updateShiftAssignment(shiftAssignment){
        return axios.put(SHIFTASSIGNMENT_API_BASE_URL + '/update', shiftAssignment);
    }

    deleteShiftAssignment(shiftAssignmentId){
        return axios.delete(SHIFTASSIGNMENT_API_BASE_URL + '/delete?shiftAssignmentId=' + shiftAssignmentId);
    }
}

export default new ShiftAssignmentService()