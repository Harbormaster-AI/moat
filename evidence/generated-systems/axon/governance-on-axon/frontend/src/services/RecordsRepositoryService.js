import axios from 'axios';

const RECORDSREPOSITORY_API_BASE_URL = "/RecordsRepository";

class RecordsRepositoryService {

    getRecordsRepositorys(){
        return axios.get(RECORDSREPOSITORY_API_BASE_URL + '/' );
    }

    createRecordsRepository(recordsRepository){
        return axios.post(RECORDSREPOSITORY_API_BASE_URL  + '/create', recordsRepository);
    }

    getRecordsRepositoryById(recordsRepositoryId){
        return axios.get(RECORDSREPOSITORY_API_BASE_URL + '/load?recordsRepositoryId=' + recordsRepositoryId);
    }

    updateRecordsRepository(recordsRepository){
        return axios.put(RECORDSREPOSITORY_API_BASE_URL + '/update', recordsRepository);
    }

    deleteRecordsRepository(recordsRepositoryId){
        return axios.delete(RECORDSREPOSITORY_API_BASE_URL + '/delete?recordsRepositoryId=' + recordsRepositoryId);
    }
}

export default new RecordsRepositoryService()