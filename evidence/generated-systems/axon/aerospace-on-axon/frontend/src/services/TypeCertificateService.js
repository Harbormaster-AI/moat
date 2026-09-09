import axios from 'axios';

const TYPECERTIFICATE_API_BASE_URL = "/TypeCertificate";

class TypeCertificateService {

    getTypeCertificates(){
        return axios.get(TYPECERTIFICATE_API_BASE_URL + '/' );
    }

    createTypeCertificate(typeCertificate){
        return axios.post(TYPECERTIFICATE_API_BASE_URL  + '/create', typeCertificate);
    }

    getTypeCertificateById(typeCertificateId){
        return axios.get(TYPECERTIFICATE_API_BASE_URL + '/load?typeCertificateId=' + typeCertificateId);
    }

    updateTypeCertificate(typeCertificate){
        return axios.put(TYPECERTIFICATE_API_BASE_URL + '/update', typeCertificate);
    }

    deleteTypeCertificate(typeCertificateId){
        return axios.delete(TYPECERTIFICATE_API_BASE_URL + '/delete?typeCertificateId=' + typeCertificateId);
    }
}

export default new TypeCertificateService()