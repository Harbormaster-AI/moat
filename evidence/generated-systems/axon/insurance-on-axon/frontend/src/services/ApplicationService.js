import axios from 'axios';

const APPLICATION_API_BASE_URL = "/Application";

class ApplicationService {

    getApplications(){
        return axios.get(APPLICATION_API_BASE_URL + '/' );
    }

    createApplication(application){
        return axios.post(APPLICATION_API_BASE_URL  + '/create', application);
    }

    getApplicationById(applicationId){
        return axios.get(APPLICATION_API_BASE_URL + '/load?applicationId=' + applicationId);
    }

    updateApplication(application){
        return axios.put(APPLICATION_API_BASE_URL + '/update', application);
    }

    deleteApplication(applicationId){
        return axios.delete(APPLICATION_API_BASE_URL + '/delete?applicationId=' + applicationId);
    }
}

export default new ApplicationService()