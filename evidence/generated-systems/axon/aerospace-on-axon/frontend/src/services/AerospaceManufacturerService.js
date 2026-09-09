import axios from 'axios';

const AEROSPACEMANUFACTURER_API_BASE_URL = "/AerospaceManufacturer";

class AerospaceManufacturerService {

    getAerospaceManufacturers(){
        return axios.get(AEROSPACEMANUFACTURER_API_BASE_URL + '/' );
    }

    createAerospaceManufacturer(aerospaceManufacturer){
        return axios.post(AEROSPACEMANUFACTURER_API_BASE_URL  + '/create', aerospaceManufacturer);
    }

    getAerospaceManufacturerById(aerospaceManufacturerId){
        return axios.get(AEROSPACEMANUFACTURER_API_BASE_URL + '/load?aerospaceManufacturerId=' + aerospaceManufacturerId);
    }

    updateAerospaceManufacturer(aerospaceManufacturer){
        return axios.put(AEROSPACEMANUFACTURER_API_BASE_URL + '/update', aerospaceManufacturer);
    }

    deleteAerospaceManufacturer(aerospaceManufacturerId){
        return axios.delete(AEROSPACEMANUFACTURER_API_BASE_URL + '/delete?aerospaceManufacturerId=' + aerospaceManufacturerId);
    }
}

export default new AerospaceManufacturerService()