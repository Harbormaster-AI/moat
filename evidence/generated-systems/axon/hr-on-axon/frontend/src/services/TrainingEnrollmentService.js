import axios from 'axios';

const TRAININGENROLLMENT_API_BASE_URL = "/TrainingEnrollment";

class TrainingEnrollmentService {

    getTrainingEnrollments(){
        return axios.get(TRAININGENROLLMENT_API_BASE_URL + '/' );
    }

    createTrainingEnrollment(trainingEnrollment){
        return axios.post(TRAININGENROLLMENT_API_BASE_URL  + '/create', trainingEnrollment);
    }

    getTrainingEnrollmentById(trainingEnrollmentId){
        return axios.get(TRAININGENROLLMENT_API_BASE_URL + '/load?trainingEnrollmentId=' + trainingEnrollmentId);
    }

    updateTrainingEnrollment(trainingEnrollment){
        return axios.put(TRAININGENROLLMENT_API_BASE_URL + '/update', trainingEnrollment);
    }

    deleteTrainingEnrollment(trainingEnrollmentId){
        return axios.delete(TRAININGENROLLMENT_API_BASE_URL + '/delete?trainingEnrollmentId=' + trainingEnrollmentId);
    }
}

export default new TrainingEnrollmentService()