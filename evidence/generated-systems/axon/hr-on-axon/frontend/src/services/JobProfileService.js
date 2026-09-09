import axios from 'axios';

const JOBPROFILE_API_BASE_URL = "/JobProfile";

class JobProfileService {

    getJobProfiles(){
        return axios.get(JOBPROFILE_API_BASE_URL + '/' );
    }

    createJobProfile(jobProfile){
        return axios.post(JOBPROFILE_API_BASE_URL  + '/create', jobProfile);
    }

    getJobProfileById(jobProfileId){
        return axios.get(JOBPROFILE_API_BASE_URL + '/load?jobProfileId=' + jobProfileId);
    }

    updateJobProfile(jobProfile){
        return axios.put(JOBPROFILE_API_BASE_URL + '/update', jobProfile);
    }

    deleteJobProfile(jobProfileId){
        return axios.delete(JOBPROFILE_API_BASE_URL + '/delete?jobProfileId=' + jobProfileId);
    }
}

export default new JobProfileService()