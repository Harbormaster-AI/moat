import axios from 'axios';

const CONTRACT_API_BASE_URL = "/Contract";

class ContractService {

    getContracts(){
        return axios.get(CONTRACT_API_BASE_URL + '/' );
    }

    createContract(contract){
        return axios.post(CONTRACT_API_BASE_URL  + '/create', contract);
    }

    getContractById(contractId){
        return axios.get(CONTRACT_API_BASE_URL + '/load?contractId=' + contractId);
    }

    updateContract(contract){
        return axios.put(CONTRACT_API_BASE_URL + '/update', contract);
    }

    deleteContract(contractId){
        return axios.delete(CONTRACT_API_BASE_URL + '/delete?contractId=' + contractId);
    }
}

export default new ContractService()