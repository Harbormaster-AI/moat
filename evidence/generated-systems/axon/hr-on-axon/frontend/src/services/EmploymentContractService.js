import axios from 'axios';

const EMPLOYMENTCONTRACT_API_BASE_URL = "/EmploymentContract";

class EmploymentContractService {

    getEmploymentContracts(){
        return axios.get(EMPLOYMENTCONTRACT_API_BASE_URL + '/' );
    }

    createEmploymentContract(employmentContract){
        return axios.post(EMPLOYMENTCONTRACT_API_BASE_URL  + '/create', employmentContract);
    }

    getEmploymentContractById(employmentContractId){
        return axios.get(EMPLOYMENTCONTRACT_API_BASE_URL + '/load?employmentContractId=' + employmentContractId);
    }

    updateEmploymentContract(employmentContract){
        return axios.put(EMPLOYMENTCONTRACT_API_BASE_URL + '/update', employmentContract);
    }

    deleteEmploymentContract(employmentContractId){
        return axios.delete(EMPLOYMENTCONTRACT_API_BASE_URL + '/delete?employmentContractId=' + employmentContractId);
    }
}

export default new EmploymentContractService()