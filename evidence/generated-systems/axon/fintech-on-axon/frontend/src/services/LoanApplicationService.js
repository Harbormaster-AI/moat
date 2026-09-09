import axios from 'axios';

const LOANAPPLICATION_API_BASE_URL = "/LoanApplication";

class LoanApplicationService {

    getLoanApplications(){
        return axios.get(LOANAPPLICATION_API_BASE_URL + '/' );
    }

    createLoanApplication(loanApplication){
        return axios.post(LOANAPPLICATION_API_BASE_URL  + '/create', loanApplication);
    }

    getLoanApplicationById(loanApplicationId){
        return axios.get(LOANAPPLICATION_API_BASE_URL + '/load?loanApplicationId=' + loanApplicationId);
    }

    updateLoanApplication(loanApplication){
        return axios.put(LOANAPPLICATION_API_BASE_URL + '/update', loanApplication);
    }

    deleteLoanApplication(loanApplicationId){
        return axios.delete(LOANAPPLICATION_API_BASE_URL + '/delete?loanApplicationId=' + loanApplicationId);
    }
}

export default new LoanApplicationService()