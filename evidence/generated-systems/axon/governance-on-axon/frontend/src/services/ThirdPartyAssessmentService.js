import axios from 'axios';

const THIRDPARTYASSESSMENT_API_BASE_URL = "/ThirdPartyAssessment";

class ThirdPartyAssessmentService {

    getThirdPartyAssessments(){
        return axios.get(THIRDPARTYASSESSMENT_API_BASE_URL + '/' );
    }

    createThirdPartyAssessment(thirdPartyAssessment){
        return axios.post(THIRDPARTYASSESSMENT_API_BASE_URL  + '/create', thirdPartyAssessment);
    }

    getThirdPartyAssessmentById(thirdPartyAssessmentId){
        return axios.get(THIRDPARTYASSESSMENT_API_BASE_URL + '/load?thirdPartyAssessmentId=' + thirdPartyAssessmentId);
    }

    updateThirdPartyAssessment(thirdPartyAssessment){
        return axios.put(THIRDPARTYASSESSMENT_API_BASE_URL + '/update', thirdPartyAssessment);
    }

    deleteThirdPartyAssessment(thirdPartyAssessmentId){
        return axios.delete(THIRDPARTYASSESSMENT_API_BASE_URL + '/delete?thirdPartyAssessmentId=' + thirdPartyAssessmentId);
    }
}

export default new ThirdPartyAssessmentService()