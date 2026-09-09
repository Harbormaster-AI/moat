import axios from 'axios';

const WORKCENTER_API_BASE_URL = "/WorkCenter";

class WorkCenterService {

    getWorkCenters(){
        return axios.get(WORKCENTER_API_BASE_URL + '/' );
    }

    createWorkCenter(workCenter){
        return axios.post(WORKCENTER_API_BASE_URL  + '/create', workCenter);
    }

    getWorkCenterById(workCenterId){
        return axios.get(WORKCENTER_API_BASE_URL + '/load?workCenterId=' + workCenterId);
    }

    updateWorkCenter(workCenter){
        return axios.put(WORKCENTER_API_BASE_URL + '/update', workCenter);
    }

    deleteWorkCenter(workCenterId){
        return axios.delete(WORKCENTER_API_BASE_URL + '/delete?workCenterId=' + workCenterId);
    }
}

export default new WorkCenterService()