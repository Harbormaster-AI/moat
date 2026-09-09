import axios from 'axios';

const DEPARTMENT_API_BASE_URL = "/Department";

class DepartmentService {

    getDepartments(){
        return axios.get(DEPARTMENT_API_BASE_URL + '/' );
    }

    createDepartment(department){
        return axios.post(DEPARTMENT_API_BASE_URL  + '/create', department);
    }

    getDepartmentById(departmentId){
        return axios.get(DEPARTMENT_API_BASE_URL + '/load?departmentId=' + departmentId);
    }

    updateDepartment(department){
        return axios.put(DEPARTMENT_API_BASE_URL + '/update', department);
    }

    deleteDepartment(departmentId){
        return axios.delete(DEPARTMENT_API_BASE_URL + '/delete?departmentId=' + departmentId);
    }
}

export default new DepartmentService()