import axios from 'axios';

const CERTIFICATION_API_BASE_URL = "/Certification";

class CertificationService {

    getCertifications(){
        return axios.get(CERTIFICATION_API_BASE_URL + '/' );
    }

    createCertification(certification){
        return axios.post(CERTIFICATION_API_BASE_URL  + '/create', certification);
    }

    getCertificationById(certificationId){
        return axios.get(CERTIFICATION_API_BASE_URL + '/load?certificationId=' + certificationId);
    }

    updateCertification(certification){
        return axios.put(CERTIFICATION_API_BASE_URL + '/update', certification);
    }

    deleteCertification(certificationId){
        return axios.delete(CERTIFICATION_API_BASE_URL + '/delete?certificationId=' + certificationId);
    }
}

export default new CertificationService()