import axios from 'axios';

const SALARYCOMPONENT_API_BASE_URL = "/SalaryComponent";

class SalaryComponentService {

    getSalaryComponents(){
        return axios.get(SALARYCOMPONENT_API_BASE_URL + '/' );
    }

    createSalaryComponent(salaryComponent){
        return axios.post(SALARYCOMPONENT_API_BASE_URL  + '/create', salaryComponent);
    }

    getSalaryComponentById(salaryComponentId){
        return axios.get(SALARYCOMPONENT_API_BASE_URL + '/load?salaryComponentId=' + salaryComponentId);
    }

    updateSalaryComponent(salaryComponent){
        return axios.put(SALARYCOMPONENT_API_BASE_URL + '/update', salaryComponent);
    }

    deleteSalaryComponent(salaryComponentId){
        return axios.delete(SALARYCOMPONENT_API_BASE_URL + '/delete?salaryComponentId=' + salaryComponentId);
    }
}

export default new SalaryComponentService()