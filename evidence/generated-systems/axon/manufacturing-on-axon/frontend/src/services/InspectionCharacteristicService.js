import axios from 'axios';

const INSPECTIONCHARACTERISTIC_API_BASE_URL = "/InspectionCharacteristic";

class InspectionCharacteristicService {

    getInspectionCharacteristics(){
        return axios.get(INSPECTIONCHARACTERISTIC_API_BASE_URL + '/' );
    }

    createInspectionCharacteristic(inspectionCharacteristic){
        return axios.post(INSPECTIONCHARACTERISTIC_API_BASE_URL  + '/create', inspectionCharacteristic);
    }

    getInspectionCharacteristicById(inspectionCharacteristicId){
        return axios.get(INSPECTIONCHARACTERISTIC_API_BASE_URL + '/load?inspectionCharacteristicId=' + inspectionCharacteristicId);
    }

    updateInspectionCharacteristic(inspectionCharacteristic){
        return axios.put(INSPECTIONCHARACTERISTIC_API_BASE_URL + '/update', inspectionCharacteristic);
    }

    deleteInspectionCharacteristic(inspectionCharacteristicId){
        return axios.delete(INSPECTIONCHARACTERISTIC_API_BASE_URL + '/delete?inspectionCharacteristicId=' + inspectionCharacteristicId);
    }
}

export default new InspectionCharacteristicService()