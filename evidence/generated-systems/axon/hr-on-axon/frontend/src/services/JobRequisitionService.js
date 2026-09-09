import axios from 'axios';

const JOBREQUISITION_API_BASE_URL = "/JobRequisition";

class JobRequisitionService {

    getJobRequisitions(){
        return axios.get(JOBREQUISITION_API_BASE_URL + '/' );
    }

    createJobRequisition(jobRequisition){
        return axios.post(JOBREQUISITION_API_BASE_URL  + '/create', jobRequisition);
    }

    getJobRequisitionById(jobRequisitionId){
        return axios.get(JOBREQUISITION_API_BASE_URL + '/load?jobRequisitionId=' + jobRequisitionId);
    }

    updateJobRequisition(jobRequisition){
        return axios.put(JOBREQUISITION_API_BASE_URL + '/update', jobRequisition);
    }

    deleteJobRequisition(jobRequisitionId){
        return axios.delete(JOBREQUISITION_API_BASE_URL + '/delete?jobRequisitionId=' + jobRequisitionId);
    }
}

export default new JobRequisitionService()