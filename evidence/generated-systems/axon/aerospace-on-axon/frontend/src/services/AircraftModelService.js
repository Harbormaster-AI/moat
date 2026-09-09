import axios from 'axios';

const AIRCRAFTMODEL_API_BASE_URL = "/AircraftModel";

class AircraftModelService {

    getAircraftModels(){
        return axios.get(AIRCRAFTMODEL_API_BASE_URL + '/' );
    }

    createAircraftModel(aircraftModel){
        return axios.post(AIRCRAFTMODEL_API_BASE_URL  + '/create', aircraftModel);
    }

    getAircraftModelById(aircraftModelId){
        return axios.get(AIRCRAFTMODEL_API_BASE_URL + '/load?aircraftModelId=' + aircraftModelId);
    }

    updateAircraftModel(aircraftModel){
        return axios.put(AIRCRAFTMODEL_API_BASE_URL + '/update', aircraftModel);
    }

    deleteAircraftModel(aircraftModelId){
        return axios.delete(AIRCRAFTMODEL_API_BASE_URL + '/delete?aircraftModelId=' + aircraftModelId);
    }
}

export default new AircraftModelService()