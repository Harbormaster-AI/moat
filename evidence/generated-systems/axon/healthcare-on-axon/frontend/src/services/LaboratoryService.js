import axios from 'axios';

const LABORATORY_API_BASE_URL = "/Laboratory";

class LaboratoryService {

    getLaboratorys(){
        return axios.get(LABORATORY_API_BASE_URL + '/' );
    }

    createLaboratory(laboratory){
        return axios.post(LABORATORY_API_BASE_URL  + '/create', laboratory);
    }

    getLaboratoryById(laboratoryId){
        return axios.get(LABORATORY_API_BASE_URL + '/load?laboratoryId=' + laboratoryId);
    }

    updateLaboratory(laboratory){
        return axios.put(LABORATORY_API_BASE_URL + '/update', laboratory);
    }

    deleteLaboratory(laboratoryId){
        return axios.delete(LABORATORY_API_BASE_URL + '/delete?laboratoryId=' + laboratoryId);
    }
}

export default new LaboratoryService()