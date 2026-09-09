import axios from 'axios';

const SERVICEBULLETIN_API_BASE_URL = "/ServiceBulletin";

class ServiceBulletinService {

    getServiceBulletins(){
        return axios.get(SERVICEBULLETIN_API_BASE_URL + '/' );
    }

    createServiceBulletin(serviceBulletin){
        return axios.post(SERVICEBULLETIN_API_BASE_URL  + '/create', serviceBulletin);
    }

    getServiceBulletinById(serviceBulletinId){
        return axios.get(SERVICEBULLETIN_API_BASE_URL + '/load?serviceBulletinId=' + serviceBulletinId);
    }

    updateServiceBulletin(serviceBulletin){
        return axios.put(SERVICEBULLETIN_API_BASE_URL + '/update', serviceBulletin);
    }

    deleteServiceBulletin(serviceBulletinId){
        return axios.delete(SERVICEBULLETIN_API_BASE_URL + '/delete?serviceBulletinId=' + serviceBulletinId);
    }
}

export default new ServiceBulletinService()