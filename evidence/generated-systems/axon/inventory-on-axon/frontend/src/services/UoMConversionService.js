import axios from 'axios';

const UOMCONVERSION_API_BASE_URL = "/UoMConversion";

class UoMConversionService {

    getUoMConversions(){
        return axios.get(UOMCONVERSION_API_BASE_URL + '/' );
    }

    createUoMConversion(uoMConversion){
        return axios.post(UOMCONVERSION_API_BASE_URL  + '/create', uoMConversion);
    }

    getUoMConversionById(uoMConversionId){
        return axios.get(UOMCONVERSION_API_BASE_URL + '/load?uoMConversionId=' + uoMConversionId);
    }

    updateUoMConversion(uoMConversion){
        return axios.put(UOMCONVERSION_API_BASE_URL + '/update', uoMConversion);
    }

    deleteUoMConversion(uoMConversionId){
        return axios.delete(UOMCONVERSION_API_BASE_URL + '/delete?uoMConversionId=' + uoMConversionId);
    }
}

export default new UoMConversionService()