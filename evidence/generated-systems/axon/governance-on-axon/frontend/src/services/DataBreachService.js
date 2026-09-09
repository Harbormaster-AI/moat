import axios from 'axios';

const DATABREACH_API_BASE_URL = "/DataBreach";

class DataBreachService {

    getDataBreachs(){
        return axios.get(DATABREACH_API_BASE_URL + '/' );
    }

    createDataBreach(dataBreach){
        return axios.post(DATABREACH_API_BASE_URL  + '/create', dataBreach);
    }

    getDataBreachById(dataBreachId){
        return axios.get(DATABREACH_API_BASE_URL + '/load?dataBreachId=' + dataBreachId);
    }

    updateDataBreach(dataBreach){
        return axios.put(DATABREACH_API_BASE_URL + '/update', dataBreach);
    }

    deleteDataBreach(dataBreachId){
        return axios.delete(DATABREACH_API_BASE_URL + '/delete?dataBreachId=' + dataBreachId);
    }
}

export default new DataBreachService()