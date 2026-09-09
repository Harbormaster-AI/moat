import axios from 'axios';

const PAYROLLRUN_API_BASE_URL = "/PayrollRun";

class PayrollRunService {

    getPayrollRuns(){
        return axios.get(PAYROLLRUN_API_BASE_URL + '/' );
    }

    createPayrollRun(payrollRun){
        return axios.post(PAYROLLRUN_API_BASE_URL  + '/create', payrollRun);
    }

    getPayrollRunById(payrollRunId){
        return axios.get(PAYROLLRUN_API_BASE_URL + '/load?payrollRunId=' + payrollRunId);
    }

    updatePayrollRun(payrollRun){
        return axios.put(PAYROLLRUN_API_BASE_URL + '/update', payrollRun);
    }

    deletePayrollRun(payrollRunId){
        return axios.delete(PAYROLLRUN_API_BASE_URL + '/delete?payrollRunId=' + payrollRunId);
    }
}

export default new PayrollRunService()