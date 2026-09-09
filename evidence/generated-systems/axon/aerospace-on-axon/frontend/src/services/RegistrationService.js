import axios from 'axios';

const REGISTRATION_API_BASE_URL = "/Registration";

class RegistrationService {

    getRegistrations(){
        return axios.get(REGISTRATION_API_BASE_URL + '/' );
    }

    createRegistration(registration){
        return axios.post(REGISTRATION_API_BASE_URL  + '/create', registration);
    }

    getRegistrationById(registrationId){
        return axios.get(REGISTRATION_API_BASE_URL + '/load?registrationId=' + registrationId);
    }

    updateRegistration(registration){
        return axios.put(REGISTRATION_API_BASE_URL + '/update', registration);
    }

    deleteRegistration(registrationId){
        return axios.delete(REGISTRATION_API_BASE_URL + '/delete?registrationId=' + registrationId);
    }
}

export default new RegistrationService()