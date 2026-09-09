import axios from 'axios';

const GEOREGION_API_BASE_URL = "/GeoRegion";

class GeoRegionService {

    getGeoRegions(){
        return axios.get(GEOREGION_API_BASE_URL + '/' );
    }

    createGeoRegion(geoRegion){
        return axios.post(GEOREGION_API_BASE_URL  + '/create', geoRegion);
    }

    getGeoRegionById(geoRegionId){
        return axios.get(GEOREGION_API_BASE_URL + '/load?geoRegionId=' + geoRegionId);
    }

    updateGeoRegion(geoRegion){
        return axios.put(GEOREGION_API_BASE_URL + '/update', geoRegion);
    }

    deleteGeoRegion(geoRegionId){
        return axios.delete(GEOREGION_API_BASE_URL + '/delete?geoRegionId=' + geoRegionId);
    }
}

export default new GeoRegionService()