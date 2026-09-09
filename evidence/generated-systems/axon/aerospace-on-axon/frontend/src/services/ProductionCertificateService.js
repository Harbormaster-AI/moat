import axios from 'axios';

const PRODUCTIONCERTIFICATE_API_BASE_URL = "/ProductionCertificate";

class ProductionCertificateService {

    getProductionCertificates(){
        return axios.get(PRODUCTIONCERTIFICATE_API_BASE_URL + '/' );
    }

    createProductionCertificate(productionCertificate){
        return axios.post(PRODUCTIONCERTIFICATE_API_BASE_URL  + '/create', productionCertificate);
    }

    getProductionCertificateById(productionCertificateId){
        return axios.get(PRODUCTIONCERTIFICATE_API_BASE_URL + '/load?productionCertificateId=' + productionCertificateId);
    }

    updateProductionCertificate(productionCertificate){
        return axios.put(PRODUCTIONCERTIFICATE_API_BASE_URL + '/update', productionCertificate);
    }

    deleteProductionCertificate(productionCertificateId){
        return axios.delete(PRODUCTIONCERTIFICATE_API_BASE_URL + '/delete?productionCertificateId=' + productionCertificateId);
    }
}

export default new ProductionCertificateService()