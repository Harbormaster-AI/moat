import axios from 'axios';

const BUSINESSGLOSSARYTERM_API_BASE_URL = "/BusinessGlossaryTerm";

class BusinessGlossaryTermService {

    getBusinessGlossaryTerms(){
        return axios.get(BUSINESSGLOSSARYTERM_API_BASE_URL + '/' );
    }

    createBusinessGlossaryTerm(businessGlossaryTerm){
        return axios.post(BUSINESSGLOSSARYTERM_API_BASE_URL  + '/create', businessGlossaryTerm);
    }

    getBusinessGlossaryTermById(businessGlossaryTermId){
        return axios.get(BUSINESSGLOSSARYTERM_API_BASE_URL + '/load?businessGlossaryTermId=' + businessGlossaryTermId);
    }

    updateBusinessGlossaryTerm(businessGlossaryTerm){
        return axios.put(BUSINESSGLOSSARYTERM_API_BASE_URL + '/update', businessGlossaryTerm);
    }

    deleteBusinessGlossaryTerm(businessGlossaryTermId){
        return axios.delete(BUSINESSGLOSSARYTERM_API_BASE_URL + '/delete?businessGlossaryTermId=' + businessGlossaryTermId);
    }
}

export default new BusinessGlossaryTermService()