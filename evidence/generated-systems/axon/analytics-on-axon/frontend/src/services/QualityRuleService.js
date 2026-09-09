import axios from 'axios';

const QUALITYRULE_API_BASE_URL = "/QualityRule";

class QualityRuleService {

    getQualityRules(){
        return axios.get(QUALITYRULE_API_BASE_URL + '/' );
    }

    createQualityRule(qualityRule){
        return axios.post(QUALITYRULE_API_BASE_URL  + '/create', qualityRule);
    }

    getQualityRuleById(qualityRuleId){
        return axios.get(QUALITYRULE_API_BASE_URL + '/load?qualityRuleId=' + qualityRuleId);
    }

    updateQualityRule(qualityRule){
        return axios.put(QUALITYRULE_API_BASE_URL + '/update', qualityRule);
    }

    deleteQualityRule(qualityRuleId){
        return axios.delete(QUALITYRULE_API_BASE_URL + '/delete?qualityRuleId=' + qualityRuleId);
    }
}

export default new QualityRuleService()