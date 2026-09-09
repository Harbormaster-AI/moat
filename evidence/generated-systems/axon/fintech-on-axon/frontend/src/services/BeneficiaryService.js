import axios from 'axios';

const BENEFICIARY_API_BASE_URL = "/Beneficiary";

class BeneficiaryService {

    getBeneficiarys(){
        return axios.get(BENEFICIARY_API_BASE_URL + '/' );
    }

    createBeneficiary(beneficiary){
        return axios.post(BENEFICIARY_API_BASE_URL  + '/create', beneficiary);
    }

    getBeneficiaryById(beneficiaryId){
        return axios.get(BENEFICIARY_API_BASE_URL + '/load?beneficiaryId=' + beneficiaryId);
    }

    updateBeneficiary(beneficiary){
        return axios.put(BENEFICIARY_API_BASE_URL + '/update', beneficiary);
    }

    deleteBeneficiary(beneficiaryId){
        return axios.delete(BENEFICIARY_API_BASE_URL + '/delete?beneficiaryId=' + beneficiaryId);
    }
}

export default new BeneficiaryService()