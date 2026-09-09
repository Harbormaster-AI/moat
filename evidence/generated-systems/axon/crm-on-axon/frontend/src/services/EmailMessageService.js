import axios from 'axios';

const EMAILMESSAGE_API_BASE_URL = "/EmailMessage";

class EmailMessageService {

    getEmailMessages(){
        return axios.get(EMAILMESSAGE_API_BASE_URL + '/' );
    }

    createEmailMessage(emailMessage){
        return axios.post(EMAILMESSAGE_API_BASE_URL  + '/create', emailMessage);
    }

    getEmailMessageById(emailMessageId){
        return axios.get(EMAILMESSAGE_API_BASE_URL + '/load?emailMessageId=' + emailMessageId);
    }

    updateEmailMessage(emailMessage){
        return axios.put(EMAILMESSAGE_API_BASE_URL + '/update', emailMessage);
    }

    deleteEmailMessage(emailMessageId){
        return axios.delete(EMAILMESSAGE_API_BASE_URL + '/delete?emailMessageId=' + emailMessageId);
    }
}

export default new EmailMessageService()