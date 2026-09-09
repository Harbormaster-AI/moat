import axios from 'axios';

const SUBSCRIPTION_API_BASE_URL = "/Subscription";

class SubscriptionService {

    getSubscriptions(){
        return axios.get(SUBSCRIPTION_API_BASE_URL + '/' );
    }

    createSubscription(subscription){
        return axios.post(SUBSCRIPTION_API_BASE_URL  + '/create', subscription);
    }

    getSubscriptionById(subscriptionId){
        return axios.get(SUBSCRIPTION_API_BASE_URL + '/load?subscriptionId=' + subscriptionId);
    }

    updateSubscription(subscription){
        return axios.put(SUBSCRIPTION_API_BASE_URL + '/update', subscription);
    }

    deleteSubscription(subscriptionId){
        return axios.delete(SUBSCRIPTION_API_BASE_URL + '/delete?subscriptionId=' + subscriptionId);
    }
}

export default new SubscriptionService()