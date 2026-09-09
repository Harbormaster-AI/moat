import axios from 'axios';

const DISPOSITIONREVIEW_API_BASE_URL = "/DispositionReview";

class DispositionReviewService {

    getDispositionReviews(){
        return axios.get(DISPOSITIONREVIEW_API_BASE_URL + '/' );
    }

    createDispositionReview(dispositionReview){
        return axios.post(DISPOSITIONREVIEW_API_BASE_URL  + '/create', dispositionReview);
    }

    getDispositionReviewById(dispositionReviewId){
        return axios.get(DISPOSITIONREVIEW_API_BASE_URL + '/load?dispositionReviewId=' + dispositionReviewId);
    }

    updateDispositionReview(dispositionReview){
        return axios.put(DISPOSITIONREVIEW_API_BASE_URL + '/update', dispositionReview);
    }

    deleteDispositionReview(dispositionReviewId){
        return axios.delete(DISPOSITIONREVIEW_API_BASE_URL + '/delete?dispositionReviewId=' + dispositionReviewId);
    }
}

export default new DispositionReviewService()