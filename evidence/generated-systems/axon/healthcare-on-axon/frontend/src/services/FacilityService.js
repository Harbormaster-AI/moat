import axios from 'axios';

const FACILITY_API_BASE_URL = "/Facility";

class FacilityService {

    getFacilitys(){
        return axios.get(FACILITY_API_BASE_URL + '/' );
    }

    createFacility(facility){
        return axios.post(FACILITY_API_BASE_URL  + '/create', facility);
    }

    getFacilityById(facilityId){
        return axios.get(FACILITY_API_BASE_URL + '/load?facilityId=' + facilityId);
    }

    updateFacility(facility){
        return axios.put(FACILITY_API_BASE_URL + '/update', facility);
    }

    deleteFacility(facilityId){
        return axios.delete(FACILITY_API_BASE_URL + '/delete?facilityId=' + facilityId);
    }
}

export default new FacilityService()