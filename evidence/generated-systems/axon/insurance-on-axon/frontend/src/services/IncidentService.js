import axios from 'axios';

const INCIDENT_API_BASE_URL = "/Incident";

class IncidentService {

    getIncidents(){
        return axios.get(INCIDENT_API_BASE_URL + '/' );
    }

    createIncident(incident){
        return axios.post(INCIDENT_API_BASE_URL  + '/create', incident);
    }

    getIncidentById(incidentId){
        return axios.get(INCIDENT_API_BASE_URL + '/load?incidentId=' + incidentId);
    }

    updateIncident(incident){
        return axios.put(INCIDENT_API_BASE_URL + '/update', incident);
    }

    deleteIncident(incidentId){
        return axios.delete(INCIDENT_API_BASE_URL + '/delete?incidentId=' + incidentId);
    }
}

export default new IncidentService()