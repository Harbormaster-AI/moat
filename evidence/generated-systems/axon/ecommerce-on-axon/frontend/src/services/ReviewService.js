import axios from 'axios';

const REVIEW_API_BASE_URL = "/Review";

class ReviewService {

    getReviews(){
        return axios.get(REVIEW_API_BASE_URL + '/' );
    }

    createReview(review){
        return axios.post(REVIEW_API_BASE_URL  + '/create', review);
    }

    getReviewById(reviewId){
        return axios.get(REVIEW_API_BASE_URL + '/load?reviewId=' + reviewId);
    }

    updateReview(review){
        return axios.put(REVIEW_API_BASE_URL + '/update', review);
    }

    deleteReview(reviewId){
        return axios.delete(REVIEW_API_BASE_URL + '/delete?reviewId=' + reviewId);
    }
}

export default new ReviewService()