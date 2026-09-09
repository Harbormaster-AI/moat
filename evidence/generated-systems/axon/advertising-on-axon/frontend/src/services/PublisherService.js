import axios from 'axios';

const PUBLISHER_API_BASE_URL = "/Publisher";

class PublisherService {

    getPublishers(){
        return axios.get(PUBLISHER_API_BASE_URL + '/' );
    }

    createPublisher(publisher){
        return axios.post(PUBLISHER_API_BASE_URL  + '/create', publisher);
    }

    getPublisherById(publisherId){
        return axios.get(PUBLISHER_API_BASE_URL + '/load?publisherId=' + publisherId);
    }

    updatePublisher(publisher){
        return axios.put(PUBLISHER_API_BASE_URL + '/update', publisher);
    }

    deletePublisher(publisherId){
        return axios.delete(PUBLISHER_API_BASE_URL + '/delete?publisherId=' + publisherId);
    }
}

export default new PublisherService()