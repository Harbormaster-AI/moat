import axios from 'axios';

const GOAL_API_BASE_URL = "/Goal";

class GoalService {

    getGoals(){
        return axios.get(GOAL_API_BASE_URL + '/' );
    }

    createGoal(goal){
        return axios.post(GOAL_API_BASE_URL  + '/create', goal);
    }

    getGoalById(goalId){
        return axios.get(GOAL_API_BASE_URL + '/load?goalId=' + goalId);
    }

    updateGoal(goal){
        return axios.put(GOAL_API_BASE_URL + '/update', goal);
    }

    deleteGoal(goalId){
        return axios.delete(GOAL_API_BASE_URL + '/delete?goalId=' + goalId);
    }
}

export default new GoalService()