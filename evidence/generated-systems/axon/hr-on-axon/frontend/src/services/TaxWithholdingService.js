import axios from 'axios';

const TAXWITHHOLDING_API_BASE_URL = "/TaxWithholding";

class TaxWithholdingService {

    getTaxWithholdings(){
        return axios.get(TAXWITHHOLDING_API_BASE_URL + '/' );
    }

    createTaxWithholding(taxWithholding){
        return axios.post(TAXWITHHOLDING_API_BASE_URL  + '/create', taxWithholding);
    }

    getTaxWithholdingById(taxWithholdingId){
        return axios.get(TAXWITHHOLDING_API_BASE_URL + '/load?taxWithholdingId=' + taxWithholdingId);
    }

    updateTaxWithholding(taxWithholding){
        return axios.put(TAXWITHHOLDING_API_BASE_URL + '/update', taxWithholding);
    }

    deleteTaxWithholding(taxWithholdingId){
        return axios.delete(TAXWITHHOLDING_API_BASE_URL + '/delete?taxWithholdingId=' + taxWithholdingId);
    }
}

export default new TaxWithholdingService()