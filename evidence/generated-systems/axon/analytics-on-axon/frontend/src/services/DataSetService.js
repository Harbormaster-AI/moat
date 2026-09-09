import axios from 'axios';

const DATASET_API_BASE_URL = "/DataSet";

class DataSetService {

    getDataSets(){
        return axios.get(DATASET_API_BASE_URL + '/' );
    }

    createDataSet(dataSet){
        return axios.post(DATASET_API_BASE_URL  + '/create', dataSet);
    }

    getDataSetById(dataSetId){
        return axios.get(DATASET_API_BASE_URL + '/load?dataSetId=' + dataSetId);
    }

    updateDataSet(dataSet){
        return axios.put(DATASET_API_BASE_URL + '/update', dataSet);
    }

    deleteDataSet(dataSetId){
        return axios.delete(DATASET_API_BASE_URL + '/delete?dataSetId=' + dataSetId);
    }
}

export default new DataSetService()