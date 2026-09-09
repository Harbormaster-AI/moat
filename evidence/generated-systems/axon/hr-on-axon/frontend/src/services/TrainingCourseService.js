import axios from 'axios';

const TRAININGCOURSE_API_BASE_URL = "/TrainingCourse";

class TrainingCourseService {

    getTrainingCourses(){
        return axios.get(TRAININGCOURSE_API_BASE_URL + '/' );
    }

    createTrainingCourse(trainingCourse){
        return axios.post(TRAININGCOURSE_API_BASE_URL  + '/create', trainingCourse);
    }

    getTrainingCourseById(trainingCourseId){
        return axios.get(TRAININGCOURSE_API_BASE_URL + '/load?trainingCourseId=' + trainingCourseId);
    }

    updateTrainingCourse(trainingCourse){
        return axios.put(TRAININGCOURSE_API_BASE_URL + '/update', trainingCourse);
    }

    deleteTrainingCourse(trainingCourseId){
        return axios.delete(TRAININGCOURSE_API_BASE_URL + '/delete?trainingCourseId=' + trainingCourseId);
    }
}

export default new TrainingCourseService()