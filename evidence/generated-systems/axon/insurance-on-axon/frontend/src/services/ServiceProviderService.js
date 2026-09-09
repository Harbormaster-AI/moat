import axios from 'axios';

const SERVICEPROVIDER_API_BASE_URL = "/ServiceProvider";

class ServiceProviderService {

    getServiceProviders(){
        return axios.get(SERVICEPROVIDER_API_BASE_URL + '/' );
    }

    createServiceProvider(serviceProvider){
        return axios.post(SERVICEPROVIDER_API_BASE_URL  + '/create', serviceProvider);
    }

    getServiceProviderById(serviceProviderId){
        return axios.get(SERVICEPROVIDER_API_BASE_URL + '/load?serviceProviderId=' + serviceProviderId);
    }

    updateServiceProvider(serviceProvider){
        return axios.put(SERVICEPROVIDER_API_BASE_URL + '/update', serviceProvider);
    }

    deleteServiceProvider(serviceProviderId){
        return axios.delete(SERVICEPROVIDER_API_BASE_URL + '/delete?serviceProviderId=' + serviceProviderId);
    }
}

export default new ServiceProviderService()