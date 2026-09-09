import axios from 'axios';

const TRACKINGPIXEL_API_BASE_URL = "/TrackingPixel";

class TrackingPixelService {

    getTrackingPixels(){
        return axios.get(TRACKINGPIXEL_API_BASE_URL + '/' );
    }

    createTrackingPixel(trackingPixel){
        return axios.post(TRACKINGPIXEL_API_BASE_URL  + '/create', trackingPixel);
    }

    getTrackingPixelById(trackingPixelId){
        return axios.get(TRACKINGPIXEL_API_BASE_URL + '/load?trackingPixelId=' + trackingPixelId);
    }

    updateTrackingPixel(trackingPixel){
        return axios.put(TRACKINGPIXEL_API_BASE_URL + '/update', trackingPixel);
    }

    deleteTrackingPixel(trackingPixelId){
        return axios.delete(TRACKINGPIXEL_API_BASE_URL + '/delete?trackingPixelId=' + trackingPixelId);
    }
}

export default new TrackingPixelService()