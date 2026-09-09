import axios from 'axios';

const BRANDSAFETYPOLICY_API_BASE_URL = "/BrandSafetyPolicy";

class BrandSafetyPolicyService {

    getBrandSafetyPolicys(){
        return axios.get(BRANDSAFETYPOLICY_API_BASE_URL + '/' );
    }

    createBrandSafetyPolicy(brandSafetyPolicy){
        return axios.post(BRANDSAFETYPOLICY_API_BASE_URL  + '/create', brandSafetyPolicy);
    }

    getBrandSafetyPolicyById(brandSafetyPolicyId){
        return axios.get(BRANDSAFETYPOLICY_API_BASE_URL + '/load?brandSafetyPolicyId=' + brandSafetyPolicyId);
    }

    updateBrandSafetyPolicy(brandSafetyPolicy){
        return axios.put(BRANDSAFETYPOLICY_API_BASE_URL + '/update', brandSafetyPolicy);
    }

    deleteBrandSafetyPolicy(brandSafetyPolicyId){
        return axios.delete(BRANDSAFETYPOLICY_API_BASE_URL + '/delete?brandSafetyPolicyId=' + brandSafetyPolicyId);
    }
}

export default new BrandSafetyPolicyService()