import axios from 'axios';

const PRIVACYNOTICE_API_BASE_URL = "/PrivacyNotice";

class PrivacyNoticeService {

    getPrivacyNotices(){
        return axios.get(PRIVACYNOTICE_API_BASE_URL + '/' );
    }

    createPrivacyNotice(privacyNotice){
        return axios.post(PRIVACYNOTICE_API_BASE_URL  + '/create', privacyNotice);
    }

    getPrivacyNoticeById(privacyNoticeId){
        return axios.get(PRIVACYNOTICE_API_BASE_URL + '/load?privacyNoticeId=' + privacyNoticeId);
    }

    updatePrivacyNotice(privacyNotice){
        return axios.put(PRIVACYNOTICE_API_BASE_URL + '/update', privacyNotice);
    }

    deletePrivacyNotice(privacyNoticeId){
        return axios.delete(PRIVACYNOTICE_API_BASE_URL + '/delete?privacyNoticeId=' + privacyNoticeId);
    }
}

export default new PrivacyNoticeService()