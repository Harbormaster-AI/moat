import axios from 'axios';

const JOBAPPLICATION_API_BASE_URL = "/JobApplication";

class JobApplicationService {

    getJobApplications(){
        return axios.get(JOBAPPLICATION_API_BASE_URL + '/' );
    }

    createJobApplication(jobApplication){
        return axios.post(JOBAPPLICATION_API_BASE_URL  + '/create', jobApplication);
    }

    getJobApplicationById(jobApplicationId){
        return axios.get(JOBAPPLICATION_API_BASE_URL + '/load?jobApplicationId=' + jobApplicationId);
    }

    updateJobApplication(jobApplication){
        return axios.put(JOBAPPLICATION_API_BASE_URL + '/update', jobApplication);
    }

    deleteJobApplication(jobApplicationId){
        return axios.delete(JOBAPPLICATION_API_BASE_URL + '/delete?jobApplicationId=' + jobApplicationId);
    }
}

export default new JobApplicationService()