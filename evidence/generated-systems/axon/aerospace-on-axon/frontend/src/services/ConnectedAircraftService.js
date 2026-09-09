import axios from 'axios';

const CONNECTEDAIRCRAFT_API_BASE_URL = "/ConnectedAircraft";

class ConnectedAircraftService {

    getConnectedAircrafts(){
        return axios.get(CONNECTEDAIRCRAFT_API_BASE_URL + '/' );
    }

    createConnectedAircraft(connectedAircraft){
        return axios.post(CONNECTEDAIRCRAFT_API_BASE_URL  + '/create', connectedAircraft);
    }

    getConnectedAircraftById(connectedAircraftId){
        return axios.get(CONNECTEDAIRCRAFT_API_BASE_URL + '/load?connectedAircraftId=' + connectedAircraftId);
    }

    updateConnectedAircraft(connectedAircraft){
        return axios.put(CONNECTEDAIRCRAFT_API_BASE_URL + '/update', connectedAircraft);
    }

    deleteConnectedAircraft(connectedAircraftId){
        return axios.delete(CONNECTEDAIRCRAFT_API_BASE_URL + '/delete?connectedAircraftId=' + connectedAircraftId);
    }
}

export default new ConnectedAircraftService()