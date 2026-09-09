import axios from 'axios';

const CONTACT_API_BASE_URL = "/Contact";

class ContactService {

    getContacts(){
        return axios.get(CONTACT_API_BASE_URL + '/' );
    }

    createContact(contact){
        return axios.post(CONTACT_API_BASE_URL  + '/create', contact);
    }

    getContactById(contactId){
        return axios.get(CONTACT_API_BASE_URL + '/load?contactId=' + contactId);
    }

    updateContact(contact){
        return axios.put(CONTACT_API_BASE_URL + '/update', contact);
    }

    deleteContact(contactId){
        return axios.delete(CONTACT_API_BASE_URL + '/delete?contactId=' + contactId);
    }
}

export default new ContactService()