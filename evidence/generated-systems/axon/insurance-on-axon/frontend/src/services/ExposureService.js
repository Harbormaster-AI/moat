import axios from 'axios';

const EXPOSURE_API_BASE_URL = "/Exposure";

class ExposureService {

    getExposures(){
        return axios.get(EXPOSURE_API_BASE_URL + '/' );
    }

    createExposure(exposure){
        return axios.post(EXPOSURE_API_BASE_URL  + '/create', exposure);
    }

    getExposureById(exposureId){
        return axios.get(EXPOSURE_API_BASE_URL + '/load?exposureId=' + exposureId);
    }

    updateExposure(exposure){
        return axios.put(EXPOSURE_API_BASE_URL + '/update', exposure);
    }

    deleteExposure(exposureId){
        return axios.delete(EXPOSURE_API_BASE_URL + '/delete?exposureId=' + exposureId);
    }
}

export default new ExposureService()