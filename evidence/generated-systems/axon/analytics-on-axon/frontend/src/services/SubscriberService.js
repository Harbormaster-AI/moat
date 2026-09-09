import axios from 'axios';

const SUBSCRIBER_API_BASE_URL = "/Subscriber";

class SubscriberService {

    getSubscribers(){
        return axios.get(SUBSCRIBER_API_BASE_URL + '/' );
    }

    createSubscriber(subscriber){
        return axios.post(SUBSCRIBER_API_BASE_URL  + '/create', subscriber);
    }

    getSubscriberById(subscriberId){
        return axios.get(SUBSCRIBER_API_BASE_URL + '/load?subscriberId=' + subscriberId);
    }

    updateSubscriber(subscriber){
        return axios.put(SUBSCRIBER_API_BASE_URL + '/update', subscriber);
    }

    deleteSubscriber(subscriberId){
        return axios.delete(SUBSCRIBER_API_BASE_URL + '/delete?subscriberId=' + subscriberId);
    }
}

export default new SubscriberService()