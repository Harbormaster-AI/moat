import axios from 'axios';

const EVALUATIONMETRIC_API_BASE_URL = "/EvaluationMetric";

class EvaluationMetricService {

    getEvaluationMetrics(){
        return axios.get(EVALUATIONMETRIC_API_BASE_URL + '/' );
    }

    createEvaluationMetric(evaluationMetric){
        return axios.post(EVALUATIONMETRIC_API_BASE_URL  + '/create', evaluationMetric);
    }

    getEvaluationMetricById(evaluationMetricId){
        return axios.get(EVALUATIONMETRIC_API_BASE_URL + '/load?evaluationMetricId=' + evaluationMetricId);
    }

    updateEvaluationMetric(evaluationMetric){
        return axios.put(EVALUATIONMETRIC_API_BASE_URL + '/update', evaluationMetric);
    }

    deleteEvaluationMetric(evaluationMetricId){
        return axios.delete(EVALUATIONMETRIC_API_BASE_URL + '/delete?evaluationMetricId=' + evaluationMetricId);
    }
}

export default new EvaluationMetricService()