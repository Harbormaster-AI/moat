import axios from 'axios';

const COMPENSATIONPACKAGE_API_BASE_URL = "/CompensationPackage";

class CompensationPackageService {

    getCompensationPackages(){
        return axios.get(COMPENSATIONPACKAGE_API_BASE_URL + '/' );
    }

    createCompensationPackage(compensationPackage){
        return axios.post(COMPENSATIONPACKAGE_API_BASE_URL  + '/create', compensationPackage);
    }

    getCompensationPackageById(compensationPackageId){
        return axios.get(COMPENSATIONPACKAGE_API_BASE_URL + '/load?compensationPackageId=' + compensationPackageId);
    }

    updateCompensationPackage(compensationPackage){
        return axios.put(COMPENSATIONPACKAGE_API_BASE_URL + '/update', compensationPackage);
    }

    deleteCompensationPackage(compensationPackageId){
        return axios.delete(COMPENSATIONPACKAGE_API_BASE_URL + '/delete?compensationPackageId=' + compensationPackageId);
    }
}

export default new CompensationPackageService()