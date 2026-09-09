import axios from 'axios';

const CREATIVEFILE_API_BASE_URL = "/CreativeFile";

class CreativeFileService {

    getCreativeFiles(){
        return axios.get(CREATIVEFILE_API_BASE_URL + '/' );
    }

    createCreativeFile(creativeFile){
        return axios.post(CREATIVEFILE_API_BASE_URL  + '/create', creativeFile);
    }

    getCreativeFileById(creativeFileId){
        return axios.get(CREATIVEFILE_API_BASE_URL + '/load?creativeFileId=' + creativeFileId);
    }

    updateCreativeFile(creativeFile){
        return axios.put(CREATIVEFILE_API_BASE_URL + '/update', creativeFile);
    }

    deleteCreativeFile(creativeFileId){
        return axios.delete(CREATIVEFILE_API_BASE_URL + '/delete?creativeFileId=' + creativeFileId);
    }
}

export default new CreativeFileService()