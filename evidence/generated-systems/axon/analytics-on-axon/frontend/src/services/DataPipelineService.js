import axios from 'axios';

const DATAPIPELINE_API_BASE_URL = "/DataPipeline";

class DataPipelineService {

    getDataPipelines(){
        return axios.get(DATAPIPELINE_API_BASE_URL + '/' );
    }

    createDataPipeline(dataPipeline){
        return axios.post(DATAPIPELINE_API_BASE_URL  + '/create', dataPipeline);
    }

    getDataPipelineById(dataPipelineId){
        return axios.get(DATAPIPELINE_API_BASE_URL + '/load?dataPipelineId=' + dataPipelineId);
    }

    updateDataPipeline(dataPipeline){
        return axios.put(DATAPIPELINE_API_BASE_URL + '/update', dataPipeline);
    }

    deleteDataPipeline(dataPipelineId){
        return axios.delete(DATAPIPELINE_API_BASE_URL + '/delete?dataPipelineId=' + dataPipelineId);
    }
}

export default new DataPipelineService()