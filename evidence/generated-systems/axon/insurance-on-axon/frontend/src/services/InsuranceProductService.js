import axios from 'axios';

const INSURANCEPRODUCT_API_BASE_URL = "/InsuranceProduct";

class InsuranceProductService {

    getInsuranceProducts(){
        return axios.get(INSURANCEPRODUCT_API_BASE_URL + '/' );
    }

    createInsuranceProduct(insuranceProduct){
        return axios.post(INSURANCEPRODUCT_API_BASE_URL  + '/create', insuranceProduct);
    }

    getInsuranceProductById(insuranceProductId){
        return axios.get(INSURANCEPRODUCT_API_BASE_URL + '/load?insuranceProductId=' + insuranceProductId);
    }

    updateInsuranceProduct(insuranceProduct){
        return axios.put(INSURANCEPRODUCT_API_BASE_URL + '/update', insuranceProduct);
    }

    deleteInsuranceProduct(insuranceProductId){
        return axios.delete(INSURANCEPRODUCT_API_BASE_URL + '/delete?insuranceProductId=' + insuranceProductId);
    }
}

export default new InsuranceProductService()