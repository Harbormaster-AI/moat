import axios from 'axios';

const TAG_API_BASE_URL = "/Tag";

class TagService {

    getTags(){
        return axios.get(TAG_API_BASE_URL + '/' );
    }

    createTag(tag){
        return axios.post(TAG_API_BASE_URL  + '/create', tag);
    }

    getTagById(tagId){
        return axios.get(TAG_API_BASE_URL + '/load?tagId=' + tagId);
    }

    updateTag(tag){
        return axios.put(TAG_API_BASE_URL + '/update', tag);
    }

    deleteTag(tagId){
        return axios.delete(TAG_API_BASE_URL + '/delete?tagId=' + tagId);
    }
}

export default new TagService()