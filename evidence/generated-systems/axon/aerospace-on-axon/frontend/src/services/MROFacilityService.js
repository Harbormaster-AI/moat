import axios from 'axios';

const MROFACILITY_API_BASE_URL = "/MROFacility";

class MROFacilityService {

    getMROFacilitys(){
        return axios.get(MROFACILITY_API_BASE_URL + '/' );
    }

    createMROFacility(mROFacility){
        return axios.post(MROFACILITY_API_BASE_URL  + '/create', mROFacility);
    }

    getMROFacilityById(mROFacilityId){
        return axios.get(MROFACILITY_API_BASE_URL + '/load?mROFacilityId=' + mROFacilityId);
    }

    updateMROFacility(mROFacility){
        return axios.put(MROFACILITY_API_BASE_URL + '/update', mROFacility);
    }

    deleteMROFacility(mROFacilityId){
        return axios.delete(MROFACILITY_API_BASE_URL + '/delete?mROFacilityId=' + mROFacilityId);
    }
}

export default new MROFacilityService()