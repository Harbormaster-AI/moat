import axios from 'axios';

const AUTHORIZATION_API_BASE_URL = "/Authorization";

class AuthorizationService {

    getAuthorizations(){
        return axios.get(AUTHORIZATION_API_BASE_URL + '/' );
    }

    createAuthorization(authorization){
        return axios.post(AUTHORIZATION_API_BASE_URL  + '/create', authorization);
    }

    getAuthorizationById(authorizationId){
        return axios.get(AUTHORIZATION_API_BASE_URL + '/load?authorizationId=' + authorizationId);
    }

    updateAuthorization(authorization){
        return axios.put(AUTHORIZATION_API_BASE_URL + '/update', authorization);
    }

    deleteAuthorization(authorizationId){
        return axios.delete(AUTHORIZATION_API_BASE_URL + '/delete?authorizationId=' + authorizationId);
    }
}

export default new AuthorizationService()