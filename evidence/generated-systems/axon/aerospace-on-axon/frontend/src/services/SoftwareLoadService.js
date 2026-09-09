import axios from 'axios';

const SOFTWARELOAD_API_BASE_URL = "/SoftwareLoad";

class SoftwareLoadService {

    getSoftwareLoads(){
        return axios.get(SOFTWARELOAD_API_BASE_URL + '/' );
    }

    createSoftwareLoad(softwareLoad){
        return axios.post(SOFTWARELOAD_API_BASE_URL  + '/create', softwareLoad);
    }

    getSoftwareLoadById(softwareLoadId){
        return axios.get(SOFTWARELOAD_API_BASE_URL + '/load?softwareLoadId=' + softwareLoadId);
    }

    updateSoftwareLoad(softwareLoad){
        return axios.put(SOFTWARELOAD_API_BASE_URL + '/update', softwareLoad);
    }

    deleteSoftwareLoad(softwareLoadId){
        return axios.delete(SOFTWARELOAD_API_BASE_URL + '/delete?softwareLoadId=' + softwareLoadId);
    }
}

export default new SoftwareLoadService()