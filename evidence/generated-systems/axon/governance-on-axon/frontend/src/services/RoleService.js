import axios from 'axios';

const ROLE_API_BASE_URL = "/Role";

class RoleService {

    getRoles(){
        return axios.get(ROLE_API_BASE_URL + '/' );
    }

    createRole(role){
        return axios.post(ROLE_API_BASE_URL  + '/create', role);
    }

    getRoleById(roleId){
        return axios.get(ROLE_API_BASE_URL + '/load?roleId=' + roleId);
    }

    updateRole(role){
        return axios.put(ROLE_API_BASE_URL + '/update', role);
    }

    deleteRole(roleId){
        return axios.delete(ROLE_API_BASE_URL + '/delete?roleId=' + roleId);
    }
}

export default new RoleService()