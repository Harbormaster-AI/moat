import axios from 'axios';

const USER_API_BASE_URL = "/User";

class UserService {

    getUsers(){
        return axios.get(USER_API_BASE_URL + '/' );
    }

    createUser(user){
        return axios.post(USER_API_BASE_URL  + '/create', user);
    }

    getUserById(userId){
        return axios.get(USER_API_BASE_URL + '/load?userId=' + userId);
    }

    updateUser(user){
        return axios.put(USER_API_BASE_URL + '/update', user);
    }

    deleteUser(userId){
        return axios.delete(USER_API_BASE_URL + '/delete?userId=' + userId);
    }
}

export default new UserService()