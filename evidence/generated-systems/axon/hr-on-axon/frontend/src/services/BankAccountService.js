import axios from 'axios';

const BANKACCOUNT_API_BASE_URL = "/BankAccount";

class BankAccountService {

    getBankAccounts(){
        return axios.get(BANKACCOUNT_API_BASE_URL + '/' );
    }

    createBankAccount(bankAccount){
        return axios.post(BANKACCOUNT_API_BASE_URL  + '/create', bankAccount);
    }

    getBankAccountById(bankAccountId){
        return axios.get(BANKACCOUNT_API_BASE_URL + '/load?bankAccountId=' + bankAccountId);
    }

    updateBankAccount(bankAccount){
        return axios.put(BANKACCOUNT_API_BASE_URL + '/update', bankAccount);
    }

    deleteBankAccount(bankAccountId){
        return axios.delete(BANKACCOUNT_API_BASE_URL + '/delete?bankAccountId=' + bankAccountId);
    }
}

export default new BankAccountService()