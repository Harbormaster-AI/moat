import axios from 'axios';

const INVESTMENTPORTFOLIO_API_BASE_URL = "/InvestmentPortfolio";

class InvestmentPortfolioService {

    getInvestmentPortfolios(){
        return axios.get(INVESTMENTPORTFOLIO_API_BASE_URL + '/' );
    }

    createInvestmentPortfolio(investmentPortfolio){
        return axios.post(INVESTMENTPORTFOLIO_API_BASE_URL  + '/create', investmentPortfolio);
    }

    getInvestmentPortfolioById(investmentPortfolioId){
        return axios.get(INVESTMENTPORTFOLIO_API_BASE_URL + '/load?investmentPortfolioId=' + investmentPortfolioId);
    }

    updateInvestmentPortfolio(investmentPortfolio){
        return axios.put(INVESTMENTPORTFOLIO_API_BASE_URL + '/update', investmentPortfolio);
    }

    deleteInvestmentPortfolio(investmentPortfolioId){
        return axios.delete(INVESTMENTPORTFOLIO_API_BASE_URL + '/delete?investmentPortfolioId=' + investmentPortfolioId);
    }
}

export default new InvestmentPortfolioService()