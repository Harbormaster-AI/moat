import axios from 'axios';

const LOANTRANSACTION_API_BASE_URL = "/LoanTransaction";

class LoanTransactionService {

    getLoanTransactions(){
        return axios.get(LOANTRANSACTION_API_BASE_URL + '/' );
    }

    createLoanTransaction(loanTransaction){
        return axios.post(LOANTRANSACTION_API_BASE_URL  + '/create', loanTransaction);
    }

    getLoanTransactionById(loanTransactionId){
        return axios.get(LOANTRANSACTION_API_BASE_URL + '/load?loanTransactionId=' + loanTransactionId);
    }

    updateLoanTransaction(loanTransaction){
        return axios.put(LOANTRANSACTION_API_BASE_URL + '/update', loanTransaction);
    }

    deleteLoanTransaction(loanTransactionId){
        return axios.delete(LOANTRANSACTION_API_BASE_URL + '/delete?loanTransactionId=' + loanTransactionId);
    }
}

export default new LoanTransactionService()