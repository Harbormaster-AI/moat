import axios from 'axios';

const SUPPLIER_API_BASE_URL = "/Supplier";

class SupplierService {

    getSuppliers(){
        return axios.get(SUPPLIER_API_BASE_URL + '/' );
    }

    createSupplier(supplier){
        return axios.post(SUPPLIER_API_BASE_URL  + '/create', supplier);
    }

    getSupplierById(supplierId){
        return axios.get(SUPPLIER_API_BASE_URL + '/load?supplierId=' + supplierId);
    }

    updateSupplier(supplier){
        return axios.put(SUPPLIER_API_BASE_URL + '/update', supplier);
    }

    deleteSupplier(supplierId){
        return axios.delete(SUPPLIER_API_BASE_URL + '/delete?supplierId=' + supplierId);
    }
}

export default new SupplierService()