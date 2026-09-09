import axios from 'axios';

const AUDITWORKPAPER_API_BASE_URL = "/AuditWorkpaper";

class AuditWorkpaperService {

    getAuditWorkpapers(){
        return axios.get(AUDITWORKPAPER_API_BASE_URL + '/' );
    }

    createAuditWorkpaper(auditWorkpaper){
        return axios.post(AUDITWORKPAPER_API_BASE_URL  + '/create', auditWorkpaper);
    }

    getAuditWorkpaperById(auditWorkpaperId){
        return axios.get(AUDITWORKPAPER_API_BASE_URL + '/load?auditWorkpaperId=' + auditWorkpaperId);
    }

    updateAuditWorkpaper(auditWorkpaper){
        return axios.put(AUDITWORKPAPER_API_BASE_URL + '/update', auditWorkpaper);
    }

    deleteAuditWorkpaper(auditWorkpaperId){
        return axios.delete(AUDITWORKPAPER_API_BASE_URL + '/delete?auditWorkpaperId=' + auditWorkpaperId);
    }
}

export default new AuditWorkpaperService()