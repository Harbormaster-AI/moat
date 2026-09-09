import axios from 'axios';

const SERIALNUMBER_API_BASE_URL = "/SerialNumber";

class SerialNumberService {

    getSerialNumbers(){
        return axios.get(SERIALNUMBER_API_BASE_URL + '/' );
    }

    createSerialNumber(serialNumber){
        return axios.post(SERIALNUMBER_API_BASE_URL  + '/create', serialNumber);
    }

    getSerialNumberById(serialNumberId){
        return axios.get(SERIALNUMBER_API_BASE_URL + '/load?serialNumberId=' + serialNumberId);
    }

    updateSerialNumber(serialNumber){
        return axios.put(SERIALNUMBER_API_BASE_URL + '/update', serialNumber);
    }

    deleteSerialNumber(serialNumberId){
        return axios.delete(SERIALNUMBER_API_BASE_URL + '/delete?serialNumberId=' + serialNumberId);
    }
}

export default new SerialNumberService()