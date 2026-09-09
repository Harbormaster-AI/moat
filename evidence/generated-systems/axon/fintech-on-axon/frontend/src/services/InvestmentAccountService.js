import axios from 'axios';

const INVESTMENTACCOUNT_API_BASE_URL = "/InvestmentAccount";

class InvestmentAccountService {

    getInvestmentAccounts(){
        return axios.get(INVESTMENTACCOUNT_API_BASE_URL + '/' );
    }

    createInvestmentAccount(investmentAccount){
        return axios.post(INVESTMENTACCOUNT_API_BASE_URL  + '/create', investmentAccount);
    }

    getInvestmentAccountById(investmentAccountId){
        return axios.get(INVESTMENTACCOUNT_API_BASE_URL + '/load?investmentAccountId=' + investmentAccountId);
    }

    updateInvestmentAccount(investmentAccount){
        return axios.put(INVESTMENTACCOUNT_API_BASE_URL + '/update', investmentAccount);
    }

    deleteInvestmentAccount(investmentAccountId){
        return axios.delete(INVESTMENTACCOUNT_API_BASE_URL + '/delete?investmentAccountId=' + investmentAccountId);
    }
}

export default new InvestmentAccountService()