import axios from 'axios';

const ONBOARDINGTASK_API_BASE_URL = "/OnboardingTask";

class OnboardingTaskService {

    getOnboardingTasks(){
        return axios.get(ONBOARDINGTASK_API_BASE_URL + '/' );
    }

    createOnboardingTask(onboardingTask){
        return axios.post(ONBOARDINGTASK_API_BASE_URL  + '/create', onboardingTask);
    }

    getOnboardingTaskById(onboardingTaskId){
        return axios.get(ONBOARDINGTASK_API_BASE_URL + '/load?onboardingTaskId=' + onboardingTaskId);
    }

    updateOnboardingTask(onboardingTask){
        return axios.put(ONBOARDINGTASK_API_BASE_URL + '/update', onboardingTask);
    }

    deleteOnboardingTask(onboardingTaskId){
        return axios.delete(ONBOARDINGTASK_API_BASE_URL + '/delete?onboardingTaskId=' + onboardingTaskId);
    }
}

export default new OnboardingTaskService()