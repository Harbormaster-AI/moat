import axios from 'axios';

const DATAPROVIDER_API_BASE_URL = "/DataProvider";

class DataProviderService {

    getDataProviders(){
        return axios.get(DATAPROVIDER_API_BASE_URL + '/' );
    }

    createDataProvider(dataProvider){
        return axios.post(DATAPROVIDER_API_BASE_URL  + '/create', dataProvider);
    }

    getDataProviderById(dataProviderId){
        return axios.get(DATAPROVIDER_API_BASE_URL + '/load?dataProviderId=' + dataProviderId);
    }

    updateDataProvider(dataProvider){
        return axios.put(DATAPROVIDER_API_BASE_URL + '/update', dataProvider);
    }

    deleteDataProvider(dataProviderId){
        return axios.delete(DATAPROVIDER_API_BASE_URL + '/delete?dataProviderId=' + dataProviderId);
    }
}

export default new DataProviderService()