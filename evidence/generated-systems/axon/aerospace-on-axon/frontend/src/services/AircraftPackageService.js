import axios from 'axios';

const AIRCRAFTPACKAGE_API_BASE_URL = "/AircraftPackage";

class AircraftPackageService {

    getAircraftPackages(){
        return axios.get(AIRCRAFTPACKAGE_API_BASE_URL + '/' );
    }

    createAircraftPackage(aircraftPackage){
        return axios.post(AIRCRAFTPACKAGE_API_BASE_URL  + '/create', aircraftPackage);
    }

    getAircraftPackageById(aircraftPackageId){
        return axios.get(AIRCRAFTPACKAGE_API_BASE_URL + '/load?aircraftPackageId=' + aircraftPackageId);
    }

    updateAircraftPackage(aircraftPackage){
        return axios.put(AIRCRAFTPACKAGE_API_BASE_URL + '/update', aircraftPackage);
    }

    deleteAircraftPackage(aircraftPackageId){
        return axios.delete(AIRCRAFTPACKAGE_API_BASE_URL + '/delete?aircraftPackageId=' + aircraftPackageId);
    }
}

export default new AircraftPackageService()