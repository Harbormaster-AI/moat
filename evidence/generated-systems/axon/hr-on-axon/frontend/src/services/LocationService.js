import axios from 'axios';

const LOCATION_API_BASE_URL = "/Location";

class LocationService {

    getLocations(){
        return axios.get(LOCATION_API_BASE_URL + '/' );
    }

    createLocation(location){
        return axios.post(LOCATION_API_BASE_URL  + '/create', location);
    }

    getLocationById(locationId){
        return axios.get(LOCATION_API_BASE_URL + '/load?locationId=' + locationId);
    }

    updateLocation(location){
        return axios.put(LOCATION_API_BASE_URL + '/update', location);
    }

    deleteLocation(locationId){
        return axios.delete(LOCATION_API_BASE_URL + '/delete?locationId=' + locationId);
    }
}

export default new LocationService()