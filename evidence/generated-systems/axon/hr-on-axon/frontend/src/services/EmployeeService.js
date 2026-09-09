import axios from 'axios';

const EMPLOYEE_API_BASE_URL = "/Employee";

class EmployeeService {

    getEmployees(){
        return axios.get(EMPLOYEE_API_BASE_URL + '/' );
    }

    createEmployee(employee){
        return axios.post(EMPLOYEE_API_BASE_URL  + '/create', employee);
    }

    getEmployeeById(employeeId){
        return axios.get(EMPLOYEE_API_BASE_URL + '/load?employeeId=' + employeeId);
    }

    updateEmployee(employee){
        return axios.put(EMPLOYEE_API_BASE_URL + '/update', employee);
    }

    deleteEmployee(employeeId){
        return axios.delete(EMPLOYEE_API_BASE_URL + '/delete?employeeId=' + employeeId);
    }
}

export default new EmployeeService()