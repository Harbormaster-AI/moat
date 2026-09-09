import axios from 'axios';

const CONTENTCATEGORY_API_BASE_URL = "/ContentCategory";

class ContentCategoryService {

    getContentCategorys(){
        return axios.get(CONTENTCATEGORY_API_BASE_URL + '/' );
    }

    createContentCategory(contentCategory){
        return axios.post(CONTENTCATEGORY_API_BASE_URL  + '/create', contentCategory);
    }

    getContentCategoryById(contentCategoryId){
        return axios.get(CONTENTCATEGORY_API_BASE_URL + '/load?contentCategoryId=' + contentCategoryId);
    }

    updateContentCategory(contentCategory){
        return axios.put(CONTENTCATEGORY_API_BASE_URL + '/update', contentCategory);
    }

    deleteContentCategory(contentCategoryId){
        return axios.delete(CONTENTCATEGORY_API_BASE_URL + '/delete?contentCategoryId=' + contentCategoryId);
    }
}

export default new ContentCategoryService()