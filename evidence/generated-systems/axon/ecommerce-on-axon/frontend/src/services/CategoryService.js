import axios from 'axios';

const CATEGORY_API_BASE_URL = "/Category";

class CategoryService {

    getCategorys(){
        return axios.get(CATEGORY_API_BASE_URL + '/' );
    }

    createCategory(category){
        return axios.post(CATEGORY_API_BASE_URL  + '/create', category);
    }

    getCategoryById(categoryId){
        return axios.get(CATEGORY_API_BASE_URL + '/load?categoryId=' + categoryId);
    }

    updateCategory(category){
        return axios.put(CATEGORY_API_BASE_URL + '/update', category);
    }

    deleteCategory(categoryId){
        return axios.delete(CATEGORY_API_BASE_URL + '/delete?categoryId=' + categoryId);
    }
}

export default new CategoryService()