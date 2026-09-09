import axios from 'axios';

const BENEFITENROLLMENT_API_BASE_URL = "/BenefitEnrollment";

class BenefitEnrollmentService {

    getBenefitEnrollments(){
        return axios.get(BENEFITENROLLMENT_API_BASE_URL + '/' );
    }

    createBenefitEnrollment(benefitEnrollment){
        return axios.post(BENEFITENROLLMENT_API_BASE_URL  + '/create', benefitEnrollment);
    }

    getBenefitEnrollmentById(benefitEnrollmentId){
        return axios.get(BENEFITENROLLMENT_API_BASE_URL + '/load?benefitEnrollmentId=' + benefitEnrollmentId);
    }

    updateBenefitEnrollment(benefitEnrollment){
        return axios.put(BENEFITENROLLMENT_API_BASE_URL + '/update', benefitEnrollment);
    }

    deleteBenefitEnrollment(benefitEnrollmentId){
        return axios.delete(BENEFITENROLLMENT_API_BASE_URL + '/delete?benefitEnrollmentId=' + benefitEnrollmentId);
    }
}

export default new BenefitEnrollmentService()