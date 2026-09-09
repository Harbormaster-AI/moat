import axios from 'axios';

const LOAN_API_BASE_URL = "/Loan";

class LoanService {

    getLoans(){
        return axios.get(LOAN_API_BASE_URL + '/' );
    }

    createLoan(loan){
        return axios.post(LOAN_API_BASE_URL  + '/create', loan);
    }

    getLoanById(loanId){
        return axios.get(LOAN_API_BASE_URL + '/load?loanId=' + loanId);
    }

    updateLoan(loan){
        return axios.put(LOAN_API_BASE_URL + '/update', loan);
    }

    deleteLoan(loanId){
        return axios.delete(LOAN_API_BASE_URL + '/delete?loanId=' + loanId);
    }
}

export default new LoanService()