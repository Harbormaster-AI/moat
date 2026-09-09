import axios from 'axios';

const UNDERWRITER_API_BASE_URL = "/Underwriter";

class UnderwriterService {

    getUnderwriters(){
        return axios.get(UNDERWRITER_API_BASE_URL + '/' );
    }

    createUnderwriter(underwriter){
        return axios.post(UNDERWRITER_API_BASE_URL  + '/create', underwriter);
    }

    getUnderwriterById(underwriterId){
        return axios.get(UNDERWRITER_API_BASE_URL + '/load?underwriterId=' + underwriterId);
    }

    updateUnderwriter(underwriter){
        return axios.put(UNDERWRITER_API_BASE_URL + '/update', underwriter);
    }

    deleteUnderwriter(underwriterId){
        return axios.delete(UNDERWRITER_API_BASE_URL + '/delete?underwriterId=' + underwriterId);
    }
}

export default new UnderwriterService()