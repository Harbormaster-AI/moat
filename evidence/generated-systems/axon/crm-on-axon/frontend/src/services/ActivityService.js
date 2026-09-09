import axios from 'axios';

const ACTIVITY_API_BASE_URL = "/Activity";

class ActivityService {

    getActivitys(){
        return axios.get(ACTIVITY_API_BASE_URL + '/' );
    }

    createActivity(activity){
        return axios.post(ACTIVITY_API_BASE_URL  + '/create', activity);
    }

    getActivityById(activityId){
        return axios.get(ACTIVITY_API_BASE_URL + '/load?activityId=' + activityId);
    }

    updateActivity(activity){
        return axios.put(ACTIVITY_API_BASE_URL + '/update', activity);
    }

    deleteActivity(activityId){
        return axios.delete(ACTIVITY_API_BASE_URL + '/delete?activityId=' + activityId);
    }
}

export default new ActivityService()