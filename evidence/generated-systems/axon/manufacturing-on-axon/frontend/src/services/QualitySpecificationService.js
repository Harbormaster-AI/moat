import axios from 'axios';

const QUALITYSPECIFICATION_API_BASE_URL = "/QualitySpecification";

class QualitySpecificationService {

    getQualitySpecifications(){
        return axios.get(QUALITYSPECIFICATION_API_BASE_URL + '/' );
    }

    createQualitySpecification(qualitySpecification){
        return axios.post(QUALITYSPECIFICATION_API_BASE_URL  + '/create', qualitySpecification);
    }

    getQualitySpecificationById(qualitySpecificationId){
        return axios.get(QUALITYSPECIFICATION_API_BASE_URL + '/load?qualitySpecificationId=' + qualitySpecificationId);
    }

    updateQualitySpecification(qualitySpecification){
        return axios.put(QUALITYSPECIFICATION_API_BASE_URL + '/update', qualitySpecification);
    }

    deleteQualitySpecification(qualitySpecificationId){
        return axios.delete(QUALITYSPECIFICATION_API_BASE_URL + '/delete?qualitySpecificationId=' + qualitySpecificationId);
    }
}

export default new QualitySpecificationService()