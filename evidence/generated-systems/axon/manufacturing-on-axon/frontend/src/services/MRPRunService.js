import axios from 'axios';

const MRPRUN_API_BASE_URL = "/MRPRun";

class MRPRunService {

    getMRPRuns(){
        return axios.get(MRPRUN_API_BASE_URL + '/' );
    }

    createMRPRun(mRPRun){
        return axios.post(MRPRUN_API_BASE_URL  + '/create', mRPRun);
    }

    getMRPRunById(mRPRunId){
        return axios.get(MRPRUN_API_BASE_URL + '/load?mRPRunId=' + mRPRunId);
    }

    updateMRPRun(mRPRun){
        return axios.put(MRPRUN_API_BASE_URL + '/update', mRPRun);
    }

    deleteMRPRun(mRPRunId){
        return axios.delete(MRPRUN_API_BASE_URL + '/delete?mRPRunId=' + mRPRunId);
    }
}

export default new MRPRunService()