import axios from 'axios';

const DATACATEGORY_API_BASE_URL = "/DataCategory";

class DataCategoryService {

    getDataCategorys(){
        return axios.get(DATACATEGORY_API_BASE_URL + '/' );
    }

    createDataCategory(dataCategory){
        return axios.post(DATACATEGORY_API_BASE_URL  + '/create', dataCategory);
    }

    getDataCategoryById(dataCategoryId){
        return axios.get(DATACATEGORY_API_BASE_URL + '/load?dataCategoryId=' + dataCategoryId);
    }

    updateDataCategory(dataCategory){
        return axios.put(DATACATEGORY_API_BASE_URL + '/update', dataCategory);
    }

    deleteDataCategory(dataCategoryId){
        return axios.delete(DATACATEGORY_API_BASE_URL + '/delete?dataCategoryId=' + dataCategoryId);
    }
}

export default new DataCategoryService()