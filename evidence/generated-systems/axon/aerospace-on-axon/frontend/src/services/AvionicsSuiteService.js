import axios from 'axios';

const AVIONICSSUITE_API_BASE_URL = "/AvionicsSuite";

class AvionicsSuiteService {

    getAvionicsSuites(){
        return axios.get(AVIONICSSUITE_API_BASE_URL + '/' );
    }

    createAvionicsSuite(avionicsSuite){
        return axios.post(AVIONICSSUITE_API_BASE_URL  + '/create', avionicsSuite);
    }

    getAvionicsSuiteById(avionicsSuiteId){
        return axios.get(AVIONICSSUITE_API_BASE_URL + '/load?avionicsSuiteId=' + avionicsSuiteId);
    }

    updateAvionicsSuite(avionicsSuite){
        return axios.put(AVIONICSSUITE_API_BASE_URL + '/update', avionicsSuite);
    }

    deleteAvionicsSuite(avionicsSuiteId){
        return axios.delete(AVIONICSSUITE_API_BASE_URL + '/delete?avionicsSuiteId=' + avionicsSuiteId);
    }
}

export default new AvionicsSuiteService()