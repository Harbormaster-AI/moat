import axios from 'axios';

const TARGETINGPROFILE_API_BASE_URL = "/TargetingProfile";

class TargetingProfileService {

    getTargetingProfiles(){
        return axios.get(TARGETINGPROFILE_API_BASE_URL + '/' );
    }

    createTargetingProfile(targetingProfile){
        return axios.post(TARGETINGPROFILE_API_BASE_URL  + '/create', targetingProfile);
    }

    getTargetingProfileById(targetingProfileId){
        return axios.get(TARGETINGPROFILE_API_BASE_URL + '/load?targetingProfileId=' + targetingProfileId);
    }

    updateTargetingProfile(targetingProfile){
        return axios.put(TARGETINGPROFILE_API_BASE_URL + '/update', targetingProfile);
    }

    deleteTargetingProfile(targetingProfileId){
        return axios.delete(TARGETINGPROFILE_API_BASE_URL + '/delete?targetingProfileId=' + targetingProfileId);
    }
}

export default new TargetingProfileService()