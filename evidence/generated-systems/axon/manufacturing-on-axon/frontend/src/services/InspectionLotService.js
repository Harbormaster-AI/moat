import axios from 'axios';

const INSPECTIONLOT_API_BASE_URL = "/InspectionLot";

class InspectionLotService {

    getInspectionLots(){
        return axios.get(INSPECTIONLOT_API_BASE_URL + '/' );
    }

    createInspectionLot(inspectionLot){
        return axios.post(INSPECTIONLOT_API_BASE_URL  + '/create', inspectionLot);
    }

    getInspectionLotById(inspectionLotId){
        return axios.get(INSPECTIONLOT_API_BASE_URL + '/load?inspectionLotId=' + inspectionLotId);
    }

    updateInspectionLot(inspectionLot){
        return axios.put(INSPECTIONLOT_API_BASE_URL + '/update', inspectionLot);
    }

    deleteInspectionLot(inspectionLotId){
        return axios.delete(INSPECTIONLOT_API_BASE_URL + '/delete?inspectionLotId=' + inspectionLotId);
    }
}

export default new InspectionLotService()