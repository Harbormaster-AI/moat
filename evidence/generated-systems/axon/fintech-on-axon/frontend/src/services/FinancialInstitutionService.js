import axios from 'axios';

const FINANCIALINSTITUTION_API_BASE_URL = "/FinancialInstitution";

class FinancialInstitutionService {

    getFinancialInstitutions(){
        return axios.get(FINANCIALINSTITUTION_API_BASE_URL + '/' );
    }

    createFinancialInstitution(financialInstitution){
        return axios.post(FINANCIALINSTITUTION_API_BASE_URL  + '/create', financialInstitution);
    }

    getFinancialInstitutionById(financialInstitutionId){
        return axios.get(FINANCIALINSTITUTION_API_BASE_URL + '/load?financialInstitutionId=' + financialInstitutionId);
    }

    updateFinancialInstitution(financialInstitution){
        return axios.put(FINANCIALINSTITUTION_API_BASE_URL + '/update', financialInstitution);
    }

    deleteFinancialInstitution(financialInstitutionId){
        return axios.delete(FINANCIALINSTITUTION_API_BASE_URL + '/delete?financialInstitutionId=' + financialInstitutionId);
    }
}

export default new FinancialInstitutionService()