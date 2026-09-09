import axios from 'axios';

const INSUREDOBJECT_API_BASE_URL = "/InsuredObject";

class InsuredObjectService {

    getInsuredObjects(){
        return axios.get(INSUREDOBJECT_API_BASE_URL + '/' );
    }

    createInsuredObject(insuredObject){
        return axios.post(INSUREDOBJECT_API_BASE_URL  + '/create', insuredObject);
    }

    getInsuredObjectById(insuredObjectId){
        return axios.get(INSUREDOBJECT_API_BASE_URL + '/load?insuredObjectId=' + insuredObjectId);
    }

    updateInsuredObject(insuredObject){
        return axios.put(INSUREDOBJECT_API_BASE_URL + '/update', insuredObject);
    }

    deleteInsuredObject(insuredObjectId){
        return axios.delete(INSUREDOBJECT_API_BASE_URL + '/delete?insuredObjectId=' + insuredObjectId);
    }
}

export default new InsuredObjectService()