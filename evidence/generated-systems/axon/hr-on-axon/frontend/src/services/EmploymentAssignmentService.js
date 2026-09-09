import axios from 'axios';

const EMPLOYMENTASSIGNMENT_API_BASE_URL = "/EmploymentAssignment";

class EmploymentAssignmentService {

    getEmploymentAssignments(){
        return axios.get(EMPLOYMENTASSIGNMENT_API_BASE_URL + '/' );
    }

    createEmploymentAssignment(employmentAssignment){
        return axios.post(EMPLOYMENTASSIGNMENT_API_BASE_URL  + '/create', employmentAssignment);
    }

    getEmploymentAssignmentById(employmentAssignmentId){
        return axios.get(EMPLOYMENTASSIGNMENT_API_BASE_URL + '/load?employmentAssignmentId=' + employmentAssignmentId);
    }

    updateEmploymentAssignment(employmentAssignment){
        return axios.put(EMPLOYMENTASSIGNMENT_API_BASE_URL + '/update', employmentAssignment);
    }

    deleteEmploymentAssignment(employmentAssignmentId){
        return axios.delete(EMPLOYMENTASSIGNMENT_API_BASE_URL + '/delete?employmentAssignmentId=' + employmentAssignmentId);
    }
}

export default new EmploymentAssignmentService()