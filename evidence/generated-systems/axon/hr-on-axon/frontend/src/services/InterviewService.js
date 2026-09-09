import axios from 'axios';

const INTERVIEW_API_BASE_URL = "/Interview";

class InterviewService {

    getInterviews(){
        return axios.get(INTERVIEW_API_BASE_URL + '/' );
    }

    createInterview(interview){
        return axios.post(INTERVIEW_API_BASE_URL  + '/create', interview);
    }

    getInterviewById(interviewId){
        return axios.get(INTERVIEW_API_BASE_URL + '/load?interviewId=' + interviewId);
    }

    updateInterview(interview){
        return axios.put(INTERVIEW_API_BASE_URL + '/update', interview);
    }

    deleteInterview(interviewId){
        return axios.delete(INTERVIEW_API_BASE_URL + '/delete?interviewId=' + interviewId);
    }
}

export default new InterviewService()