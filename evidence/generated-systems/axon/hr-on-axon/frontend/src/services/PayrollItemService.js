import axios from 'axios';

const PAYROLLITEM_API_BASE_URL = "/PayrollItem";

class PayrollItemService {

    getPayrollItems(){
        return axios.get(PAYROLLITEM_API_BASE_URL + '/' );
    }

    createPayrollItem(payrollItem){
        return axios.post(PAYROLLITEM_API_BASE_URL  + '/create', payrollItem);
    }

    getPayrollItemById(payrollItemId){
        return axios.get(PAYROLLITEM_API_BASE_URL + '/load?payrollItemId=' + payrollItemId);
    }

    updatePayrollItem(payrollItem){
        return axios.put(PAYROLLITEM_API_BASE_URL + '/update', payrollItem);
    }

    deletePayrollItem(payrollItemId){
        return axios.delete(PAYROLLITEM_API_BASE_URL + '/delete?payrollItemId=' + payrollItemId);
    }
}

export default new PayrollItemService()