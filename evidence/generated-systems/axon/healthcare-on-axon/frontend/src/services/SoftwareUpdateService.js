import axios from 'axios';

const SOFTWAREUPDATE_API_BASE_URL = "/SoftwareUpdate";

class SoftwareUpdateService {

    getSoftwareUpdates(){
        return axios.get(SOFTWAREUPDATE_API_BASE_URL + '/' );
    }

    createSoftwareUpdate(softwareUpdate){
        return axios.post(SOFTWAREUPDATE_API_BASE_URL  + '/create', softwareUpdate);
    }

    getSoftwareUpdateById(softwareUpdateId){
        return axios.get(SOFTWAREUPDATE_API_BASE_URL + '/load?softwareUpdateId=' + softwareUpdateId);
    }

    updateSoftwareUpdate(softwareUpdate){
        return axios.put(SOFTWAREUPDATE_API_BASE_URL + '/update', softwareUpdate);
    }

    deleteSoftwareUpdate(softwareUpdateId){
        return axios.delete(SOFTWAREUPDATE_API_BASE_URL + '/delete?softwareUpdateId=' + softwareUpdateId);
    }
}

export default new SoftwareUpdateService()