import axios from 'axios';

const MEDICALSUPPLIER_API_BASE_URL = "/MedicalSupplier";

class MedicalSupplierService {

    getMedicalSuppliers(){
        return axios.get(MEDICALSUPPLIER_API_BASE_URL + '/' );
    }

    createMedicalSupplier(medicalSupplier){
        return axios.post(MEDICALSUPPLIER_API_BASE_URL  + '/create', medicalSupplier);
    }

    getMedicalSupplierById(medicalSupplierId){
        return axios.get(MEDICALSUPPLIER_API_BASE_URL + '/load?medicalSupplierId=' + medicalSupplierId);
    }

    updateMedicalSupplier(medicalSupplier){
        return axios.put(MEDICALSUPPLIER_API_BASE_URL + '/update', medicalSupplier);
    }

    deleteMedicalSupplier(medicalSupplierId){
        return axios.delete(MEDICALSUPPLIER_API_BASE_URL + '/delete?medicalSupplierId=' + medicalSupplierId);
    }
}

export default new MedicalSupplierService()