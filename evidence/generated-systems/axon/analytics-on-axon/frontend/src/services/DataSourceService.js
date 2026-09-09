import axios from 'axios';

const DATASOURCE_API_BASE_URL = "/DataSource";

class DataSourceService {

    getDataSources(){
        return axios.get(DATASOURCE_API_BASE_URL + '/' );
    }

    createDataSource(dataSource){
        return axios.post(DATASOURCE_API_BASE_URL  + '/create', dataSource);
    }

    getDataSourceById(dataSourceId){
        return axios.get(DATASOURCE_API_BASE_URL + '/load?dataSourceId=' + dataSourceId);
    }

    updateDataSource(dataSource){
        return axios.put(DATASOURCE_API_BASE_URL + '/update', dataSource);
    }

    deleteDataSource(dataSourceId){
        return axios.delete(DATASOURCE_API_BASE_URL + '/delete?dataSourceId=' + dataSourceId);
    }
}

export default new DataSourceService()