import axios from 'axios';

const PERFORMANCEREVIEW_API_BASE_URL = "/PerformanceReview";

class PerformanceReviewService {

    getPerformanceReviews(){
        return axios.get(PERFORMANCEREVIEW_API_BASE_URL + '/' );
    }

    createPerformanceReview(performanceReview){
        return axios.post(PERFORMANCEREVIEW_API_BASE_URL  + '/create', performanceReview);
    }

    getPerformanceReviewById(performanceReviewId){
        return axios.get(PERFORMANCEREVIEW_API_BASE_URL + '/load?performanceReviewId=' + performanceReviewId);
    }

    updatePerformanceReview(performanceReview){
        return axios.put(PERFORMANCEREVIEW_API_BASE_URL + '/update', performanceReview);
    }

    deletePerformanceReview(performanceReviewId){
        return axios.delete(PERFORMANCEREVIEW_API_BASE_URL + '/delete?performanceReviewId=' + performanceReviewId);
    }
}

export default new PerformanceReviewService()