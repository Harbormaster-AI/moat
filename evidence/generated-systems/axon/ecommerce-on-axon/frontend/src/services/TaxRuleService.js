import axios from 'axios';

const TAXRULE_API_BASE_URL = "/TaxRule";

class TaxRuleService {

    getTaxRules(){
        return axios.get(TAXRULE_API_BASE_URL + '/' );
    }

    createTaxRule(taxRule){
        return axios.post(TAXRULE_API_BASE_URL  + '/create', taxRule);
    }

    getTaxRuleById(taxRuleId){
        return axios.get(TAXRULE_API_BASE_URL + '/load?taxRuleId=' + taxRuleId);
    }

    updateTaxRule(taxRule){
        return axios.put(TAXRULE_API_BASE_URL + '/update', taxRule);
    }

    deleteTaxRule(taxRuleId){
        return axios.delete(TAXRULE_API_BASE_URL + '/delete?taxRuleId=' + taxRuleId);
    }
}

export default new TaxRuleService()