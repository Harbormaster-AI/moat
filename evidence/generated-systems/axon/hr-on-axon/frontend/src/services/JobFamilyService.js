import axios from 'axios';

const JOBFAMILY_API_BASE_URL = "/JobFamily";

class JobFamilyService {

    getJobFamilys(){
        return axios.get(JOBFAMILY_API_BASE_URL + '/' );
    }

    createJobFamily(jobFamily){
        return axios.post(JOBFAMILY_API_BASE_URL  + '/create', jobFamily);
    }

    getJobFamilyById(jobFamilyId){
        return axios.get(JOBFAMILY_API_BASE_URL + '/load?jobFamilyId=' + jobFamilyId);
    }

    updateJobFamily(jobFamily){
        return axios.put(JOBFAMILY_API_BASE_URL + '/update', jobFamily);
    }

    deleteJobFamily(jobFamilyId){
        return axios.delete(JOBFAMILY_API_BASE_URL + '/delete?jobFamilyId=' + jobFamilyId);
    }
}

export default new JobFamilyService()