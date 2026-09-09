import axios from 'axios';

const INSPECTIONPLAN_API_BASE_URL = "/InspectionPlan";

class InspectionPlanService {

    getInspectionPlans(){
        return axios.get(INSPECTIONPLAN_API_BASE_URL + '/' );
    }

    createInspectionPlan(inspectionPlan){
        return axios.post(INSPECTIONPLAN_API_BASE_URL  + '/create', inspectionPlan);
    }

    getInspectionPlanById(inspectionPlanId){
        return axios.get(INSPECTIONPLAN_API_BASE_URL + '/load?inspectionPlanId=' + inspectionPlanId);
    }

    updateInspectionPlan(inspectionPlan){
        return axios.put(INSPECTIONPLAN_API_BASE_URL + '/update', inspectionPlan);
    }

    deleteInspectionPlan(inspectionPlanId){
        return axios.delete(INSPECTIONPLAN_API_BASE_URL + '/delete?inspectionPlanId=' + inspectionPlanId);
    }
}

export default new InspectionPlanService()