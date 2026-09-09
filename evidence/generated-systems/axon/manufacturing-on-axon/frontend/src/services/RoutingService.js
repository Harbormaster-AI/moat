import axios from 'axios';

const ROUTING_API_BASE_URL = "/Routing";

class RoutingService {

    getRoutings(){
        return axios.get(ROUTING_API_BASE_URL + '/' );
    }

    createRouting(routing){
        return axios.post(ROUTING_API_BASE_URL  + '/create', routing);
    }

    getRoutingById(routingId){
        return axios.get(ROUTING_API_BASE_URL + '/load?routingId=' + routingId);
    }

    updateRouting(routing){
        return axios.put(ROUTING_API_BASE_URL + '/update', routing);
    }

    deleteRouting(routingId){
        return axios.delete(ROUTING_API_BASE_URL + '/delete?routingId=' + routingId);
    }
}

export default new RoutingService()