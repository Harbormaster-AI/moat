import axios from 'axios';

const IMAGINGCENTER_API_BASE_URL = "/ImagingCenter";

class ImagingCenterService {

    getImagingCenters(){
        return axios.get(IMAGINGCENTER_API_BASE_URL + '/' );
    }

    createImagingCenter(imagingCenter){
        return axios.post(IMAGINGCENTER_API_BASE_URL  + '/create', imagingCenter);
    }

    getImagingCenterById(imagingCenterId){
        return axios.get(IMAGINGCENTER_API_BASE_URL + '/load?imagingCenterId=' + imagingCenterId);
    }

    updateImagingCenter(imagingCenter){
        return axios.put(IMAGINGCENTER_API_BASE_URL + '/update', imagingCenter);
    }

    deleteImagingCenter(imagingCenterId){
        return axios.delete(IMAGINGCENTER_API_BASE_URL + '/delete?imagingCenterId=' + imagingCenterId);
    }
}

export default new ImagingCenterService()