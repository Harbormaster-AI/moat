import axios from 'axios';

const STORAGELOCATION_API_BASE_URL = "/StorageLocation";

class StorageLocationService {

    getStorageLocations(){
        return axios.get(STORAGELOCATION_API_BASE_URL + '/' );
    }

    createStorageLocation(storageLocation){
        return axios.post(STORAGELOCATION_API_BASE_URL  + '/create', storageLocation);
    }

    getStorageLocationById(storageLocationId){
        return axios.get(STORAGELOCATION_API_BASE_URL + '/load?storageLocationId=' + storageLocationId);
    }

    updateStorageLocation(storageLocation){
        return axios.put(STORAGELOCATION_API_BASE_URL + '/update', storageLocation);
    }

    deleteStorageLocation(storageLocationId){
        return axios.delete(STORAGELOCATION_API_BASE_URL + '/delete?storageLocationId=' + storageLocationId);
    }
}

export default new StorageLocationService()