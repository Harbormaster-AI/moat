import axios from 'axios';

const ROLEASSIGNMENT_API_BASE_URL = "/RoleAssignment";

class RoleAssignmentService {

    getRoleAssignments(){
        return axios.get(ROLEASSIGNMENT_API_BASE_URL + '/' );
    }

    createRoleAssignment(roleAssignment){
        return axios.post(ROLEASSIGNMENT_API_BASE_URL  + '/create', roleAssignment);
    }

    getRoleAssignmentById(roleAssignmentId){
        return axios.get(ROLEASSIGNMENT_API_BASE_URL + '/load?roleAssignmentId=' + roleAssignmentId);
    }

    updateRoleAssignment(roleAssignment){
        return axios.put(ROLEASSIGNMENT_API_BASE_URL + '/update', roleAssignment);
    }

    deleteRoleAssignment(roleAssignmentId){
        return axios.delete(ROLEASSIGNMENT_API_BASE_URL + '/delete?roleAssignmentId=' + roleAssignmentId);
    }
}

export default new RoleAssignmentService()