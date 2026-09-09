import axios from 'axios';

const INSURANCEPAYER_API_BASE_URL = "/InsurancePayer";

class InsurancePayerService {

    getInsurancePayers(){
        return axios.get(INSURANCEPAYER_API_BASE_URL + '/' );
    }

    createInsurancePayer(insurancePayer){
        return axios.post(INSURANCEPAYER_API_BASE_URL  + '/create', insurancePayer);
    }

    getInsurancePayerById(insurancePayerId){
        return axios.get(INSURANCEPAYER_API_BASE_URL + '/load?insurancePayerId=' + insurancePayerId);
    }

    updateInsurancePayer(insurancePayer){
        return axios.put(INSURANCEPAYER_API_BASE_URL + '/update', insurancePayer);
    }

    deleteInsurancePayer(insurancePayerId){
        return axios.delete(INSURANCEPAYER_API_BASE_URL + '/delete?insurancePayerId=' + insurancePayerId);
    }
}

export default new InsurancePayerService()